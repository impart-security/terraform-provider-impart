# Create a new list
resource "impart_list" "example" {
  name        = "list_example"
  kind        = "string"
  description = "list description"
  items = [
    {
      value = "item1",
    },
    {
      value = "item2",
    }
  ]
}
