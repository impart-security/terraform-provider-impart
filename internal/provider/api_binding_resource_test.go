package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccApiBindingResource(t *testing.T) {
	name := acctest.RandomWithPrefix("tfacc-binding")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + fmt.Sprintf(`
resource "impart_spec" "test_binding" {
  name        = %[1]q
  source_file = "./testdata/spec.yaml"
}

resource "impart_api_binding" "test" {
  name      = %[1]q
  port      = 443
  spec_id   = resource.impart_spec.test_binding.id
  hostname  = "%[1]s.example.com"
  base_path = "/"
}
`, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_api_binding.test", "name", name),
					resource.TestCheckResourceAttr("impart_api_binding.test", "port", "443"),
					resource.TestCheckResourceAttr("impart_api_binding.test", "hostname", name+".example.com"),
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
				Config: providerConfig + fmt.Sprintf(`
resource "impart_spec" "test_binding" {
  name        = %[1]q
  source_file = "./testdata/spec.yaml"
}

resource "impart_api_binding" "test" {
  name      = "%[1]s-updated"
  port      = 445
  spec_id   = resource.impart_spec.test_binding.id
  hostname  = "%[1]s.example.net"
  base_path = "/example"
}
`, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_api_binding.test", "name", name+"-updated"),
					resource.TestCheckResourceAttr("impart_api_binding.test", "port", "445"),
					resource.TestCheckResourceAttr("impart_api_binding.test", "hostname", name+".example.net"),
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
