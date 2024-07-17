package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSpecDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "impart_spec" "test" {
  name        = "terraform_test"
  source_file = "./testdata/spec.yaml"
  source_hash = "9bf6af660fcce87f4909073928d8bb051fafb6f6bb3245322de871d3c316e2a4"
}

data "impart_spec" "test" {
	id = impart_spec.test.id
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.impart_spec.test", "name", "terraform_test"),
					// Verify placeholder id attribute
					resource.TestCheckResourceAttrSet("data.impart_spec.test", "id"),
				),
			},
		},
	})
}
