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
  source_hash = data.local_file.example.content_sha256 # set hash if spec changes tracking is disirable

  # lifecycle {
  #   ignore_changes = all
  # }
}

# Create a new notification template
resource "impart_notification_template" "test" {
  name         = "terraform_notification_template"
  connector_id = data.local.example_connector.id
  payload      = "This is notification template made from terraform"
  subject      = "Terraform Subject"
  destination  = ["C060E7NGB61"]
}

#Create a new event monitor
resource "impart_monitor" "test_event" {
  name                      = "terraform_event_monitor"
  description               = "test event monitor"
  notification_template_ids = [impart_notification_template.test.id]
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

resource "impart_monitor" "test_metric" {
  name                      = "terraform_event_monitor"
  description               = "test event monitor"
  notification_template_ids = [impart_notification_template.test.id]
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
