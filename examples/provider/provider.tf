terraform {
  required_providers {
    impart = {
      source = "impart-security/impart"
    }
  }
}

# Configure the connection details for the impart service
provider "impart" {

}

# Read in an existing specification
data "impart_spec" "example" {
  id = "<id>"
}

# Reference local spec file
data "local_file" "example" {
  filename = "${path.module}/spec.yaml"
}

# Create a new specification
resource "impart_spec" "example" {
  name        = "terraform_test"
  source_file = data.local_file.example.filename
  source_hash = data.local_file.example.content_sha256 # set hash if spec changes tracking is desirable

  # lifecycle {
  #   ignore_changes = all
  # }
}

# Create a new list
resource "impart_list" "example" {
  name = "list_example"
  kind = "string"
}

# Create a new script rule
resource "impart_rule" "example" {
  name        = "example"
  type        = "script"
  disabled    = false
  description = "Rule description"
  content     = <<EOF
import { List } from "asruleslib";

// Reference the list
const blockedAgents = new List("${resource.impart_list.example.id}");

class BlockListedAgents extends BaseProcessorV2 {
  inspectRequestV2(metaInfo: MetaInfo, request: Request, tags: Set<string>): InspectResultV2 {
    if (request.header.has("user-agent") && blockedAgents.contains(request.header.get("user-agent"))) {
      tags.add("BLOCKED_AGENT");
    }

    return InspectResultV2.from(tags);
  }

  enforceRequestV2(metaInfo: MetaInfo, request: Request, tags: Set<string>): EnforceResultV2 {
    return tags.has("BLOCKED_AGENT") ? EnforceResultV2.BLOCK : EnforceResultV2.ALLOW;
  }
}

registerV2(new BlockListedAgents());
  EOF
}

# Read in an existing connector to deliver monitor notifications through
data "impart_connector" "example" {
  id = "<id>"
}

# Create a new notification template
resource "impart_notification_template" "test" {
  name    = "terraform_notification_template"
  payload = "This is notification template made from terraform"
  subject = "Terraform Subject"
}

# Create a new event monitor
resource "impart_monitor" "test_event" {
  name        = "terraform_event_monitor"
  description = "test event monitor"
  notifications = [
    {
      notification_template_id = resource.impart_notification_template.test.id
      connector_id             = data.impart_connector.example.id
      destination              = "C060E7NGB61"
    }
  ]
  conditions = [
    {
      threshold   = 1,
      comparator  = "gt",
      time_period = 60000,
      delay       = 0,
      details = {
        type         = "event",
        action       = "api_access_token_created",
        subject_type = "api_access_token_id",
        actor_type   = "user_id"
      }
    }
  ]
}

# Create a new metric monitor
resource "impart_monitor" "test_metric" {
  name        = "terraform_metric_monitor"
  description = "test metric monitor"
  notifications = [
    {
      notification_template_id = resource.impart_notification_template.test.id
      connector_id             = data.impart_connector.example.id
      destination              = "C060E7NGB61"
    }
  ]
  conditions = [
    {
      threshold   = 1,
      comparator  = "lt",
      time_period = 60000,
      delay       = 0,
      details = {
        type = "metric",
        tag  = "http-request"
      }
    }
  ]
}
