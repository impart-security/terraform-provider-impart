package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapiclient "github.com/impart-security/terraform-provider-impart/internal/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &specDataSource{}
	_ datasource.DataSourceWithConfigure = &specDataSource{}
)

// NewSpecDataSource is a helper function to simplify the provider implementation.
func NewSpecDataSource() datasource.DataSource {
	return &specDataSource{}
}

// specDataSource is the data source implementation.
type specDataSource struct {
	client *impartAPIClient
}

// specDataSourceModel maps the data source schema data.
type specDataSourceModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// Configure adds the provider configured client to the data source.
func (d *specDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*impartAPIClient)
	if !ok {
		tflog.Error(ctx, "Unable to prepare client")
		return
	}
	d.client = client
}

// Metadata returns the data source type name.
func (d *specDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_spec"
}

// Schema defines the schema for the data source.
func (d *specDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage a specification.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for this specification.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name for this specification.",
				Computed:    true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *specDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read spec data source")
	var state specDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	specResponse, _, err := d.client.SpecsApi.GetSpec(ctx, d.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to read specification",
			message,
		)
		return
	}

	// Map response body to model
	state = specDataSourceModel{
		ID:   types.StringValue(specResponse.Id),
		Name: types.StringValue(specResponse.Name),
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	tflog.Debug(ctx, "Finished reading spec data source", map[string]any{"success": true})
}
