---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "impart_external_link Resource - impart"
subcategory: ""
description: |-
  
---

# impart_external_link (Resource)



## Example Usage

```terraform
# Create a new external link
resource "impart_external_link" "external_link1" {
  name              = "Datadog client IP address"
  description       = "A link to Datadog dashboard for client IP address"
  url               = "https://app.datadoghq.com/dashboard/3tm-mpc-863?tpl_var_ClientIp=9.37.130.233"
  entity            = "request"
  json_path_element = "$.client_ip.address"
  vendor            = "Datadog"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `entity` (String) The entity to which the links should be applied.
- `json_path_element` (String) A JSONPath to the element for which this link should apply (e.g. $.client_ip.address).
- `name` (String) The name of the external link.
- `url` (String) The external URL template with JSONPath element variables.
- `vendor` (String) The vendor for the external link.

### Optional

- `description` (String) The description of the external link.
- `spec_ids` (List of String) A list of spec IDs this external link applies to (empty means all).

### Read-Only

- `id` (String) Identifier for this list.
