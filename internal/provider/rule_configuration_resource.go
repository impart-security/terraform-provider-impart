package provider

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/impart-security/terraform-provider-impart/internal/apiclient"
	openapiclient "github.com/impart-security/terraform-provider-impart/internal/apiclient"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &ruleConfigurationResource{}
	_ resource.ResourceWithConfigure   = &ruleConfigurationResource{}
	_ resource.ResourceWithImportState = &ruleConfigurationResource{}
)

// NewRuleConfigurationResource is a helper function to simplify the provider implementation.
func NewRuleConfigurationResource() resource.Resource {
	return &ruleConfigurationResource{}
}

// ruleConfigurationResource is the resource implementation.
type ruleConfigurationResource struct {
	client   *impartAPIClient
	typeName string
}

// ruleConfigurationResourceModel maps the resource schema data.
type ruleConfigurationResourceModel struct {
	ID       types.String `tfsdk:"id"`
	Slug     types.String `tfsdk:"slug"`
	Config   types.String `tfsdk:"config"`
	Disabled types.Bool   `tfsdk:"disabled"`
	//Labels   []types.String `tfsdk:"labels"` // will be added once the API supports it
}

// Configure adds the provider configured client to the resource.
func (r *ruleConfigurationResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *ruleConfigurationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_core_rule_config"
}

// Schema defines the schema for the resource.
func (r *ruleConfigurationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage a core rule configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for the core rule.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"slug": schema.StringAttribute{
				Description: "The slug for the core rule.",
				Required:    true,
			},
			"config": schema.StringAttribute{
				Description: "the core rule configuration.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					conditionalEqualJSONPlanModifier{
						Subset: true,
					},
				},
			},
			"disabled": schema.BoolAttribute{
				Description: "Set true to disable the core rule.",
				Required:    true,
			},
			// "labels": schema.ListAttribute{
			// 	Description: "The applied labels.",
			// 	ElementType: types.StringType,
			// 	Optional:    true,
			// },
		},
	}
}

func (r *ruleConfigurationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Create a new resource.
func (r *ruleConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the core rule configuration resource")
	// Retrieve values from plan
	var plan ruleConfigurationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ruleConfigurationPostBody, err := toRuleConfigurationPostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create the core rule configuration",
			err.Error(),
		)
		return
	}

	ruleRequest := r.client.CoreRulesAPI.UpdateCoreRule(ctx, r.client.OrgID, plan.Slug.ValueString()).
		CoreRulePostBody(ruleConfigurationPostBody)

	ruleResponse, _, err := ruleRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the core rule configuration",
			message,
		)
		return
	}

	// Map response body to model
	state, err := toRuleConfigurationModel(ruleResponse, plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to map the core rule configuration response to a model",
			err.Error(),
		)
		return
	}

	state.Config = plan.Config

	// Set state to fully populated data
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Created the core rule configuration resource", map[string]any{"success": true})
}

// Read resource information.
func (r *ruleConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the core rule configuration resource")
	// Get current state
	var state ruleConfigurationResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ruleResponse, httpResp, err := r.client.CoreRulesAPI.GetCoreRule(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
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
			"Unable to read the core rule configuration",
			message,
		)
		return
	}

	// Map response body to model
	newState, err := toRuleConfigurationModel(ruleResponse, state)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to map the core rule configuration response to model",
			err.Error(),
		)
		return
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Finished reading the core rule configuration resource", map[string]any{"success": true})
}

func (r *ruleConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the core rule configuration resource")
	// Retrieve values from plan
	var plan ruleConfigurationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ruleConfigurationPostBody, err := toRuleConfigurationPostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update the core rule configuration",
			err.Error(),
		)
		return
	}

	ruleRequest := r.client.CoreRulesAPI.UpdateCoreRule(ctx, r.client.OrgID, plan.Slug.ValueString()).
		CoreRulePostBody(ruleConfigurationPostBody)

	// update rule
	ruleResponse, _, err := ruleRequest.Execute()

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the core rule configuration",
			message,
		)
		return
	}

	state, err := toRuleConfigurationModel(ruleResponse, plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to map the core rule configuration response to model",
			err.Error(),
		)
		return
	}

	state.Config = plan.Config

	// Set refreshed state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updated the core rule configuration resource", map[string]any{"success": true})
}

func (r *ruleConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the core rule configuration resource")
	// Retrieve values from the state
	var state ruleConfigurationResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	slug := strings.ToLower(state.Slug.ValueString())

	ruleConfigurationPostBody := openapiclient.CoreRulePostBody{
		Disabled: true,
		Config:   openapiclient.CoreRuleConfig(fmt.Sprintf(`{"app_slug": "%s"}`, slug)),
		Labels:   []string{},
	}

	ruleRequest := r.client.CoreRulesAPI.UpdateCoreRule(ctx, r.client.OrgID, state.ID.ValueString()).
		CoreRulePostBody(ruleConfigurationPostBody)

	// update rule
	_, _, err := ruleRequest.Execute()

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the core rule configuration",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the core rule configuration resource", map[string]any{"success": true})
}

func (r *ruleConfigurationResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var plan ruleConfigurationResourceModel
	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func toRuleConfigurationPostBody(plan ruleConfigurationResourceModel) (apiclient.CoreRulePostBody, error) {
	postBody := openapiclient.CoreRulePostBody{
		Disabled: plan.Disabled.ValueBool(),
	}

	if !plan.Config.IsNull() {
		postBody.Config = openapiclient.CoreRuleConfig(plan.Config.ValueString())
	}

	// if len(plan.Labels) > 0 {
	// 	postBody.Labels = make([]string, 0, len(plan.Labels))
	// 	for _, label := range plan.Labels {
	// 		postBody.Labels = append(postBody.Labels, label.ValueString())
	// 	}
	// }

	return postBody, nil
}

func toRuleConfigurationModel(coreRuleResponse *openapiclient.CoreRule, plan ruleConfigurationResourceModel) (ruleConfigurationResourceModel, error) {
	ruleConfigurationModel := ruleConfigurationResourceModel{
		ID:       types.StringValue(coreRuleResponse.Id),
		Slug:     types.StringValue(coreRuleResponse.Name),
		Disabled: types.BoolValue(coreRuleResponse.Disabled),
	}

	//ruleConfigurationModel.Labels = buildStateList(plan.Labels, coreRuleResponse.Labels)

	// normalize name to use plan
	if strings.EqualFold(coreRuleResponse.Name, plan.Slug.ValueString()) {
		ruleConfigurationModel.Slug = types.StringValue(plan.Slug.ValueString())
	}

	// track only if content was set
	if !plan.Config.IsNull() {
		ruleConfigurationModel.Config = types.StringValue(string(coreRuleResponse.Config))
	}

	return ruleConfigurationModel, nil
}
