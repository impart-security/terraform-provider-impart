
# Impart Terraform Provider

A Terraform provider to manage Impart resources. It works with both **Terraform** (1.5 and later) and
**OpenTofu** (1.6 and later). The same binary serves both. See [Using with OpenTofu](#using-with-opentofu).

## Installation

- Login to https://console.impartsecurity.net/. Under manage section click Settings => Access Tokens => New API access token. Select scopes: read:org_api_access_tokens and scopes for resources will be managed by terraform.
- Set environment variable `IMPART_TOKEN`
- Set the following provider configuration

```
terraform {
  required_providers {
    impart = {
      source  = "impart-security/impart"
      version = "<plugin_version>"
    }
  }
}

provider "impart" {
}
```

Alternatively token can provided explicitly in the configuration:

```
provider "impart" {
  token = <token>
}
```

## Using Impart terraform provider

```
data "local_file" "example" {
  filename = "${path.module}/spec.yaml"
}

# manage specification resource
resource "impart_spec" "example" {
  name        = "example"
  source_file = data.local_file.example.filename
  source_hash = data.local_file.example.content_sha256 #optional to detect specification changes

  # uncomment if you would like to ignore changes to the specification resource
  # lifecycle {
  #   ignore_changes = all
  # }
}

# manage binding resource
resource "impart_api_binding" "example" {
  name      = "example"
  port      = 443
  spec_id   = resource.impart_spec.example.id
  hostname  = "example.com"
  base_path = "/"
}

data "local_file" "rule_example" {
  filename = "${path.module}/rule.js"
}

# manage rule script resource
resource "impart_rule_script" "example" {
  name        = "rule_script_example"
  disabled    = false
  description = "Rule description"
  source_file = data.local_file.rule_example.filename
  source_hash = data.local_file.rule_example.content_sha256 #optional to detect rule script changes
}
```

## Using with OpenTofu

The provider is the same binary for Terraform and OpenTofu; use the same `required_providers` block and
run `tofu init`. Under OpenTofu, `impart-security/impart` resolves to
`registry.opentofu.org/impart-security/impart`, which serves the same signed release files. Existing
Terraform state works unchanged. When migrating a directory, pin the version your lock file already
records before the first `tofu init`: `tofu init` rewrites the lock entries to the OpenTofu registry
but may select a newer version where the constraint allows one. The full guide is
[docs/guides/opentofu.md](./docs/guides/opentofu.md).

## Development

### Build provider

Run the following command to build & install the provider

```shell
make install
```

### Test provider

Acceptance testing will create and destroy real resources. IMPART_TOKEN needs to be specified.

```shell
make testacc          # with terraform on PATH
make testacc-tofu     # the same suite with OpenTofu (tofu on PATH)
```

`make testacc` registers the provider under `impart-security/impart` (`TF_ACC_PROVIDER_NAMESPACE`);
`make testacc-tofu` also sets `TF_ACC_TERRAFORM_PATH` to `tofu` and `TF_ACC_PROVIDER_HOST` to
`registry.opentofu.org`, and refuses to run without `tofu` rather than silently using Terraform.
Use `usingOpenTofu()` (in `internal/provider/provider_test.go`) for any CLI-specific test step; don't
use `tfversion` checks for that, because OpenTofu reports its version as a Terraform version.

### Validate examples

Every directory under `examples/` must pass `init` + `validate` with both CLIs. No credentials are
needed:

```shell
make validate-examples            # terraform
make validate-examples CLI=tofu   # OpenTofu
```

If you want to run acceptance tests from the vs code add env vars to the settings.json file
```
  "go.testEnvVars": {
    "TF_ACC": "1",
    "IMPART_TOKEN": "<access_token>"
  },
```
If you want to test out the provider locally with the `terraform` CLI.

- Run make
- Edit `$HOME/.terraformrc` (OpenTofu reads `$HOME/.tofurc`, falling back to `.terraformrc`) and point "impart-security/impart" to your ${GOBIN} directory.

```hcl
provider_installation {
  dev_overrides {
    "impart-security/impart" = "/home/me/go/path/bin/"
  }
  direct {}
}
```

And then create and test a few runs based on the files under examples. **NOTE:** You should not run `terraform init` when using _dev_overrides_.

### Documentation

Documentation is generated with [tfplugindocs](https://github.com/hashicorp/terraform-plugin-docs) and exists in the [docs](./docs/) directory.

```shell
make generate
```

### Debugging with VS code

Create launch configuration with -debug argument and IMPART_TOKEN env variable:

```
    {
      "name": "Debug Terraform Provider",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceFolder}/terraform-provider-impart",
      "env": {
        "IMPART_TOKEN": "<token>"
      },
      "args": ["-debug"]
    }
```

- Compile the plugin and set dev overrides as described above
- Run the launch configuration
- From the debug console copy TF_REATTACH_PROVIDERS env variable and set in the shell which runs terraform commands. Example:

```
export TF_REATTACH_PROVIDERS='{"registry.terraform.io/impart-security/impart":{"Protocol":"grpc","ProtocolVersion":6,"Pid":26776,"Test":true,"Addr":{"Network":"unix","String":"/tmp/plugin460798854"}}}'
```

For OpenTofu, add `"-address=registry.opentofu.org/impart-security/impart"` to `args` so the printed
`TF_REATTACH_PROVIDERS` is keyed by the address `tofu` expects.

## Parity between Terraform and OpenTofu

A large share of users run this provider with OpenTofu only, so both CLIs are first-class:

- Validate the examples with both CLIs before releasing: `make validate-examples` and
  `make validate-examples CLI=tofu`.
- `TestOpenTofuParity` fails if the provider adopts a capability OpenTofu doesn't support (list
  resources, actions, state stores). Don't relax it without product sign-off.
- When a new OpenTofu line reaches end of life, drop it from the supported versions in the docs.

## Releasing

1. Update `CHANGELOG.md` and bump `VERSION` in the monorepo. Merging the bump publishes the provider:
   the monorepo syncs this directory to `impart-security/terraform-provider-impart` and tags it, and
   that repository's release workflow builds and signs the release with GoReleaser.
2. Confirm the version appears **signed** in both registries (usually within an hour):

   ```shell
   V=<version>
   for r in registry.terraform.io registry.opentofu.org; do
     curl -s "https://$r/v1/providers/impart-security/impart/$V/download/linux/amd64" \
       | jq -r --arg r "$r" '"\($r): " + ([.signing_keys.gpg_public_keys[]?.key_id] | join(","))'
   done
   ```

   Both lines should show `8223F728ECE52F4C`.
3. **Signing key rotation.** Once a key is on file, OpenTofu rejects releases signed by any other key,
   so `tofu init` would fail for that version. Register the new key with the OpenTofu Registry
   (the "Submit new Provider Signing Key" issue form at <https://github.com/opentofu/registry>,
   submitted through the web UI by a *public* member of the `impart-security` GitHub org) **before**
   publishing a release signed with it, and add it to the Terraform Registry as usual.
