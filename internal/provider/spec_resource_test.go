package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSpecResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "impart_spec" "test" {
  name        = "terraform_test"
  source_file = "./testdata/spec.yaml"
  source_hash = "4f501b53775586d59458a5d1c3eda6e1ef195d746895dd37b93db033f378e04c"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_spec.test", "name", "terraform_test"),
					resource.TestCheckResourceAttr("impart_spec.test", "source_file", "./testdata/spec.yaml"),
					resource.TestCheckResourceAttr("impart_spec.test", "source_hash", "4f501b53775586d59458a5d1c3eda6e1ef195d746895dd37b93db033f378e04c"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("impart_spec.test", "id"),
				),
			},
			// ImportState testing
			// This test is failing due to the issue with unpublished plugins
			// https://github.com/hashicorp/terraform-plugin-sdk/issues/1171
			// {
			// 	ResourceName:      "impart_spec.test",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
			// Update testing
			{
				Config: providerConfig + `
resource "impart_spec" "test" {
  name        = "terraform_test"
  source_file = "./testdata/spec_update.yaml"
  source_hash = "82cef44de7a05cbfda9f58be61c00d521f872777534890cc8c2d5ecc0f832d9a"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_spec.test", "name", "terraform_test"),
					resource.TestCheckResourceAttr("impart_spec.test", "source_file", "./testdata/spec_update.yaml"),
					resource.TestCheckResourceAttr("impart_spec.test", "source_hash", "82cef44de7a05cbfda9f58be61c00d521f872777534890cc8c2d5ecc0f832d9a"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("impart_spec.test", "id"),
				),
			},
		},
	})
}
