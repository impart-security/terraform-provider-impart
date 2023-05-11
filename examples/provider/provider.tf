terraform {
  required_providers {
    impart = {
      source = "impart-security/impart"
    }
  }
}

# Configure the connection details for the impart service
provider "impart" {

}

# Read in an existing specification
data "impart_spec" "example" {
  id = "<id>"
}

# Reference local spec file
data "local_file" "example" {
  filename = "${path.module}/spec.yaml"
}

# Create a new specification
resource "impart_spec" "example" {
  name        = "terraform_test"
  source_file = data.local_file.example.filename
  source_hash = data.local_file.example.content_sha256 # set hash if spec changes tracking is disirable

  # lifecycle {
  #   ignore_changes = all
  # }
}
