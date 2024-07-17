package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLogBindingResource(t *testing.T) {

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
  name        = "terraform_test"
  source_file = "./testdata/spec.yaml"
}

resource "impart_log_binding" "test" {
  name         = "terraform_test"
  spec_id      = resource.impart_spec.test.id
  pattern_type = "grok"
  pattern = <<EOF
%s
  EOF
  logstream_id = "logstream"
}`, strings.Replace(grokPattern, "%", "%%", -1)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_log_binding.test", "name", "terraform_test"),
					resource.TestCheckResourceAttr("impart_log_binding.test", "logstream_id", "logstream"),
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
  name        = "terraform_test"
  source_file = "./testdata/spec.yaml"
}
resource "impart_log_binding" "test" {
  name      = "terraform_test_updated"
  spec_id   = resource.impart_spec.test.id
  pattern_type = "grok"
  pattern = <<EOF
%s
  EOF
  logstream_id = "updated_logstream"
}`, strings.Replace(updatedGrokPattern, "%", "%%", -1)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_log_binding.test", "name", "terraform_test_updated"),
					resource.TestCheckResourceAttr("impart_log_binding.test", "pattern_type", "grok"),
					resource.TestCheckResourceAttr("impart_log_binding.test", "pattern", fmt.Sprintf("%s\n", updatedGrokPattern)),
					resource.TestCheckResourceAttr("impart_log_binding.test", "logstream_id", "updated_logstream"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("impart_log_binding.test", "spec_id"),
					resource.TestCheckResourceAttrSet("impart_log_binding.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
