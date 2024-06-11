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
  source_hash = "4f501b53775586d59458a5d1c3eda6e1ef195d746895dd37b93db033f378e04c"
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
