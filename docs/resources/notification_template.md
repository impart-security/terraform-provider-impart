---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "impart_notification_template Resource - impart"
subcategory: ""
description: |-
  Manage an notification template.
---

# impart_notification_template (Resource)

Manage an notification template.

## Example Usage

```terraform
# Create a new notification template
resource "impart_notification_template" "example" {
  name         = "notification_template_example"
  connector_id = "<example_connector.id>"
  payload      = "This is a test message payload"
  subject      = "Test subject"
  destination  = ["test-destination-id"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connector_id` (String) The connector id.
- `destination` (List of String) An array of destination ids to which the payloads will be sent.
- `name` (String) The name for this notification template.
- `payload` (String) The payload message that will be sent to the Third Party API.
- `subject` (String) The subject message that will be sent to the Third Party API.

### Read-Only

- `id` (String) Identifier for this notification template.
