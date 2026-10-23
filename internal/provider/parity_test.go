package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
)

// TestOpenTofuParity fails when the provider adopts a capability that OpenTofu
// does not support. Customers run this provider with OpenTofu as well as
// Terraform, so a Terraform-only capability would silently not work for them.
// Read docs/specs/2026-10-19-terraform-provider-opentofu-parity.md in the
// monorepo and get product sign-off before relaxing a check here.
func TestOpenTofuParity(t *testing.T) {
	p := New("test")()

	if _, ok := p.(provider.ProviderWithListResources); ok {
		t.Error("provider implements list resources, which OpenTofu does not support (opentofu/opentofu#3787)")
	}
	if _, ok := p.(provider.ProviderWithActions); ok {
		t.Error("provider implements actions, which OpenTofu does not support (opentofu/opentofu#3309)")
	}
	if _, ok := p.(provider.ProviderWithStateStores); ok {
		t.Error("provider implements state stores, which OpenTofu does not support")
	}
}

func TestUsingOpenTofu(t *testing.T) {
	for path, want := range map[string]bool{
		"":                           false,
		"/usr/local/bin/terraform":   false,
		"/usr/local/bin/tofu":        true,
		"/opt/hostedtoolcache/tofu":  true,
		"tofu.exe":                   true, // native Windows binary name
		"/home/runner/bin/tofu-1.13": true,
	} {
		t.Setenv("TF_ACC_TERRAFORM_PATH", path)
		if got := usingOpenTofu(); got != want {
			t.Errorf("usingOpenTofu() with TF_ACC_TERRAFORM_PATH=%q = %t, want %t", path, got, want)
		}
	}
}
