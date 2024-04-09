package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestMonitorResource(t *testing.T) {
	t.Skip()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
data "impart_connector" "test" {
	id = ""
}

resource "impart_notification_template" "test" {
  name         = "terraform_notification_template"
  connector_id      = data.impart_connector.test.id
  payload = "This is notification template made from terraform"
  subject = "Terraform Subject"
  destination = ["C060E7NGB6X"]
}

resource "impart_monitor" "test_metric" {
  name                      = "terraform_metric_monitor"
  description               = "test metric monitor"
  notification_template_ids = [impart_notification_template.test.id]
  conditions = [
    {
      threshold   = 5,
      comparator  = "gt",
      time_period = 60000,
      delay       = 0,
      details = {
        type = "metric",
        tag  = "whoa"
      }
    }
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_monitor.test", "name", "terraform_monitor"),
					resource.TestCheckResourceAttr("impart_monitor.test", "description", "test monitor"),
					resource.TestCheckResourceAttrSet("impart_monitor.test", "conditions.0.threshold"),
					// Verify dynamic values have any impart_monitor set in the state.
					resource.TestCheckResourceAttrSet("impart_monitor.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
