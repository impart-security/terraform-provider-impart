# ⚠️ This resource is deprecated. Please migrate to `impart_rule` instead.
# Create a new rule script
resource "impart_rule_script" "example" {
  name            = "example"
  disabled        = false
  description     = "Rule description"
  source_file     = "${path.module}/rule.js"
  source_hash     = "<sha256 hash for the source_file content>"
  blocking_effect = "block"
}

# Create a new rule script with content
resource "impart_rule_script" "example_content" {
  name            = "example"
  disabled        = false
  description     = "Rule description"
  content         = file("${path.module}/rule.js")
  blocking_effect = "block"
}
