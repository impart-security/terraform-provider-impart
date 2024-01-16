package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
  source_hash = "4f501b53775586d59458a5d1c3eda6e1ef195d746895dd37b93db033f378e04c"
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
			// // ImportState testing
			// {
			// 	ResourceName:      "impart_spec.test",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
			// Update and Read testing
			{
				Config: providerConfig + `
resource "impart_spec" "test_binding" {
  name        = "terraform_test"
  source_file = "./testdata/spec.yaml"
  source_hash = "4f501b53775586d59458a5d1c3eda6e1ef195d746895dd37b93db033f378e04c"
}
resource "impart_api_binding" "test" {
  name      = "terraform_test_updated"
  port      = 444
  spec_id   = resource.impart_spec.test_binding.id
  hostname  = "example.net"
  base_path = "/"
}
						`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_api_binding.test", "name", "terraform_test_updated"),
					resource.TestCheckResourceAttr("impart_api_binding.test", "port", "444"),
					resource.TestCheckResourceAttr("impart_api_binding.test", "hostname", "example.net"),
					resource.TestCheckResourceAttr("impart_api_binding.test", "base_path", "/"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("impart_api_binding.test", "spec_id"),
					resource.TestCheckResourceAttrSet("impart_api_binding.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
