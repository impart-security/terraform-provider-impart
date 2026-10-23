package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMonitorResource(t *testing.T) {
	t.Skip()
	name := acctest.RandomWithPrefix("tfacc-monitor")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + fmt.Sprintf(`
data "impart_connector" "test" {
	id = ""
}

resource "impart_notification_template" "test" {
  name    = %[1]q
  payload = "This is notification template made from terraform"
  subject = "Terraform Subject"
}

resource "impart_monitor" "test_metric" {
  name        = %[2]q
  description = "test metric monitor"
  notifications = [
    {
      notification_template_id = impart_notification_template.test.id
      connector_id             = data.impart_connector.test.id
      destination              = "C060E7NGB6X"
    }
  ]
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
`, name+"-template", name),
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
