package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	// providerConfig is a shared configuration to combine with the actual
	// test configuration so the Impart client is properly configured.
	providerConfig = `terraform {
  required_providers {
    impart = {
      source = "impart-security/impart"
    }
  }
}

# Configure the connection details for the Impart service
provider "impart" {

}
`
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){"impart": providerserver.NewProtocol6WithError(New("test")())}
)

// testAccPreCheck validates the necessary test API key exists
// in the testing environment.
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("IMPART_TOKEN"); v == "" {
		t.Fatal("IMPART_TOKEN must be set for acceptance tests")
	}
}
