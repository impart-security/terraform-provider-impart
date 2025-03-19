package provider

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapiclient "github.com/impart-security/terraform-provider-impart/internal/apiclient"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &impartProvider{}
)

type impartAPIClient struct {
	*openapiclient.APIClient
	OrgID string
}

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &impartProvider{
			version: version,
		}
	}
}

// impartProvider is the provider implementation.
type impartProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// impartProviderModel maps provider schema data to a Go type.
type impartProviderModel struct {
	Token    types.String `tfsdk:"token"`
	Endpoint types.String `tfsdk:"endpoint"`
}

// Metadata returns the provider type name.
func (p *impartProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "impart"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *impartProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"token": schema.StringAttribute{
				Optional:    true,
				Description: "The Impart api token",
			},
			"endpoint": schema.StringAttribute{
				Optional:    true,
				Description: "The Impart api endpoint",
			},
		},
		Blocks:      map[string]schema.Block{},
		Description: "Interface with the Impart service API.",
	}
}

// Configure prepares a Impart API client for data sources and resources.
//
//gocyclo:ignore
func (p *impartProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Impart client")

	// Retrieve provider data from configuration
	var config impartProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}

	token := os.Getenv("IMPART_TOKEN")
	endpoint := os.Getenv("IMPART_ENDPOINT")

	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}

	if !config.Endpoint.IsNull() {
		endpoint = config.Endpoint.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if token == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Missing Impart API token",
			"Set the token value in the configuration or use the IMPART_TOKEN environment variable. ",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating Impart client")

	// Instantiate impart client
	configuration := openapiclient.NewConfiguration()
	configuration.UserAgent = fmt.Sprintf("terraform-provider-impart (%s)", p.version)
	configuration.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))

	if endpoint != "" {
		u, err := url.Parse(endpoint)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to parse Impart API endpoint",
				"An unexpected error occurred when creating the Impart API client. "+
					"If the error is not clear, please contact the provider developers.\n\n"+
					"Endpoint parse error: "+err.Error(),
			)
			return
		}
		configuration.Host = u.Host
		configuration.Scheme = u.Scheme
	}

	apiClient := openapiclient.NewAPIClient(configuration)

	tokenInfo, httpResp, err := apiClient.UserAPI.GetTokenInfo(context.Background()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Impart API Client",
			"An unexpected error occurred when creating the Impart API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Impart Client Error: "+err.Error(),
		)
		return
	}

	if httpResp.StatusCode != 200 {
		resp.Diagnostics.AddError(
			"Unable to Create Impart API Client",
			fmt.Sprintf("Impart API token_info returned unsuccessful status code %d", httpResp.StatusCode),
		)
		return
	}

	// Make the Impart client available during DataSource and Resource
	// type Configure methods.
	impartClient := impartAPIClient{apiClient, tokenInfo.OrgId}

	resp.DataSourceData = &impartClient
	resp.ResourceData = &impartClient

	tflog.Info(ctx, "Configured Impart client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *impartProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewSpecDataSource,
		NewConnectorDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *impartProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSpecResource,
		NewApiBindingResource,
		NewRuleScriptResource,
		NewLogBindingResource,
		NewNotificationTemplateResource,
		NewMonitorResource,
		NewRuleScriptDependenciesResource,
		NewListResource,
		NewRuleTestcaseResource,
		NewLabelResource,
		NewTagMetadataResource,
		NewExternalLinkResource,
	}
}
