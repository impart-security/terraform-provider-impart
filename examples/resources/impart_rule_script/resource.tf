data "local_file" "rule_example" {
  filename = "${path.module}/rule.js"
}

# Create a new rule script
resource "impart_rule_script" "example" {
  name        = "example"
  disabled    = false
  description = "Rule description"
  source_file = data.local_file.rule_example.filename
  source_hash = data.local_file.rule_example.content_sha256
}
