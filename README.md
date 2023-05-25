
# Impart Terraform Provider

A Terraform provider to manage Impart resources.

## Installation

- Login to https://console.impartsecurity.net/. Under manage section click Settings => Access Tokens => New API access token. Select scopes: read:org_api_access_tokens and scopes for resources will be managed by terraform.
- Set environment variable `IMPART_TOKEN`
- Set the following provider configuration

```
terraform {
  required_providers {
    impart = {
      source  = "impart-security/impart"
      "version" = "<plugin_version>"
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
resource "impart_binding" "example" {
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

## Development

### Build provider

Run the following command to build & install the provider

```shell
make
```

### Test provider

Acceptance testing will create and destroy real resources. IMPART_TOKEN needs to be specified.

```shell
make testacc
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
- Edit `$HOME/.terraformrc` and point "impart-security/impart" to your ${GOBIN} directory.

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
- Run the lanuch configuration
- From the debug console copy TF_REATTACH_PROVIDERS env variable and set in the shell which runs terraform commands. Example:

```
export TF_REATTACH_PROVIDERS='{"registry.terraform.io/impart-security/impart":{"Protocol":"grpc","ProtocolVersion":6,"Pid":26776,"Test":true,"Addr":{"Network":"unix","String":"/tmp/plugin460798854"}}}'
```
