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
	_ datasource.DataSource              = &connectorDataSource{}
	_ datasource.DataSourceWithConfigure = &connectorDataSource{}
)

// NewConnectorDataSource is a helper function to simplify the provider implementation.
func NewConnectorDataSource() datasource.DataSource {
	return &connectorDataSource{}
}

// connectorDataSource is the data source implementation.
type connectorDataSource struct {
	client *impartAPIClient
}

// connectorDataSourceModel maps the data source schema data.
type connectorDataSourceModel struct {
	ID              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	ConnectorTypeId types.String `tfsdk:"connector_type_id"`
	IsConnected     types.Bool   `tfsdk:"is_connected"`
}

// Configure adds the provider configured client to the data source.
func (d *connectorDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
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
func (d *connectorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connector"
}

// Schema defines the schema for the data source.
func (d *connectorDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage a connector.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for this connector.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name for this connector.",
				Optional:    true,
			},
			"connector_type_id": schema.StringAttribute{
				Description: "ID of the connector type (eg. ID for our Slack or Jira connector types).",
				Optional:    true,
			},
			"is_connected": schema.BoolAttribute{
				Description: "Whether or not the connector is authenticated via OAuth2.",
				Optional:    true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *connectorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the connector data source")
	var state connectorDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	connectorResponse, _, err := d.client.ConnectorsAPI.GetConnector(ctx, d.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to read the connector",
			message,
		)
		return
	}

	// Map response body to model
	state = connectorDataSourceModel{
		ID:              types.StringValue(connectorResponse.Id),
		Name:            types.StringValue(connectorResponse.Name),
		ConnectorTypeId: types.StringValue(connectorResponse.ConnectorTypeId),
		IsConnected:     types.BoolValue(connectorResponse.IsConnected),
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	tflog.Debug(ctx, "Finished reading the connector data source", map[string]any{"success": true})
}
