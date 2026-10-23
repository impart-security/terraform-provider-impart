#!/usr/bin/env bash
#
# Validate every example under examples/ with the Terraform or OpenTofu CLI,
# against a freshly built provider binary. No credentials are needed.
#
#   scripts/validate-examples.sh            # terraform
#   CLI=tofu scripts/validate-examples.sh   # OpenTofu
#
# The built binary is served from a filesystem mirror rather than dev_overrides:
# dev_overrides requires skipping `init`, which leaves the other providers the
# examples use (e.g. hashicorp/local) uninstalled. The mirror is laid out under
# both registry hostnames because a short source address resolves to
# registry.terraform.io under terraform and registry.opentofu.org under tofu.

set -euo pipefail

CLI="${CLI:-terraform}"
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
# Any version works; it only has to satisfy the unconstrained source in the examples.
MIRROR_VERSION="99.0.0"
HOSTS=(registry.terraform.io registry.opentofu.org)

for tool in "$CLI" go; do
    if ! command -v "$tool" >/dev/null 2>&1; then
        echo "error: '$tool' not found on PATH" >&2
        exit 2
    fi
done

WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

echo "Using $("$CLI" version | head -n 1)"

platform="$(go env GOOS)_$(go env GOARCH)"
binary="terraform-provider-impart_v${MIRROR_VERSION}"
(cd "$ROOT" && go build -o "$WORK/$binary" .)

includes=""
for host in "${HOSTS[@]}"; do
    dir="$WORK/mirror/$host/impart-security/impart/$MIRROR_VERSION/$platform"
    mkdir -p "$dir"
    cp "$WORK/$binary" "$dir/"
    includes="${includes:+$includes, }\"$host/impart-security/impart\""
done

# Hostname wildcards ("*/impart-security/impart") are rejected by the CLI, so
# both hosts are listed explicitly.
cat >"$WORK/cli.tfrc" <<EOF
provider_installation {
  filesystem_mirror {
    path    = "$WORK/mirror"
    include = [$includes]
  }
  direct {
    exclude = [$includes]
  }
}
EOF

export TF_CLI_CONFIG_FILE="$WORK/cli.tfrc"
export TF_PLUGIN_CACHE_DIR="$WORK/plugin-cache"
export TF_IN_AUTOMATION=1
mkdir -p "$TF_PLUGIN_CACHE_DIR"

failed=()
count=0
while IFS= read -r example; do
    count=$((count + 1))
    dest="$WORK/examples/$example"
    mkdir -p "$dest"
    cp -R "$ROOT/$example/." "$dest/"

    # A lock file left by a local `init` pins a published provider version,
    # and the mirror serves only $MIRROR_VERSION, so init would fail.
    rm -rf "$dest/.terraform" "$dest/.terraform.lock.hcl"

    # Resource and data source examples are documentation snippets without a
    # terraform block; without one the CLI would look for hashicorp/impart.
    if ! grep -qs required_providers "$dest"/*.tf; then
        cat >"$dest/zz_required_providers.tf" <<EOF
terraform {
  required_providers {
    impart = {
      source = "impart-security/impart"
    }
  }
}
EOF
    fi

    if output="$(cd "$dest" && "$CLI" init -backend=false -input=false -no-color 2>&1 &&
        "$CLI" validate -no-color 2>&1)"; then
        echo "ok   $example"
    else
        echo "FAIL $example"
        printf '     %s\n' "${output//$'\n'/$'\n     '}"
        failed+=("$example")
    fi
done < <(cd "$ROOT" && find examples -name '*.tf' -exec dirname {} \; | sort -u)

echo
if [ "${#failed[@]}" -gt 0 ]; then
    echo "${#failed[@]} of $count examples failed validation with $CLI:"
    printf '  %s\n' "${failed[@]}"
    exit 1
fi
echo "All $count examples are valid with $CLI."
