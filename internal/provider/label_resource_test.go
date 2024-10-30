package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLabelResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "impart_label" "test" {
  slug         = "terraform-test-1"
  display_name = "terraform-test-1"
  description  = "test label description"
  color        = "red"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_label.test", "slug", "terraform-test-1"),
					resource.TestCheckResourceAttr("impart_label.test", "display_name", "terraform-test-1"),
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
				Config: providerConfig + `
resource "impart_label" "test" {
  slug         = "terraform-test-1"
  display_name = "updated-terraform-test-1"
  description  = "updated label description"
  color        = "blue"
}
									`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_label.test", "slug", "terraform-test-1"),
					resource.TestCheckResourceAttr("impart_label.test", "display_name", "updated-terraform-test-1"),
					resource.TestCheckResourceAttr("impart_label.test", "description", "updated label description"),
					resource.TestCheckResourceAttr("impart_label.test", "color", "blue"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
