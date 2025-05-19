resource "impart_rule_dependencies" "example" {
  dependencies = [
    {
      "rule_id" : "<example_1.id>",
      "depends_on" : ["<example_2.id>"]
    }
  ]
}
