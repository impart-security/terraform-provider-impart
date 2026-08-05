# Create a new notification template
#
# A notification template is message-only: name, subject, and payload. The connector
# and destination it is delivered through are set on a monitor's notifications, not here.
resource "impart_notification_template" "example" {
  name    = "notification_template_example"
  payload = "This is a test message payload"
  subject = "Test subject"
}
