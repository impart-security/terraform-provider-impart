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
	_ resource.Resource                = &apiBindingResource{}
	_ resource.ResourceWithConfigure   = &apiBindingResource{}
	_ resource.ResourceWithImportState = &apiBindingResource{}
)

// NewApiBindingResource is a helper function to simplify the provider implementation.
func NewApiBindingResource() resource.Resource {
	return &apiBindingResource{}
}

// apiBindingResource is the resource implementation.
type apiBindingResource struct {
	client *impartAPIClient
}

// apiBindingResourceModel maps the resource schema data.
type apiBindingResourceModel struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	SpecID         types.String `tfsdk:"spec_id"`
	Hostname       types.String `tfsdk:"hostname"`
	Port           types.Int64  `tfsdk:"port"`
	BasePath       types.String `tfsdk:"base_path"`
	Disabled       types.Bool   `tfsdk:"disabled"`
	UpstreamOrigin types.String `tfsdk:"upstream_origin"`
	Hops           types.Int64  `tfsdk:"hops"`
	UseForwarded   types.Bool   `tfsdk:"use_forwarded"`
	ForwardedFor   []string     `tfsdk:"forwarded_for"`
	ForwardedHost  []string     `tfsdk:"forwarded_host"`
	ForwardedProto []string     `tfsdk:"forwarded_proto"`
	ForwardedID    []string     `tfsdk:"forwarded_id"`
}

// Configure adds the provider configured client to the resource.
func (r *apiBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*impartAPIClient)
	if !ok {
		tflog.Error(ctx, "Unable to prepare the client")
		return
	}
	r.client = client
}

// Metadata returns the resource type name.
func (r *apiBindingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_binding"
}

// Schema defines the schema for the resource.
func (r *apiBindingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage an api binding.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for this api binding.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name for this api binding.",
				Required:    true,
			},
			"spec_id": schema.StringAttribute{
				Description: "The specification id.",
				Required:    true,
			},
			"hostname": schema.StringAttribute{
				Description: "The hostname for this api binding.",
				Required:    true,
			},
			"port": schema.Int64Attribute{
				Description: "The port for this api binding.",
				Required:    true,
			},
			"base_path": schema.StringAttribute{
				Description: "The base_path for this api binding.",
				Required:    true,
			},
			"upstream_origin": schema.StringAttribute{
				Description: "The upstream_origin for this api binding.",
				Optional:    true,
			},
			"disabled": schema.BoolAttribute{
				Description: "The disabled for this api binding.",
				Optional:    true,
			},
			"hops": schema.Int64Attribute{
				Description: "The hops for this api binding.",
				Optional:    true,
			},
			"use_forwarded": schema.BoolAttribute{
				Description: "The use_forwarded for this api binding.",
				Optional:    true,
			},
			"forwarded_for": schema.ListAttribute{
				Description: "The forwarded_for for this api binding.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"forwarded_host": schema.ListAttribute{
				Description: "The forwarded_host for this api binding.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"forwarded_proto": schema.ListAttribute{
				Description: "The forwarded_proto for this api binding.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"forwarded_id": schema.ListAttribute{
				Description: "The forwarded_id for this api binding.",
				ElementType: types.StringType,
				Optional:    true,
			},
		},
	}
}

func (r *apiBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Create a new resource.
func (r *apiBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the api binding resource")
	// Retrieve values from plan
	var plan apiBindingResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	upstreamOrigin := plan.UpstreamOrigin.ValueString()

	postBody := openapiclient.ApiBindingPostBody{
		Name:           plan.Name.ValueString(),
		SpecId:         plan.SpecID.ValueString(),
		Hostname:       plan.Hostname.ValueString(),
		Port:           int32(plan.Port.ValueInt64()),
		BasePath:       plan.BasePath.ValueString(),
		UpstreamOrigin: &upstreamOrigin,
		ForwardedFor:   plan.ForwardedFor,
		ForwardedHost:  plan.ForwardedHost,
		ForwardedId:    plan.ForwardedID,
		ForwardedProto: plan.ForwardedProto,
	}

	if !plan.Disabled.IsNull() {
		disabled := plan.Disabled.ValueBool()
		postBody.Disabled = &disabled
	}

	if !plan.Hops.IsNull() {
		hops := int32(plan.Hops.ValueInt64())
		postBody.Hops = &hops
	}

	if !plan.UseForwarded.IsNull() {
		useForwarded := plan.UseForwarded.IsNull()
		postBody.UseForwarded = &useForwarded
	}

	bindingResponse, _, err := r.client.ApiBindingsAPI.CreateAPIBinding(ctx, r.client.OrgID).ApiBindingPostBody(postBody).Execute()

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the api binding",
			message,
		)
		return
	}

	// Map response body to model
	plan.ID = types.StringValue(bindingResponse.Id)
	plan.Name = types.StringValue(bindingResponse.Name)
	plan.BasePath = types.StringValue(bindingResponse.BasePath)
	plan.Hostname = types.StringValue(bindingResponse.Hostname)
	plan.SpecID = types.StringValue(bindingResponse.SpecId)
	plan.Port = types.Int64Value(int64(bindingResponse.Port))
	plan.ForwardedFor = bindingResponse.ForwardedFor
	plan.ForwardedHost = bindingResponse.ForwardedHost
	plan.ForwardedID = bindingResponse.ForwardedId
	plan.ForwardedProto = bindingResponse.ForwardedProto

	if !plan.Disabled.IsNull() || bindingResponse.Disabled {
		plan.Disabled = types.BoolValue(bindingResponse.Disabled)
	}
	if bindingResponse.UpstreamOrigin != "" {
		plan.UpstreamOrigin = types.StringValue(bindingResponse.UpstreamOrigin)
	}
	if !(plan.Hops.IsNull() && bindingResponse.Hops == 0) {
		plan.Hops = types.Int64Value(int64(bindingResponse.Hops))
	}
	if !plan.UseForwarded.IsNull() || bindingResponse.UseForwarded {
		plan.UseForwarded = types.BoolValue(bindingResponse.UseForwarded)
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Created the binding resource", map[string]any{"success": true})
}

// Read resource information.
func (r *apiBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the api binding resource")
	// Get current state
	var state apiBindingResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	bindingResponse, httpResp, err := r.client.ApiBindingsAPI.GetAPIBinding(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
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
			"Unable to read the api binding",
			message,
		)
		return
	}

	// Map response body to model
	state.ID = types.StringValue(bindingResponse.Id)
	state.Name = types.StringValue(bindingResponse.Name)
	state.BasePath = types.StringValue(bindingResponse.BasePath)
	state.Hostname = types.StringValue(bindingResponse.Hostname)
	state.SpecID = types.StringValue(bindingResponse.SpecId)
	state.Port = types.Int64Value(int64(bindingResponse.Port))
	state.ForwardedFor = bindingResponse.ForwardedFor
	state.ForwardedHost = bindingResponse.ForwardedHost
	state.ForwardedID = bindingResponse.ForwardedId
	state.ForwardedProto = bindingResponse.ForwardedProto

	if !state.Disabled.IsNull() || bindingResponse.Disabled {
		state.Disabled = types.BoolValue(bindingResponse.Disabled)
	}
	if bindingResponse.UpstreamOrigin != "" {
		state.UpstreamOrigin = types.StringValue(bindingResponse.UpstreamOrigin)
	}
	if !state.Hops.IsNull() || bindingResponse.Hops != 0 {
		state.Hops = types.Int64Value(int64(bindingResponse.Hops))
	}
	if !state.UseForwarded.IsNull() || bindingResponse.UseForwarded {
		state.UseForwarded = types.BoolValue(bindingResponse.UseForwarded)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Finished reading the api binding resource", map[string]any{"success": true})
}

func (r *apiBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the api binding resource")
	// Retrieve values from plan
	var plan apiBindingResourceModel
	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	upstreamOrigin := plan.UpstreamOrigin.ValueString()
	postBody := openapiclient.ApiBindingPostBody{
		Name:           plan.Name.ValueString(),
		SpecId:         plan.SpecID.ValueString(),
		Hostname:       plan.Hostname.ValueString(),
		Port:           int32(plan.Port.ValueInt64()),
		BasePath:       plan.BasePath.ValueString(),
		UpstreamOrigin: &upstreamOrigin,
		ForwardedFor:   plan.ForwardedFor,
		ForwardedHost:  plan.ForwardedHost,
		ForwardedId:    plan.ForwardedID,
		ForwardedProto: plan.ForwardedProto,
	}

	if !plan.Disabled.IsNull() {
		disabled := plan.Disabled.ValueBool()
		postBody.Disabled = &disabled
	}

	if !plan.Hops.IsNull() {
		hops := int32(plan.Hops.ValueInt64())
		postBody.Hops = &hops
	}

	if !plan.UseForwarded.IsNull() {
		useForwarded := plan.UseForwarded.ValueBool()
		postBody.UseForwarded = &useForwarded
	}

	bindingReauest := r.client.ApiBindingsAPI.UpdateAPIBinding(ctx, r.client.OrgID, plan.ID.ValueString()).
		ApiBindingPostBody(postBody)

	// Update the api binding
	bindingResponse, _, err := bindingReauest.Execute()

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the api binding",
			message,
		)
		return
	}

	// Overwrite the api binding with refreshed state
	state := apiBindingResourceModel{
		ID:             types.StringValue(bindingResponse.Id),
		Name:           types.StringValue(bindingResponse.Name),
		BasePath:       types.StringValue(bindingResponse.BasePath),
		Hostname:       types.StringValue(bindingResponse.Hostname),
		SpecID:         types.StringValue(bindingResponse.SpecId),
		Port:           types.Int64Value(int64(bindingResponse.Port)),
		ForwardedFor:   bindingResponse.ForwardedFor,
		ForwardedHost:  bindingResponse.ForwardedHost,
		ForwardedID:    bindingResponse.ForwardedId,
		ForwardedProto: bindingResponse.ForwardedProto,
	}

	if !plan.Disabled.IsNull() || bindingResponse.Disabled {
		state.Disabled = types.BoolValue(bindingResponse.Disabled)
	}
	if !plan.UpstreamOrigin.IsNull() || bindingResponse.UpstreamOrigin != "" {
		state.UpstreamOrigin = types.StringValue(bindingResponse.UpstreamOrigin)
	}
	if !plan.Hops.IsNull() || bindingResponse.Hops != 0 {
		state.Hops = types.Int64Value(int64(bindingResponse.Hops))
	}
	if !plan.UseForwarded.IsNull() || bindingResponse.UseForwarded {
		state.UseForwarded = types.BoolValue(bindingResponse.UseForwarded)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updated the api binding resource", map[string]any{"success": true})
}

func (r *apiBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the api binding resource")
	// Retrieve values from state
	var state apiBindingResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete the api binding
	_, err := r.client.ApiBindingsAPI.DeleteAPIBinding(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the api binding",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the api binding resource", map[string]any{"success": true})
}
