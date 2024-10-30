# Create a new label
resource "impart_label" "example" {
  slug         = "example"
  display_name = "Example"
  description  = "example label"
  color        = "red"
}
