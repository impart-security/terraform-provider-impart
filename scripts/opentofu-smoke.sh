#!/usr/bin/env bash
#
# End-to-end smoke test of a *published* provider release under OpenTofu.
# Creates and destroys one impart_label in the org of IMPART_TOKEN.
#
#   IMPART_TOKEN=... scripts/opentofu-smoke.sh            # OpenTofu only
#   IMPART_TOKEN=... MIGRATE=1 scripts/opentofu-smoke.sh  # Terraform -> OpenTofu
#
# Environment:
#   IMPART_TOKEN      required; IMPART_ENDPOINT optional (provider defaults apply)
#   SMOKE_VERSION     provider version to install (default: the VERSION file)
#   MIGRATE=1         create the resource with terraform, then take the same
#                     directory and state over with tofu
#   ALLOW_UNSIGNED=1  only warn (don't fail) when tofu skips signature
#                     verification, e.g. to run the migration check before
#                     the signing key is registered with OpenTofu
#
# Fails if:
#   - `tofu init` installs the provider without verifying its signature
#     (the OpenTofu Registry has no key on file for impart-security);
#   - `plan -detailed-exitcode` after apply / after migration shows changes;
#   - any init/apply/plan/destroy step fails.

set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
version="${SMOKE_VERSION:-$(tr -d '[:space:]' <"$ROOT/VERSION")}"
version="${version#v}"
migrate="${MIGRATE:-0}"

: "${IMPART_TOKEN:?IMPART_TOKEN must be set (the label is created in that org)}"
tools=(tofu jq)
if [[ $migrate == 1 ]]; then
    tools+=(terraform)
fi
for tool in "${tools[@]}"; do
    if ! command -v "$tool" >/dev/null 2>&1; then
        echo "error: '$tool' not found on PATH" >&2
        exit 2
    fi
done

work="$(mktemp -d)"
slug="tfacc-smoke-$(date +%s)-$RANDOM"
created_by=""

# created_by is the CLI that last initialized the working directory and owns its
# state; cleanup destroys with it first and falls back to the other CLI
# (re-initializing, since each CLI rewrites the lock file to its own registry).
cleanup() {
    status=$?
    if [[ -n $created_by && -f "$work/terraform.tfstate" ]]; then
        local cli other=tofu destroyed=0
        [[ $created_by == tofu ]] && other=terraform
        for cli in "$created_by" "$other"; do
            command -v "$cli" >/dev/null 2>&1 || continue
            echo "== cleanup: destroying $slug with $cli"
            if (cd "$work" && { [[ $cli == "$created_by" ]] || "$cli" init -input=false -no-color >/dev/null 2>&1; } &&
                "$cli" destroy -auto-approve -input=false -no-color >/dev/null 2>&1); then
                destroyed=1
                break
            fi
        done
        if [[ $destroyed == 0 ]]; then
            echo "::warning::could not destroy impart_label $slug; state kept in $work for manual cleanup"
            exit "$status"
        fi
    fi
    rm -rf "$work"
    exit "$status"
}
trap cleanup EXIT

cat >"$work/main.tf" <<EOF
terraform {
  required_providers {
    impart = {
      source  = "impart-security/impart"
      version = "= ${version}"
    }
  }
}

provider "impart" {}

resource "impart_label" "smoke" {
  slug         = "${slug}"
  display_name = "${slug}"
  description  = "OpenTofu smoke test; safe to delete"
  color        = "gray"
}
EOF

cd "$work"
# Use the public registries, not a developer's dev_overrides or mirrors.
export TF_CLI_CONFIG_FILE=/dev/null
export TF_IN_AUTOMATION=1

tofu_init() {
    local out
    out="$(tofu init -input=false -no-color 2>&1)" || {
        echo "$out"
        return 1
    }
    # tofu now owns the directory (it rewrote the lock file to registry.opentofu.org),
    # even if the signature policy below rejects the install.
    created_by=tofu
    echo "$out" | grep -E "impart-security/impart|[Ss]ignature" || true
    if grep -q "Signature validation was skipped" <<<"$out"; then
        msg="tofu installed impart-security/impart $version WITHOUT signature verification; the OpenTofu Registry has no signing key on file (issue #18241 H1)"
        if [[ ${ALLOW_UNSIGNED:-0} == 1 ]]; then
            echo "::warning::$msg (continuing: ALLOW_UNSIGNED=1)"
        else
            echo "::error::$msg"
            return 1
        fi
    fi
}

if [[ $migrate == 1 ]]; then
    echo "== terraform: init + apply ($version)"
    terraform init -input=false -no-color >/dev/null
    created_by=terraform
    terraform apply -auto-approve -input=false -no-color | tail -n 1
    echo "state provider address: $(jq -r '.resources[0].provider' terraform.tfstate)"

    echo "== tofu: init on the Terraform working directory"
    tofu_init
    created_by=tofu
    echo "== tofu: plan on Terraform-created state (expect no changes)"
    tofu plan -detailed-exitcode -input=false -no-color | tail -n 3
else
    echo "== tofu: init ($version)"
    tofu_init
    created_by=tofu
    echo "== tofu: apply"
    tofu apply -auto-approve -input=false -no-color | tail -n 1
    echo "== tofu: plan (expect no changes)"
    tofu plan -detailed-exitcode -input=false -no-color | tail -n 3
fi

echo "== tofu: destroy"
tofu destroy -auto-approve -input=false -no-color | tail -n 1
created_by=""
echo "smoke test passed (version $version, migrate=$migrate)"
