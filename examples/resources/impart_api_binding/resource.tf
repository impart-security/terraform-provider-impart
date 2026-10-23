# Create a specification for the binding
resource "impart_spec" "example" {
  name        = "spec_example"
  source_file = "${path.module}/spec.yaml"
  source_hash = "<sha256 hash for the source_file content>"
}

# Create a new api binding
resource "impart_api_binding" "example" {
  name      = "api_binding_example"
  port      = 443
  spec_id   = resource.impart_spec.example.id
  hostname  = "example.com"
  base_path = "/"
}
