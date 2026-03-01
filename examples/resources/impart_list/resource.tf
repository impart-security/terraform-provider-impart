# Create a new string list
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

# Create a regex pattern list
resource "impart_list" "regex_example" {
  name        = "regex_list_example"
  kind        = "regex"
  description = "A list of regex patterns for matching"
  items = [
    {
      value = "^/api/v[0-9]+/users/.*$",
    },
    {
      value = "\\.(php|asp|aspx)$",
    }
  ]
}

# Create an Aho-Corasick pattern list
resource "impart_list" "aho_corasick_example" {
  name        = "aho_corasick_list_example"
  kind        = "aho_corasick"
  description = "A list of patterns for Aho-Corasick matching"
  items = [
    {
      value = "SELECT * FROM",
    },
    {
      value = "DROP TABLE",
    }
  ]
}
