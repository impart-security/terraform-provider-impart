# Create a new specification
resource "impart_spec" "example" {
  name        = "spec_example"
  source_file = "${path.module}/spec.yaml"
  source_hash = "<sha256 hash for the source_file content>"
}
