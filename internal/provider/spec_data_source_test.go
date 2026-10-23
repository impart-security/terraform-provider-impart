package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSpecDataSource(t *testing.T) {
	name := acctest.RandomWithPrefix("tfacc-specds")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: providerConfig + fmt.Sprintf(`
resource "impart_spec" "test" {
  name        = %q
  source_file = "./testdata/spec.yaml"
  source_hash = "9bf6af660fcce87f4909073928d8bb051fafb6f6bb3245322de871d3c316e2a4"
}

data "impart_spec" "test" {
	id = impart_spec.test.id
}
`, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.impart_spec.test", "name", name),
					// Verify placeholder id attribute
					resource.TestCheckResourceAttrSet("data.impart_spec.test", "id"),
				),
			},
		},
	})
}
