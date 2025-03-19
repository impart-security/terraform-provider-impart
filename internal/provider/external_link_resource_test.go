package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccExternalLinkResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "impart_external_link" "test" {
  name              = "test"
  description       = "test description"
  url               = "http://datadog.com"
  entity            = "tag"
  json_path_element = "$.client_ip.address"
  vendor = "datadog"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_external_link.test", "name", "test"),
					resource.TestCheckResourceAttr("impart_external_link.test", "description", "test description"),
					resource.TestCheckResourceAttr("impart_external_link.test", "url", "http://datadog.com"),
					resource.TestCheckResourceAttr("impart_external_link.test", "entity", "tag"),
					resource.TestCheckResourceAttr("impart_external_link.test", "json_path_element", "$.client_ip.address"),
					resource.TestCheckResourceAttr("impart_external_link.test", "vendor", "datadog"),
				),
			},
			// Update and Read testing
			{
				Config: providerConfig + `
resource "impart_external_link" "test" {
  name              = "test2"
  description       = "test2 description"
  url               = "http://splunk.com"
  entity            = "request"
  json_path_element = "$.client_ip"
  vendor = "splunk"
}
									`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_external_link.test", "name", "test2"),
					resource.TestCheckResourceAttr("impart_external_link.test", "description", "test2 description"),
					resource.TestCheckResourceAttr("impart_external_link.test", "url", "http://splunk.com"),
					resource.TestCheckResourceAttr("impart_external_link.test", "entity", "request"),
					resource.TestCheckResourceAttr("impart_external_link.test", "json_path_element", "$.client_ip"),
					resource.TestCheckResourceAttr("impart_external_link.test", "vendor", "splunk"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
