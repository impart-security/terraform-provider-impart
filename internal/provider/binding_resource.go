package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapiclient "github.com/impart-security/terraform-provider-impart/internal/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &bindingResource{}
	_ resource.ResourceWithConfigure   = &bindingResource{}
	_ resource.ResourceWithImportState = &bindingResource{}
)

// NewBindingResource is a helper function to simplify the provider implementation.
func NewBindingResource() resource.Resource {
	return &bindingResource{}
}

// bindingResource is the resource implementation.
type bindingResource struct {
	client *impartAPIClient
}

// bindingResourceModel maps the resource schema data.
type bindingResourceModel struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	SpecId         types.String `tfsdk:"spec_id"`
	Hostname       types.String `tfsdk:"hostname"`
	Port           types.Int64  `tfsdk:"port"`
	BasePath       types.String `tfsdk:"base_path"`
	UpstreamOrigin types.String `tfsdk:"upstream_origin"`
}

// Configure adds the provider configured client to the resource.
func (r *bindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*impartAPIClient)
	if !ok {
		tflog.Error(ctx, "Unable to prepare client")
		return
	}
	r.client = client
}

// Metadata returns the resource type name.
func (r *bindingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_binding"
}

// Schema defines the schema for the resource.
func (r *bindingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage a binding.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for this binding.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name for this binding.",
				Required:    true,
			},
			"spec_id": schema.StringAttribute{
				Description: "The specification id.",
				Required:    true,
			},
			"hostname": schema.StringAttribute{
				Description: "The hostname for this binding.",
				Required:    true,
			},
			"port": schema.Int64Attribute{
				Description: "The port for this binding.",
				Required:    true,
			},
			"base_path": schema.StringAttribute{
				Description: "The base_path for this binding.",
				Required:    true,
			},
			"upstream_origin": schema.StringAttribute{
				Description: "The upstream_origin for this binding.",
				Optional:    true,
			},
		},
	}
}

func (r *bindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Create a new resource.
func (r *bindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create specification resource")
	// Retrieve values from plan
	var plan bindingResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	upstreamOrigin := plan.UpstreamOrigin.ValueString()

	bindingResponse, _, err := r.client.ApiBindingsApi.CreateAPIBinding(ctx, r.client.OrgID).ApiBindingPostBody(openapiclient.ApiBindingPostBody{
		Name:           plan.Name.ValueString(),
		SpecId:         plan.SpecId.ValueString(),
		Hostname:       plan.Hostname.ValueString(),
		Port:           int32(plan.Port.ValueInt64()),
		BasePath:       plan.BasePath.ValueString(),
		UpstreamOrigin: &upstreamOrigin,
	}).Execute()

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create binding",
			message,
		)
		return
	}

	// Map response body to model
	plan.ID = types.StringValue(bindingResponse.Id)
	plan.Name = types.StringValue(bindingResponse.Name)
	plan.BasePath = types.StringValue(bindingResponse.BasePath)
	plan.Hostname = types.StringValue(bindingResponse.Hostname)
	plan.SpecId = types.StringValue(bindingResponse.SpecId)
	plan.Port = types.Int64Value(int64(bindingResponse.Port))

	if bindingResponse.UpstreamOrigin != "" {
		plan.UpstreamOrigin = types.StringValue(bindingResponse.UpstreamOrigin)
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Created binding resource", map[string]any{"success": true})
}

// Read resource information.
func (r *bindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read binding resource")
	// Get current state
	var state bindingResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	bindingResponse, httpResp, err := r.client.ApiBindingsApi.GetAPIBinding(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		// Treat HTTP 404 Not Found status as a signal to remove/recreate resource
		if httpResp.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}

		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to read binding",
			message,
		)
		return
	}

	// Map response body to model
	state = bindingResourceModel{
		ID:       types.StringValue(bindingResponse.Id),
		Name:     types.StringValue(bindingResponse.Name),
		BasePath: types.StringValue(bindingResponse.BasePath),
		Hostname: types.StringValue(bindingResponse.Hostname),
		SpecId:   types.StringValue(bindingResponse.SpecId),
		Port:     types.Int64Value(int64(bindingResponse.Port)),
	}

	if bindingResponse.UpstreamOrigin != "" {
		state.UpstreamOrigin = types.StringValue(bindingResponse.UpstreamOrigin)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Finished reading binding resource", map[string]any{"success": true})
}

func (r *bindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update binding resource")
	// Retrieve values from plan
	var plan bindingResourceModel
	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	upstreamOrigin := plan.UpstreamOrigin.ValueString()

	specRequest := r.client.ApiBindingsApi.UpdateAPIBinding(ctx, r.client.OrgID, plan.ID.ValueString()).
		ApiBindingPostBody(openapiclient.ApiBindingPostBody{
			Name:           plan.Name.ValueString(),
			SpecId:         plan.SpecId.ValueString(),
			Hostname:       plan.Hostname.ValueString(),
			Port:           int32(plan.Port.ValueInt64()),
			BasePath:       plan.BasePath.ValueString(),
			UpstreamOrigin: &upstreamOrigin,
		})

	// update specification
	bindingResponse, _, err := specRequest.Execute()

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update binding",
			message,
		)
		return
	}

	// Overwrite specifications with refreshed state
	state := bindingResourceModel{
		ID:       types.StringValue(bindingResponse.Id),
		Name:     types.StringValue(bindingResponse.Name),
		BasePath: types.StringValue(bindingResponse.BasePath),
		Hostname: types.StringValue(bindingResponse.Hostname),
		SpecId:   types.StringValue(bindingResponse.SpecId),
		Port:     types.Int64Value(int64(bindingResponse.Port)),
	}

	if !plan.UpstreamOrigin.IsNull() || bindingResponse.UpstreamOrigin != "" {
		state.UpstreamOrigin = types.StringValue(bindingResponse.UpstreamOrigin)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updated binding resource", map[string]any{"success": true})
}

func (r *bindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete binding resource")
	// Retrieve values from state
	var state bindingResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// delete specification
	_, err := r.client.ApiBindingsApi.DeleteAPIBinding(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete binding",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted binding resource", map[string]any{"success": true})
}
