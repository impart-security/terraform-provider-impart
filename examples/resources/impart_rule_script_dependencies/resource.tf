resource "impart_rule_script_dependencies" "example" {
  dependencies = [
    {
      "rule_script_id" : "<example_1.id>",
      "depends_on_rule_script_ids" : ["<example_2.id>"]
    }
  ]
}
