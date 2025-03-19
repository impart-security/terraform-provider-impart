# Create a new external link
resource "impart_external_link" "external_link1" {
  name              = "Datadog client IP address"
  description       = "A link to Datadog dashboard for client IP address"
  url               = "https://app.datadoghq.com/dashboard/3tm-mpc-863?tpl_var_ClientIp=9.37.130.233"
  entity            = "request"
  json_path_element = "$.client_ip.address"
  vendor            = "Datadog"
}
