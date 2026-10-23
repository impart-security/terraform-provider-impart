package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLogBindingResource(t *testing.T) {
	name := acctest.RandomWithPrefix("tfacc-logbinding")
	grokPattern := `%{HTTPDATE:timestamp} "(?:%{WORD:http_method}|-) (?:%{GREEDYDATA:request}|-) (?:HTTP/%{NUMBER:httpversion}|-( )?)" (?:%{NUMBER:response_code}|-)`
	updatedGrokPattern := grokPattern + " (?:%{WORD:variable}|-)"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + fmt.Sprintf(`
resource "impart_spec" "test" {
  name        = %[1]q
  source_file = "./testdata/spec.yaml"
}

resource "impart_log_binding" "test" {
  name         = %[1]q
  spec_id      = resource.impart_spec.test.id
  pattern_type = "grok"
  pattern = <<EOF
%[2]s
  EOF
  logstream_id = "%[1]s-logstream"
}`, name, strings.ReplaceAll(grokPattern, "%", "%%")),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_log_binding.test", "name", name),
					resource.TestCheckResourceAttr("impart_log_binding.test", "logstream_id", name+"-logstream"),
					resource.TestCheckResourceAttr("impart_log_binding.test", "pattern_type", "grok"),
					resource.TestCheckResourceAttr("impart_log_binding.test", "pattern", fmt.Sprintf("%s\n", grokPattern)),
					// Verify dynamic values have any impart_log_binding set in the state.
					resource.TestCheckResourceAttrSet("impart_log_binding.test", "spec_id"),
					resource.TestCheckResourceAttrSet("impart_log_binding.test", "id"),
				),
			},
			{
				Config: providerConfig + fmt.Sprintf(`
resource "impart_spec" "test" {
  name        = %[1]q
  source_file = "./testdata/spec.yaml"
}

resource "impart_log_binding" "test" {
  name      = "%[1]s-updated"
  spec_id   = resource.impart_spec.test.id
  pattern_type = "grok"
  pattern = <<EOF
%[2]s
  EOF
  logstream_id = "%[1]s-logstream-updated"
}`, name, strings.ReplaceAll(updatedGrokPattern, "%", "%%")),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_log_binding.test", "name", name+"-updated"),
					resource.TestCheckResourceAttr("impart_log_binding.test", "pattern_type", "grok"),
					resource.TestCheckResourceAttr("impart_log_binding.test", "pattern", fmt.Sprintf("%s\n", updatedGrokPattern)),
					resource.TestCheckResourceAttr("impart_log_binding.test", "logstream_id", name+"-logstream-updated"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("impart_log_binding.test", "spec_id"),
					resource.TestCheckResourceAttrSet("impart_log_binding.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
