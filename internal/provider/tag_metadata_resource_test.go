package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTagMetadatalResource(t *testing.T) {
	name := acctest.RandomWithPrefix("tfacc-tag")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + fmt.Sprintf(`
resource "impart_tag_metadata" "test" {
  name         = %q
  description  = "test1 description"
  external_url = "http://example.com"
  risk_statement = "risk statement"
}
`, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_tag_metadata.test", "name", name),
					resource.TestCheckResourceAttr("impart_tag_metadata.test", "description", "test1 description"),
					resource.TestCheckResourceAttr("impart_tag_metadata.test", "external_url", "http://example.com"),
					resource.TestCheckResourceAttr("impart_tag_metadata.test", "risk_statement", "risk statement"),
				),
			},
			// ImportState testing
			// {
			// 	ResourceName:      "impart_tag_metadata.test",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
			// Update and Read testing
			{
				Config: providerConfig + fmt.Sprintf(`
resource "impart_tag_metadata" "test" {
  name         = %q
  description  = "test2 description"
  external_url = "http://example2.com"
  risk_statement = "updated risk statement"
}
`, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_tag_metadata.test", "name", name),
					resource.TestCheckResourceAttr("impart_tag_metadata.test", "description", "test2 description"),
					resource.TestCheckResourceAttr("impart_tag_metadata.test", "external_url", "http://example2.com"),
					resource.TestCheckResourceAttr("impart_tag_metadata.test", "risk_statement", "updated risk statement"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
