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

  assertions = [
    {
      message_indexes = [0]
      assertion_type  = "output"
      condition       = "contains" #contains|not_contains
      expected        = "test"
    },
    {
      message_indexes = [0]
      assertion_type  = "block"
      location        = "req" # req|res
      expected        = "true"
    },
    {
      message_indexes = [0]
      assertion_type  = "status_code"
      location        = "req"    # req|res
      condition       = "one_of" # equal|not_equal|greater_than|less_than|one_of
      expected        = "201,200"
    },
    {
      message_indexes = [0]
      assertion_type  = "tags"
      location        = "req"      # req|res
      condition       = "contains" # contains|not_contains
      expected        = "tagname"
    }
  ]
}
