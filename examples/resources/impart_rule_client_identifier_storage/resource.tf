# Create a new rule client identifier storage
resource "impart_rule_client_identifier_storage" "client_identifier_storage" {
  name        = "client identifier storage"
  description = "terraform client identifier storage"
  capacity    = 10000
}

# Store a client identifier in it
resource "impart_rule_client_identifier" "client_identifier" {
  name        = "client identifier"
  description = "terraform client identifier"
  storage_id  = resource.impart_rule_client_identifier_storage.client_identifier_storage.id
  hash_fields = [
    {
      field = "client_ip"
    },
  ]
}
