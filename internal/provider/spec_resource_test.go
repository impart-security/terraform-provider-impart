package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
  source_hash = "9bf6af660fcce87f4909073928d8bb051fafb6f6bb3245322de871d3c316e2a4"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_spec.test", "name", "terraform_test"),
					resource.TestCheckResourceAttr("impart_spec.test", "source_file", "./testdata/spec.yaml"),
					resource.TestCheckResourceAttr("impart_spec.test", "source_hash", "9bf6af660fcce87f4909073928d8bb051fafb6f6bb3245322de871d3c316e2a4"),
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
  source_hash = "cd3d7f3b6ac83c28d47c400dffd958e8840aed9f8804ff2030595e19837a784c"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_spec.test", "name", "terraform_test"),
					resource.TestCheckResourceAttr("impart_spec.test", "source_file", "./testdata/spec_update.yaml"),
					resource.TestCheckResourceAttr("impart_spec.test", "source_hash", "cd3d7f3b6ac83c28d47c400dffd958e8840aed9f8804ff2030595e19837a784c"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("impart_spec.test", "id"),
				),
			},
		},
	})
}
