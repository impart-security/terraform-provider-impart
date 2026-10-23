---
page_title: "Using the provider with OpenTofu"
subcategory: ""
description: |-
  Install and use the Impart provider with OpenTofu, verify its signature, and move existing Terraform configurations to OpenTofu.
---

# Using the provider with OpenTofu

The Impart provider is the same binary for Terraform and OpenTofu. Configuration, resources and
state are identical; nothing about the provider depends on which CLI you use.

Supported versions: OpenTofu 1.6 and later, Terraform 1.5 and later.

## Install

Use the same `required_providers` block as with Terraform:

```terraform
terraform {
  required_providers {
    impart = {
      source  = "impart-security/impart"
      version = "~> 0.18"
    }
  }
}

provider "impart" {
  # token can also be set with the IMPART_TOKEN environment variable
}
```

Then run:

```shell
tofu init
```

Under OpenTofu, `impart-security/impart` resolves to `registry.opentofu.org/impart-security/impart`.
The OpenTofu Registry serves the same release files that Impart publishes on GitHub.

## Signature verification

Every release's checksum file is signed with Impart's provider signing key:

| | |
|---|---|
| Key ID | `8223F728ECE52F4C` |
| Fingerprint | `591E FC11 04DB 6B51 FFA0 B52A 8223 F728 ECE5 2F4C` |
| User ID | Impart Security Terraform &lt;support@impart.security&gt; |

`tofu init` verifies the download against this key and reports it, for example:

```
- Installed impart-security/impart v0.18.0 (signed, key ID 8223F728ECE52F4C)
```

## Lock file for several platforms

When `tofu init` installs the provider from the OpenTofu Registry, it records the release's signed
package checksums for every platform in `.terraform.lock.hcl`, so the committed lock file already
works on other operating systems and architectures. If you install from a provider mirror or local
cache instead, OpenTofu can only record checksums for the platform it ran on; pre-populate the other
platforms with:

```shell
tofu providers lock \
  -platform=linux_amd64 \
  -platform=linux_arm64 \
  -platform=darwin_amd64 \
  -platform=darwin_arm64
```

Pin the provider version (as above) so every platform installs the same release.

## Moving an existing Terraform configuration to OpenTofu

No configuration or state changes are needed, and you don't need `tofu state replace-provider`.

`tofu init` rewrites the `.terraform.lock.hcl` entries from `registry.terraform.io/...` to
`registry.opentofu.org/...`, but it does **not** reliably carry over the version Terraform had
selected: where the constraint allows a range, it can install a newer version that still satisfies
it, so the migration silently becomes a provider upgrade. Pin the version you are already on first,
so the move changes the registry and nothing else.

1. Read the version recorded in `.terraform.lock.hcl` and pin it exactly in `required_providers`,
   for example `version = "0.18.0"`.
2. In the existing working directory, run `tofu init`. Check that the lock file records that same
   version against `registry.opentofu.org`, then commit it. State written by Terraform keeps
   working as is.
3. Run `tofu plan`. It should report no changes.
4. Restore the constraint you want afterwards, and take any provider upgrade as its own change.

Use one CLI per working directory. Each CLI's `tofu init` / `terraform init` rewrites the lock file to
its own registry and drops the other's entries, so alternating CLIs in the same directory keeps
changing the lock file (and fails `init -lockfile=readonly` in CI).

## State encryption

OpenTofu's state and plan encryption happens entirely in the CLI. It needs nothing from the
provider and works with the Impart provider as with any other.

## Debugging

To attach a debugger to the provider under OpenTofu, start it with the OpenTofu provider address:

```shell
terraform-provider-impart -debug -address=registry.opentofu.org/impart-security/impart
```

and export the `TF_REATTACH_PROVIDERS` value it prints in the shell where you run `tofu`.
