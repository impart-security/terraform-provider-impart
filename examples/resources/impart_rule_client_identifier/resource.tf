
# Create a new rule client identifier
resource "impart_rule_client_identifier" "client_identifier" {
  name        = "client identifier"
  description = "terraform client identifier"
  hash_fields = [
    {
      field = "header_value"
      key   = "Authorization"
    },
    {
      field = "client_ip"
    },
  ]
}

