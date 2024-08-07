package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccApiBindingResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "impart_spec" "test_binding" {
  name        = "terraform_test"
  source_file = "./testdata/spec.yaml"
}

resource "impart_api_binding" "test" {
  name      = "terraform_test"
  port      = 443
  spec_id   = resource.impart_spec.test_binding.id
  hostname  = "example.com"
  base_path = "/"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_api_binding.test", "name", "terraform_test"),
					resource.TestCheckResourceAttr("impart_api_binding.test", "port", "443"),
					resource.TestCheckResourceAttr("impart_api_binding.test", "hostname", "example.com"),
					resource.TestCheckResourceAttr("impart_api_binding.test", "base_path", "/"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("impart_api_binding.test", "spec_id"),
					resource.TestCheckResourceAttrSet("impart_api_binding.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "impart_api_binding.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: providerConfig + `
			resource "impart_spec" "test_binding" {
			  name        = "terraform_test"
			  source_file = "./testdata/spec.yaml"
			}
			resource "impart_api_binding" "test" {
			  name      = "terraform_test_updated"
			  port      = 445
			  spec_id   = resource.impart_spec.test_binding.id
			  hostname  = "example2.net"
			  base_path = "/example"
			}
									`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_api_binding.test", "name", "terraform_test_updated"),
					resource.TestCheckResourceAttr("impart_api_binding.test", "port", "445"),
					resource.TestCheckResourceAttr("impart_api_binding.test", "hostname", "example2.net"),
					resource.TestCheckResourceAttr("impart_api_binding.test", "base_path", "/example"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("impart_api_binding.test", "spec_id"),
					resource.TestCheckResourceAttrSet("impart_api_binding.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
