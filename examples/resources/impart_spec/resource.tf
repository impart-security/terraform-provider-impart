# Create a new specification
data "local_file" "example" {
  filename = "${path.module}/spec.yaml"
}

resource "impart_spec" "example" {
  name        = "spec_example"
  source_file = data.local_file.example.filename
  source_hash = data.local_file.example.content_sha256

  # uncomment to ignore spec name and source changes
  # lifecycle {
  #   ignore_changes = all
  # }
}
