package provider

import (
	"os"
	"path/filepath"
	"strings"
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

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){"impart": providerserver.NewProtocol6WithError(New("test")())}

// testAccPreCheck validates the necessary test API key exists
// in the testing environment.
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("IMPART_TOKEN"); v == "" {
		t.Fatal("IMPART_TOKEN must be set for acceptance tests")
	}
	t.Logf("acceptance CLI: %q (opentofu=%t)", os.Getenv("TF_ACC_TERRAFORM_PATH"), usingOpenTofu())
}

// usingOpenTofu reports whether acceptance tests are driving the OpenTofu CLI.
// tofu reports its version as "terraform_version", so tfversion checks can't tell
// the CLIs apart; gate CLI-specific test steps on this instead (for example
// ExpectNonEmptyPlan: usingOpenTofu() for ephemeral resources, opentofu/opentofu#3781).
func usingOpenTofu() bool {
	return strings.HasPrefix(filepath.Base(os.Getenv("TF_ACC_TERRAFORM_PATH")), "tofu")
}
