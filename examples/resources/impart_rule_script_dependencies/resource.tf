resource "impart_rule_script_dependencies" "example" {
  dependencies = [
    {
      "rule_script_id" : resource.impart_rule_script.example_1.id,
      "depends_on_rule_script_ids" : [resource.impart_rule_script.example_2.id]
    }
  ]
}
