# Create a label to attach to the tag
resource "impart_label" "example" {
  slug         = "example"
  display_name = "Example"
  description  = "example label"
  color        = "red"
}

# Create a new tag metadata
resource "impart_tag_metadata" "example" {
  name         = "tag"
  description  = "tag description"
  external_url = "http://example.com"
  labels = [
    resource.impart_label.example.slug,
  ]
  risk_statement = "risk statement"
}
