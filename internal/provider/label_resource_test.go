package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLabelResource(t *testing.T) {
	slug := acctest.RandomWithPrefix("tfacc-label")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + fmt.Sprintf(`
resource "impart_label" "test" {
  slug         = %[1]q
  display_name = %[1]q
  description  = "test label description"
  color        = "red"
}
`, slug),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_label.test", "slug", slug),
					resource.TestCheckResourceAttr("impart_label.test", "display_name", slug),
					resource.TestCheckResourceAttr("impart_label.test", "description", "test label description"),
					resource.TestCheckResourceAttr("impart_label.test", "color", "red"),
				),
			},
			// ImportState testing
			// {
			// 	ResourceName:      "impart_label.test",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
			// Update and Read testing
			{
				Config: providerConfig + fmt.Sprintf(`
resource "impart_label" "test" {
  slug         = %[1]q
  display_name = "updated-%[1]s"
  description  = "updated label description"
  color        = "blue"
}
`, slug),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_label.test", "slug", slug),
					resource.TestCheckResourceAttr("impart_label.test", "display_name", "updated-"+slug),
					resource.TestCheckResourceAttr("impart_label.test", "description", "updated label description"),
					resource.TestCheckResourceAttr("impart_label.test", "color", "blue"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
