
# Create a new rule script
resource "impart_rule" "example" {
  name            = "example"
  disabled        = false
  description     = "Rule description"
  source_file     = "${path.module}/rule.js"
  source_hash     = "<sha256 hash for the source_file content>"
  blocking_effect = "block"
  type            = "script"
}

# Create a new rule recipe with content
resource "impart_rule" "example_rule_recipe" {
  name            = "example"
  disabled        = false
  description     = "Rule description"
  content         = file("${path.module}/rule.json")
  blocking_effect = "block"
  type            = "recipe"
}
