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

	openapiclient "github.com/impart-security/terraform-provider-impart/internal/apiclient"
)

var (
	_ resource.Resource                = &ruleClientIdentifierStorageResource{}
	_ resource.ResourceWithConfigure   = &ruleClientIdentifierStorageResource{}
	_ resource.ResourceWithImportState = &ruleClientIdentifierStorageResource{}
)

// NewRuleClientIdentifierStorageResource is a helper function to simplify the provider implementation.
func NewRuleClientIdentifierStorageResource() resource.Resource {
	return &ruleClientIdentifierStorageResource{}
}

// ruleClientIdentifierStorageResource is the resource implementation.
type ruleClientIdentifierStorageResource struct {
	client *impartAPIClient
}

// Configure adds the provider configured client to the resource.
func (r *ruleClientIdentifierStorageResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *ruleClientIdentifierStorageResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rule_client_identifier_storage"
}

type ruleClientIdentifierStorageModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Capacity    types.Int32  `tfsdk:"capacity"`
}

func (r ruleClientIdentifierStorageResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for the rule client identifier storage.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name for the rule client identifier storage.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description for the rule client identifier storage.",
				Optional:    true,
			},
			"capacity": schema.Int32Attribute{
				Description: "The capacity for the rule client identifier storage.",
				Required:    true,
			},
		},
	}
}

func (r *ruleClientIdentifierStorageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r ruleClientIdentifierStorageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the client identifier storage resource")

	var plan ruleClientIdentifierStorageModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody, err := toClientIdentifierStoragePostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create the client identifier storage resource",
			err.Error(),
		)
		return
	}

	clientIdentifierStorageRequest := r.client.RuleClientIdentifierStoragesAPI.
		CreateRuleClientIdentifierStorage(ctx, r.client.OrgID).
		RuleClientIdentifierStoragePostBody(postBody)

	clientIdentifierStorageResponse, _, err := clientIdentifierStorageRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the client identifier storage resource",
			message,
		)
		return
	}

	state := toRuleClientIdentifierStorageModel(clientIdentifierStorageResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Created the client identifier storage resource", map[string]any{"success": true})
}

func (r ruleClientIdentifierStorageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the client identifier storage resource")

	var data ruleClientIdentifierStorageModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clientIdentifierStorageResponse, httpResp, err := r.client.RuleClientIdentifierStoragesAPI.
		GetRuleClientIdentifierStorage(ctx, r.client.OrgID, data.ID.ValueString()).
		Execute()

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
			"Unable to read the client identifier storage resource "+data.ID.ValueString(),
			message,
		)
		return
	}

	state := toRuleClientIdentifierStorageModel(clientIdentifierStorageResponse, data)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Finished reading the client identifier storage resource", map[string]any{"success": true})
}

func (r ruleClientIdentifierStorageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the client identifier storage resource")

	var plan ruleClientIdentifierStorageModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody, err := toClientIdentifierStoragePostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update the client identifier storage resource",
			err.Error(),
		)
		return
	}

	clientIdentifierStorageRequest := r.client.RuleClientIdentifierStoragesAPI.
		UpdateRuleClientIdentifierStorage(ctx, r.client.OrgID, plan.ID.ValueString()).
		RuleClientIdentifierStoragePostBody(postBody)

	clientIdentifierStorageResponse, _, err := clientIdentifierStorageRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the client identifier storage resource",
			message,
		)
		return
	}

	state := toRuleClientIdentifierStorageModel(clientIdentifierStorageResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updated the client identifier storage resource", map[string]any{"success": true})
}

func (r ruleClientIdentifierStorageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the client identifier storage resource")
	// Retrieve values from a state
	var state ruleClientIdentifierStorageModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete a client identifier storage
	_, err := r.client.RuleClientIdentifierStoragesAPI.DeleteRuleClientIdentifierStorage(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the client identifier storage resource",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the client identifier storage resource", map[string]any{"success": true})
}

func toClientIdentifierStoragePostBody(plan ruleClientIdentifierStorageModel) (openapiclient.RuleClientIdentifierStoragePostBody, error) {
	postBody := openapiclient.RuleClientIdentifierStoragePostBody{
		Name:     plan.Name.ValueString(),
		Capacity: int32(plan.Capacity.ValueInt32()),
	}

	if !plan.Description.IsNull() {
		description := plan.Description.ValueString()
		postBody.Description = &description
	}

	return postBody, nil
}

func toRuleClientIdentifierStorageModel(clientIdentifierResponse *openapiclient.RuleClientIdentifierStorage, plan ruleClientIdentifierStorageModel) ruleClientIdentifierStorageModel {
	ruleClientIdentifier := ruleClientIdentifierStorageModel{
		ID:       types.StringValue(clientIdentifierResponse.Id),
		Name:     types.StringValue(clientIdentifierResponse.Name),
		Capacity: types.Int32Value(int32(clientIdentifierResponse.Capacity)),
	}

	if !plan.Description.IsNull() || clientIdentifierResponse.Description != "" {
		ruleClientIdentifier.Description = types.StringValue(clientIdentifierResponse.Description)
	}

	return ruleClientIdentifier
}

func (r *ruleClientIdentifierStorageResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var plan ruleClientIdentifierStorageModel
	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
