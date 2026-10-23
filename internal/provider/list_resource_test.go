package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccListResource(t *testing.T) {
	name := acctest.RandomWithPrefix("tfacc-list")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + fmt.Sprintf(`
resource "impart_list" "test" {
  name = %q
  kind = "string"
}
`, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_list.test", "name", name),
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
				Config: providerConfig + fmt.Sprintf(`
resource "impart_list" "test" {
  name = "%s-updated"
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
`, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_list.test", "name", name+"-updated"),
					resource.TestCheckResourceAttr("impart_list.test", "kind", "string"),
					resource.TestCheckResourceAttr("impart_list.test", "items.0.value", "item1"),
					resource.TestCheckResourceAttr("impart_list.test", "items.1.value", "item2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
