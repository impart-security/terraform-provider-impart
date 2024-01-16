# Create a new api binding
resource "impart_api_binding" "example" {
  name      = "api_binding_example"
  port      = 443
  spec_id   = resource.impart_spec.example.id
  hostname  = "example.com"
  base_path = "/"
}
