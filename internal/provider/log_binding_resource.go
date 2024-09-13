package provider

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapiclient "github.com/impart-security/terraform-provider-impart/internal/apiclient"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &logBindingResource{}
	_ resource.ResourceWithConfigure   = &logBindingResource{}
	_ resource.ResourceWithImportState = &logBindingResource{}
)

// NewLogBindingResource creates log binding resource.
func NewLogBindingResource() resource.Resource {
	return &logBindingResource{}
}

// logBindingResource is the resource implementation for log bindins.
type logBindingResource struct {
	client *impartAPIClient
}

// logBindingResourceModel maps the resource schema data.
type logBindingResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	SpecID      types.String `tfsdk:"spec_id"`
	LogstreamID types.String `tfsdk:"logstream_id"`
	Pattern     types.String `tfsdk:"pattern"`
	PatternType types.String `tfsdk:"pattern_type"`
}

// Configure adds the provider configured client to the resource.
func (r *logBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *logBindingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_log_binding"
}

// Schema defines the schema for the resource.
func (r *logBindingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage a log binding.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for this log binding.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name for this log binding.",
				Required:    true,
			},
			"spec_id": schema.StringAttribute{
				Description: "The specification id.",
				Required:    true,
			},
			"pattern": schema.StringAttribute{
				Description: "The grok/json pattern for this log binding.",
				Required:    true,
			},
			"pattern_type": schema.StringAttribute{
				Description: "The pattern type for this log binding. Accepted values: grok, json",
				Required:    true,
			},
			"logstream_id": schema.StringAttribute{
				Description: "The logstream id for this log binding.",
				Optional:    true,
			},
		},
	}
}

func (r *logBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Create a new resource.
func (r *logBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the log binding resource")
	// Retrieve values from plan
	var plan logBindingResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	logstreamID := plan.LogstreamID.ValueString()

	bindingResponse, _, err := r.client.LogBindingsAPI.CreateLogBinding(ctx, r.client.OrgID).LogBindingPostBody(openapiclient.LogBindingPostBody{
		Name:        plan.Name.ValueString(),
		SpecId:      plan.SpecID.ValueString(),
		Pattern:     base64.StdEncoding.EncodeToString([]byte(plan.Pattern.ValueString())),
		PatternType: plan.PatternType.ValueString(),
		LogstreamId: &logstreamID,
	}).Execute()

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the log binding",
			message,
		)
		return
	}

	// Map response body to model
	plan.ID = types.StringValue(bindingResponse.Id)
	plan.Name = types.StringValue(bindingResponse.Name)
	plan.SpecID = types.StringValue(bindingResponse.SpecId)

	pattern, err := base64.StdEncoding.DecodeString(bindingResponse.Pattern)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to base64 decode the pattern",
			err.Error(),
		)
	}
	plan.Pattern = types.StringValue(string(pattern))

	if logstreamID != "" {
		plan.LogstreamID = types.StringValue(bindingResponse.LogstreamId)
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Created the log binding resource", map[string]any{"success": true})
}

// Read resource information.
func (r *logBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the log binding resource")
	// Get current state
	var state logBindingResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	logstreamID := state.LogstreamID.ValueString()

	bindingResponse, httpResp, err := r.client.LogBindingsAPI.GetLogBinding(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
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
			"Unable to read the log binding",
			message,
		)
		return
	}

	// Map response body to model
	state = logBindingResourceModel{
		ID:          types.StringValue(bindingResponse.Id),
		Name:        types.StringValue(bindingResponse.Name),
		SpecID:      types.StringValue(bindingResponse.SpecId),
		PatternType: types.StringValue(bindingResponse.PatternType),
	}

	if logstreamID != "" {
		state.LogstreamID = types.StringValue(bindingResponse.LogstreamId)
	}

	pattern, err := base64.StdEncoding.DecodeString(bindingResponse.Pattern)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to base64 decode the pattern",
			err.Error(),
		)
	}
	state.Pattern = types.StringValue(string(pattern))

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Finished reading the log binding resource", map[string]any{"success": true})
}

func (r *logBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the log binding resource")
	// Retrieve values from plan
	var plan logBindingResourceModel
	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	logstreamId := plan.ID.ValueString()
	if !plan.LogstreamID.IsNull() {
		logstreamId = plan.LogstreamID.ValueString()
	}

	specRequest := r.client.LogBindingsAPI.UpdateLogBinding(ctx, r.client.OrgID, plan.ID.ValueString()).
		LogBindingPutBody(openapiclient.LogBindingPutBody{
			Name:        plan.Name.ValueString(),
			SpecId:      plan.SpecID.ValueString(),
			Pattern:     base64.StdEncoding.EncodeToString([]byte(plan.Pattern.ValueString())),
			LogstreamId: logstreamId,
			PatternType: plan.PatternType.ValueString(),
		})

	// update the log binding
	logBindingResponse, _, err := specRequest.Execute()

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the log binding",
			message,
		)
		return
	}

	// Overwrite the log binding with refreshed state
	state := logBindingResourceModel{
		ID:          types.StringValue(logBindingResponse.Id),
		Name:        types.StringValue(logBindingResponse.Name),
		Pattern:     types.StringValue(logBindingResponse.Pattern),
		SpecID:      types.StringValue(logBindingResponse.SpecId),
		PatternType: types.StringValue(logBindingResponse.PatternType),
	}

	pattern, err := base64.StdEncoding.DecodeString(logBindingResponse.Pattern)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to base64 decode the pattern",
			err.Error(),
		)
	}
	state.Pattern = types.StringValue(string(pattern))

	// update state if logstream id was set
	if !plan.LogstreamID.IsNull() {
		state.LogstreamID = types.StringValue(logBindingResponse.LogstreamId)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updated the log binding resource", map[string]any{"success": true})
}

func (r *logBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the log binding resource")
	// Retrieve values from state
	var state logBindingResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// delete the log binding
	_, err := r.client.LogBindingsAPI.DeleteLogBinding(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the log binding",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the log binding resource", map[string]any{"success": true})
}

func (r logBindingResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data logBindingResourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.PatternType.ValueString() != "json" && data.PatternType.ValueString() != "grok" {
		resp.Diagnostics.AddError(
			"Invalid pattern_type",
			"Accepted values: grok, json",
		)
	}
}
