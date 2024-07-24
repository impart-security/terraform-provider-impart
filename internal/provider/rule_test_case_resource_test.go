package provider

import (
	"encoding/base64"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestRuleTestCaseResource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "impart_rule_test_case" "test" {
  name        = "example"
  description = "test case description"
  messages = [
    {
      count = 1
      req = {
        url           = "http://example.com"
        method        = "GET"
        body          = base64encode("request body")
        header_keys   = ["Header1"]
        header_values = ["value1"]
        remote_addr   = "192.168.1.1"
		truncated_body = true
      }
      res = {
        status_code   = 201
        header_keys   = ["Header1"]
        header_values = ["value1"]
        body          = base64encode("response body")
		truncated_body = false
      }
    }
  ]
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "name", "example"),
					resource.TestCheckResourceAttrSet("impart_rule_test_case.test", "id"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "description", "test case description"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.#", "1"),

					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.count", "1"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.url", "http://example.com"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.method", "GET"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.body", base64.StdEncoding.EncodeToString([]byte("request body"))),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.header_keys.0", "Header1"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.header_values.0", "value1"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.remote_addr", "192.168.1.1"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.truncated_body", "true"),

					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.res.status_code", "201"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.res.body", base64.StdEncoding.EncodeToString([]byte("response body"))),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.res.header_keys.0", "Header1"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.res.header_values.0", "value1"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.res.truncated_body", "false"),
				),
			},
			// Update
			{
				Config: providerConfig + `
resource "impart_rule_test_case" "test" {
  name        = "example"
  description = "test case description"
  messages = [
    {
      count = 3
	  delay = 1000
	  post_delay = 2000
      req = {
        url           = "http://example2.com"
        method        = "POST"
        body          = base64encode("request body updated")
        header_keys   = ["Header2"]
        header_values = ["value2"]
        remote_addr   = "192.168.1.2"
		truncated_body = false
      }
      res = {
        status_code   = 200
        header_keys   = ["Header2"]
        header_values = ["value2"]
        body          = base64encode("response body updated")
		truncated_body = true
      }
    }
  ]
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "name", "example"),
					resource.TestCheckResourceAttrSet("impart_rule_test_case.test", "id"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "description", "test case description"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.#", "1"),

					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.count", "3"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.delay", "1000"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.post_delay", "2000"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.url", "http://example2.com"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.method", "POST"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.body", base64.StdEncoding.EncodeToString([]byte("request body updated"))),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.header_keys.0", "Header2"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.header_values.0", "value2"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.remote_addr", "192.168.1.2"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.req.truncated_body", "false"),

					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.res.status_code", "200"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.res.body", base64.StdEncoding.EncodeToString([]byte("response body updated"))),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.res.header_keys.0", "Header2"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.res.header_values.0", "value2"),
					resource.TestCheckResourceAttr("impart_rule_test_case.test", "messages.0.res.truncated_body", "true"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}
