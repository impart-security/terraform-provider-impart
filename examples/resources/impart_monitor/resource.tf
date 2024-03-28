# Create a new event monitor
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

# Create a new metric monitor
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
