package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccListResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "impart_list" "test" {
  name = "list_example"
  kind = "string"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_list.test", "name", "list_example"),
					resource.TestCheckResourceAttr("impart_list.test", "kind", "string"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "impart_list.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: providerConfig + `
resource "impart_list" "test" {
  name = "list_example_updated"
  kind = "string"
  items = [
	{
		value = "item1"
	},
	{
		value = "item2"
	}
  ]
}
									`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_list.test", "name", "list_example_updated"),
					resource.TestCheckResourceAttr("impart_list.test", "kind", "string"),
					resource.TestCheckResourceAttr("impart_list.test", "items.0.value", "item1"),
					resource.TestCheckResourceAttr("impart_list.test", "items.1.value", "item2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
