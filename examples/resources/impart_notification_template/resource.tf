# Create a new notification template
resource "impart_notification_template" "example" {
  name         = "notification_template_example"
  connector_id = "<example_connector.id>"
  payload      = "This is a test message payload"
  subject      = "Test subject"
  destination  = ["test-destination-id"]
}
