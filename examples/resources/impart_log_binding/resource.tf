# Create a new log binding
resource "impart_log_binding" "example" {
  name         = "log_binding_example"
  pattern_type = "grok"
  # Example patterns
  # for api gateway log format: $context.requestTime "$context.httpMethod $context.path $context.protocol" $context.status $context.identity.sourceIp $context.requestId
  # %%{HTTPDATE:timestamp} "(?:%%{WORD:http_method}|-) (?:%%{GREEDYDATA:request}|-) (?:HTTP/%%{NUMBER:httpversion}|-( )?)" (?:%%{NUMBER:response_code}|-)
  # for aws loadbalancer access logs
  # %%{TIMESTAMP_ISO8601:timestamp} %%{NOTSPACE:loadbalancer} %%{IP:client_ip}:%{NUMBER:client_port} (?:%{IP:backend_ip}:%{NUMBER:backend_port}|-) %%{NUMBER:request_processing_time} %%{NUMBER:backend_processing_time} %%{NUMBER:response_processing_time} (?:%{NUMBER:response_code}|-) (?:%{NUMBER:backend_status_code}|-) %%{NUMBER:received_bytes} %%{NUMBER:sent_bytes} "(?:%{WORD:http_method}|-) (?:%{GREEDYDATA:request}|-) (?:HTTP/%{NUMBER:http_version}|-( )?)" "%{DATA:user_agent}"( %%{NOTSPACE:ssl_cipher} %%{NOTSPACE:ssl_protocol})?
  pattern      = <<EOF
<pattern>
  EOF
  logstream_id = "logstream_id"
  spec_id      = resource.impart_spec.example.id
}
