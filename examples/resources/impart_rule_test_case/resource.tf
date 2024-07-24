resource "impart_rule_test_case" "example" {
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
      }
      res = {
        status_code   = 201
        header_keys   = ["Header1"]
        header_values = ["value1"]
        body          = base64encode("response body")
      }
    }
  ]
}
