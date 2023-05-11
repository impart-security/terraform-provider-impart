# Create a new binding
resource "impart_binding" "example" {
  name      = "binding_example"
  port      = 443
  spec_id   = resource.impart_spec.example.id
  hostname  = "example.com"
  base_path = "/"
}
