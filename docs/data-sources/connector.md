---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "impart_connector Data Source - impart"
subcategory: ""
description: |-
  Manage a connector.
---

# impart_connector (Data Source)

Manage a connector.

## Example Usage

```terraform
# Read in an existing specification
data "impart_connector" "example_connector" {
  id = "<id>"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Identifier for this connector.

### Optional

- `connector_type_id` (String) ID of the connector type (eg. ID for our Slack or Jira connector types).
- `is_connected` (Boolean) Whether or not the connector is authenticated via OAuth2.
- `name` (String) Name for this connector.
