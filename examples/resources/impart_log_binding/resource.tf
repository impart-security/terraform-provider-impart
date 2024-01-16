# Create a new log binding
resource "impart_log_binding" "example" {
  name         = "log_binding_example"
  grok_pattern = <<EOF
%%{HTTPDATE:timestamp} "(?:%%{WORD:http_method}|-) (?:%%{GREEDYDATA:request}|-) (?:HTTP/%%{NUMBER:httpversion}|-( )?)" (?:%%{NUMBER:response_code}|-)
  EOF
  logstream_id = "logstream_id"
}
