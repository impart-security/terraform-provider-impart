
# Create a new rule client identifier storage
resource "impart_rule_client_identifier_storage" "client_identifier_storage" {
  name        = "client identifier storage"
  description = "terraform client identifier storage"
  capacity    = 10000
  storage_id  = resource.impart_rule_client_identifier.example.id
}

