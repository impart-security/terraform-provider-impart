package provider

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapiclient "github.com/impart-security/terraform-provider-impart/internal/apiclient"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &ruleResource{}
	_ resource.ResourceWithConfigure   = &ruleResource{}
	_ resource.ResourceWithImportState = &ruleResource{}
)

// NewRuleResource is a helper function to simplify the provider implementation.
func NewRuleResource() resource.Resource {
	return &ruleResource{}
}

// ruleResource is the resource implementation.
type ruleResource struct {
	client   *impartAPIClient
	typeName string
}

// ruleResourceModel maps the resource schema data.
type ruleResourceModel struct {
	ID             types.String   `tfsdk:"id"`
	Name           types.String   `tfsdk:"name"`
	Description    types.String   `tfsdk:"description"`
	Disabled       types.Bool     `tfsdk:"disabled"`
	SourceFile     types.String   `tfsdk:"source_file"`
	SourceHash     types.String   `tfsdk:"source_hash"`
	Content        types.String   `tfsdk:"content"`
	BlockingEffect types.String   `tfsdk:"blocking_effect"`
	Labels         []types.String `tfsdk:"labels"`
	Type           types.String   `tfsdk:"type"`
}

// Configure adds the provider configured client to the resource.
func (r *ruleResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *ruleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rule"
}

// Schema defines the schema for the resource.
func (r *ruleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage a rule.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for the rule.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name for this rule.",
				Required:    true,
			},

			"description": schema.StringAttribute{
				Description: "The description for this rule.",
				Optional:    true,
			},

			"disabled": schema.BoolAttribute{
				Description: "Set true to disable the rule.",
				Required:    true,
			},

			"source_file": schema.StringAttribute{
				Description: "The rule source file.",
				Optional:    true,
			},

			"source_hash": schema.StringAttribute{
				Description: "The rule source hash.",
				Optional:    true,
			},

			"content": schema.StringAttribute{
				Description: "The rule body content.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					conditionalEqualJSONPlanModifier{
						AttrPath:  path.Root("type"),
						AttrValue: "recipe",
					},
				},
			},

			"blocking_effect": schema.StringAttribute{
				Description: "The rule blocking effect. Allowed values: block, simulate. If not set effect will be block.",
				Validators:  []validator.String{stringvalidator.OneOf(string(openapiclient.BLOCKINGEFFECTTYPE_BLOCK), string(openapiclient.BLOCKINGEFFECTTYPE_SIMULATE))},
				Optional:    true,
			},

			"labels": schema.ListAttribute{
				Description: "The applied labels.",
				ElementType: types.StringType,
				Optional:    true,
			},

			"type": schema.StringAttribute{
				Description: "The type of the rule. Allowed values: script, recipe.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *ruleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Create a new resource.
func (r *ruleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the rule resource")
	// Retrieve values from plan
	var plan ruleResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if plan.SourceFile.IsNull() && plan.Content.IsNull() {
		resp.Diagnostics.AddError(
			"Configuration Error: Missing Required Argument",
			"Either 'source_file' or 'content' must be set.",
		)
		return
	}

	if !plan.SourceFile.IsNull() && !plan.Content.IsNull() {
		resp.Diagnostics.AddError(
			"Configuration Error: Conflicting Arguments",
			"Both 'source_file' and 'content' cannot be set at the same time.",
		)
		return
	}

	if plan.SourceFile.IsNull() && !plan.SourceHash.IsNull() {
		resp.Diagnostics.AddError(
			"Configuration Error: Conflicting Arguments",
			"'source_hash' can only be set when `source_file` is set.",
		)
		return
	}

	rulesScriptPostBody, err := toRulePostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create the rule",
			err.Error(),
		)
		return
	}

	ruleResponse, err := r.createRule(ctx, plan.Type.ValueString(), rulesScriptPostBody, &resp.Diagnostics)

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the rule",
			message,
		)
		return
	}

	// Map response body to model
	state, err := toRuleModel(ruleResponse, plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to map the rule response to model",
			err.Error(),
		)
		return
	}

	// Set the content as it was set in the plan to avoid errors
	if plan.Type.ValueString() == "recipe" {
		state.Content = plan.Content
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Created the rule resource", map[string]any{"success": true})
}

// Read resource information.
func (r *ruleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the rule resource")
	// Get current state
	var state ruleResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ruleResponse, httpResp, err := r.readRule(ctx, state.Type.ValueString(), state.ID.ValueString())

	if err != nil {
		// Treat HTTP 404 Not Found status as a signal to remove/recreate resource
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}

		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to read the rule",
			message,
		)
		return
	}

	// Map response body to model
	newState, err := toRuleModel(ruleResponse, state)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to map the rule response to model",
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
	tflog.Debug(ctx, "Finished reading the rule resource", map[string]any{"success": true})
}

func (r *ruleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the rule resource")
	// Retrieve values from plan
	var plan ruleResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	rulesScriptPostBody, err := toRulePostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update the rule",
			err.Error(),
		)
		return
	}

	ruleResponse, err := r.updateRule(ctx, plan.Type.ValueString(), plan.ID.ValueString(), rulesScriptPostBody)

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the rule",
			message,
		)
		return
	}

	state, err := toRuleModel(ruleResponse, plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to map the rule response to model",
			err.Error(),
		)
		return
	}

	// Set the content as it was set in the plan to avoid errors
	if plan.Type.ValueString() == "recipe" {
		state.Content = plan.Content
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updated the rule resource", map[string]any{"success": true})
}

func (r *ruleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the rule resource")
	// Retrieve values from the state
	var state ruleResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.deleteRule(ctx, state.Type.ValueString(), state.ID.ValueString(), openapiclient.RulesScriptPostBody{})
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the rule",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the rule resource", map[string]any{"success": true})
}

func (r *ruleResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var plan ruleResourceModel
	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.typeName == "rule_script" {
		resp.Diagnostics.AddWarning(
			"Deprecated Resource Type",
			"The `impart_rule_script` resource is deprecated and will be removed in a future release. Please migrate to `impart_rule`.",
		)
	}

	if plan.SourceFile.IsNull() && plan.Content.IsNull() {
		resp.Diagnostics.AddError(
			"Configuration Error: Missing Required Argument",
			"Either 'source_file' or 'content' must be set.",
		)
		return
	}

	if !plan.SourceFile.IsNull() && !plan.Content.IsNull() {
		resp.Diagnostics.AddError(
			"Configuration Error: Conflicting Arguments",
			"Both 'source_file' and 'content' cannot be set at the same time.",
		)
		return
	}

	if plan.SourceFile.IsNull() && !plan.SourceHash.IsNull() {
		resp.Diagnostics.AddError(
			"Configuration Error: Conflicting Arguments",
			"'source_hash' can only be set when `source_file` is set.",
		)
		return
	}
}

func toRulePostBody(plan ruleResourceModel) (openapiclient.RulesScriptPostBody, error) {
	postBody := openapiclient.RulesScriptPostBody{
		Name:     plan.Name.ValueString(),
		Disabled: plan.Disabled.ValueBool(),
	}

	if !plan.Description.IsNull() {
		description := plan.Description.ValueString()
		postBody.Description = &description
	}

	if len(plan.Labels) > 0 {
		postBody.Labels = make([]string, 0, len(plan.Labels))
		for _, label := range plan.Labels {
			postBody.Labels = append(postBody.Labels, label.ValueString())
		}
	}

	blockingEffectVal := string(openapiclient.BLOCKINGEFFECTTYPE_BLOCK)
	if !plan.BlockingEffect.IsNull() {
		blockingEffectVal = plan.BlockingEffect.ValueString()
	}

	blockingEffect, err := openapiclient.NewBlockingEffectTypeFromValue(blockingEffectVal)
	if err != nil {
		return postBody, fmt.Errorf("Unable to create blocking effect: %w", err)
	}
	postBody.BlockingEffect = blockingEffect

	var ruleBytes []byte
	if !plan.SourceFile.IsNull() {
		bytes, err := os.ReadFile(plan.SourceFile.ValueString())
		if err != nil {
			return postBody, fmt.Errorf("Unable to read the rule source file: %w", err)
		}
		ruleBytes = bytes
	} else {
		ruleBytes = []byte(plan.Content.ValueString())
	}

	// base64 only for rules
	if plan.Type.ValueString() == "script" {
		ruleb64 := base64.StdEncoding.EncodeToString(ruleBytes)
		postBody.Src = ruleb64
	} else {
		// rule recipe
		postBody.Src = string(ruleBytes)
	}

	return postBody, nil
}

func toRuleModel(ruleScriptResponse *openapiclient.RulesScript, plan ruleResourceModel) (ruleResourceModel, error) {
	ruleModel := ruleResourceModel{
		ID:         types.StringValue(ruleScriptResponse.Id),
		Name:       types.StringValue(ruleScriptResponse.Name),
		SourceFile: plan.SourceFile,
		Disabled:   types.BoolValue(ruleScriptResponse.Disabled),
		SourceHash: plan.SourceHash,
	}

	if !plan.Description.IsNull() || ruleScriptResponse.Description != "" {
		ruleModel.Description = types.StringValue(ruleScriptResponse.Description)
	}

	ruleModel.Labels = buildStateList(plan.Labels, ruleScriptResponse.Labels)

	if !(plan.BlockingEffect.IsNull() && string(ruleScriptResponse.BlockingEffect) == string(openapiclient.BLOCKINGEFFECTTYPE_BLOCK)) {
		ruleModel.BlockingEffect = types.StringValue(string(ruleScriptResponse.BlockingEffect))
	}

	// track only if content was set
	if !plan.Content.IsNull() {
		if plan.Type.ValueString() == "script" {
			bytes, err := base64.StdEncoding.DecodeString(ruleScriptResponse.Src)
			if err != nil {
				return ruleModel, fmt.Errorf("Unable to base64 decode the rule: %w", err)
			}

			ruleModel.Content = types.StringValue(string(bytes))
		} else {
			// rule recipe
			ruleModel.Content = types.StringValue(string(ruleScriptResponse.Src))
		}
	}

	// track hash only if user originally set source_hash or content
	if !plan.SourceHash.IsNull() || !plan.Content.IsNull() {
		bytes := []byte(ruleScriptResponse.Src)
		if plan.Type.ValueString() == "script" {
			b, err := base64.StdEncoding.DecodeString(ruleScriptResponse.Src)
			if err != nil {
				return ruleModel, fmt.Errorf("Unable to base64 decode the rule: %w", err)
			}
			bytes = b
		}

		hash, err := calculateSha256(string(bytes))
		if err != nil {
			return ruleModel, fmt.Errorf("Unable to calculate sha256: %w", err)
		}

		if !plan.SourceHash.IsNull() {
			ruleModel.SourceHash = types.StringValue(hash)
		}

		if !plan.Content.IsNull() {
			ruleModel.Content = types.StringValue(string(bytes))
		}
	}

	// Type is ethermal field defined in the terraform schema
	ruleModel.Type = plan.Type
	return ruleModel, nil
}

func (r *ruleResource) createRule(ctx context.Context, ruleType string, ruleBody openapiclient.RulesScriptPostBody, diag *diag.Diagnostics) (*openapiclient.RulesScript, error) {
	if ruleType == "" || ruleType == "script" {
		ruleRequest := r.client.RulesScriptsAPI.CreateRulesScript(ctx, r.client.OrgID).
			RulesScriptPostBody(ruleBody)

		ruleResponse, _, err := ruleRequest.Execute()
		if err != nil {
			if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
				return nil, fmt.Errorf("%s %s", apiErr.Error(), string(apiErr.Body()))
			}
			return nil, fmt.Errorf("Unable to create the rule: %w", err)
		}

		return ruleResponse, nil
	} else if ruleType == "recipe" {
		ruleRecipeBody := openapiclient.RuleRecipePostBody{
			Name:           ruleBody.Name,
			Description:    ruleBody.Description,
			Disabled:       ruleBody.Disabled,
			Labels:         ruleBody.Labels,
			BlockingEffect: *ruleBody.BlockingEffect,
			Components:     openapiclient.RuleRecipeComponents(ruleBody.Src),
		}

		ruleRequest := r.client.RuleRecipesAPI.CreateRuleRecipe(ctx, r.client.OrgID).
			RuleRecipePostBody(ruleRecipeBody)

		ruleRecipeResponse, _, err := ruleRequest.Execute()
		if err != nil {
			if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
				return nil, fmt.Errorf("%s %s", apiErr.Error(), string(apiErr.Body()))
			}
			return nil, fmt.Errorf("Unable to create the rule recipe: %w", err)
		}

		response := &openapiclient.RulesScript{
			Id:             ruleRecipeResponse.Id,
			Name:           ruleRecipeResponse.Name,
			Description:    ruleRecipeResponse.Description,
			Disabled:       ruleRecipeResponse.Disabled,
			Labels:         ruleRecipeResponse.Labels,
			BlockingEffect: ruleRecipeResponse.BlockingEffect,
			Src:            string(ruleRecipeResponse.Components),
		}

		return response, nil
	}
	return nil, fmt.Errorf("Invalid rule type: %s", ruleType)
}

func (r *ruleResource) readRule(ctx context.Context, ruleType string, ruleID string) (*openapiclient.RulesScript, *http.Response, error) {
	if ruleType == "" || ruleType == "script" {
		ruleRequest := r.client.RulesScriptsAPI.GetRulesScript(ctx, r.client.OrgID, ruleID)

		ruleResponse, resp, err := ruleRequest.Execute()
		if err != nil {
			if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
				return nil, resp, fmt.Errorf("%s %s", apiErr.Error(), string(apiErr.Body()))
			}
			return nil, resp, fmt.Errorf("Unable to read the rule: %w", err)
		}

		return ruleResponse, resp, nil
	} else if ruleType == "recipe" {
		ruleRequest := r.client.RuleRecipesAPI.GetRuleRecipe(ctx, r.client.OrgID, ruleID)

		ruleRecipeResponse, resp, err := ruleRequest.Execute()
		if err != nil {
			if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
				return nil, resp, fmt.Errorf("%s %s", apiErr.Error(), string(apiErr.Body()))
			}
			return nil, resp, fmt.Errorf("Unable to read the rule recipe: %w", err)
		}

		response := &openapiclient.RulesScript{
			Id:             ruleRecipeResponse.Id,
			Name:           ruleRecipeResponse.Name,
			Description:    ruleRecipeResponse.Description,
			Disabled:       ruleRecipeResponse.Disabled,
			Labels:         ruleRecipeResponse.Labels,
			BlockingEffect: ruleRecipeResponse.BlockingEffect,
			Src:            string(ruleRecipeResponse.Components),
		}

		return response, resp, nil
	}
	return nil, nil, fmt.Errorf("Invalid rule type: %s", ruleType)
}

func (r *ruleResource) updateRule(ctx context.Context, ruleType string, ruleID string, ruleBody openapiclient.RulesScriptPostBody) (*openapiclient.RulesScript, error) {
	if ruleType == "" || ruleType == "script" {
		ruleRequest := r.client.RulesScriptsAPI.UpdateRulesScript(ctx, r.client.OrgID, ruleID).
			RulesScriptPostBody(ruleBody)

		ruleResponse, _, err := ruleRequest.Execute()
		if err != nil {
			if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
				return nil, fmt.Errorf("%s %s", apiErr.Error(), string(apiErr.Body()))
			}
			return nil, fmt.Errorf("Unable to update the rule: %w", err)
		}

		return ruleResponse, nil
	} else if ruleType == "recipe" {
		ruleRecipeBody := openapiclient.RuleRecipePostBody{
			Name:           ruleBody.Name,
			Disabled:       ruleBody.Disabled,
			Description:    ruleBody.Description,
			Labels:         ruleBody.Labels,
			BlockingEffect: *ruleBody.BlockingEffect,
			Components:     openapiclient.RuleRecipeComponents(ruleBody.Src),
		}
		ruleRequest := r.client.RuleRecipesAPI.UpdateRuleRecipe(ctx, r.client.OrgID, ruleID).
			RuleRecipePostBody(ruleRecipeBody)

		ruleRecipeResponse, _, err := ruleRequest.Execute()
		if err != nil {
			if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
				return nil, fmt.Errorf("%s %s", apiErr.Error(), string(apiErr.Body()))
			}
			return nil, fmt.Errorf("Unable to update the rule recipe: %w", err)
		}

		response := &openapiclient.RulesScript{
			Id:             ruleRecipeResponse.Id,
			Name:           ruleRecipeResponse.Name,
			Description:    ruleRecipeResponse.Description,
			Disabled:       ruleRecipeResponse.Disabled,
			Labels:         ruleRecipeResponse.Labels,
			BlockingEffect: ruleRecipeResponse.BlockingEffect,
			Src:            string(ruleRecipeResponse.Components),
		}

		return response, nil
	}
	return nil, fmt.Errorf("Invalid rule type: %s", ruleType)
}

func (r *ruleResource) deleteRule(ctx context.Context, ruleType string, ruleID string, ruleBody openapiclient.RulesScriptPostBody) error {
	if ruleType == "" || ruleType == "script" {
		ruleRequest := r.client.RulesScriptsAPI.DeleteRulesScript(ctx, r.client.OrgID, ruleID)

		_, err := ruleRequest.Execute()
		if err != nil {
			if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
				return fmt.Errorf("%s %s", apiErr.Error(), string(apiErr.Body()))
			}
			return fmt.Errorf("Unable to delete the rule: %w", err)
		}

		return nil
	} else if ruleType == "recipe" {
		ruleRequest := r.client.RuleRecipesAPI.DeleteRuleRecipe(ctx, r.client.OrgID, ruleID)

		_, err := ruleRequest.Execute()
		if err != nil {
			if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
				return fmt.Errorf("%s %s", apiErr.Error(), string(apiErr.Body()))
			}
			return fmt.Errorf("Unable to delete the rule recipe: %w", err)
		}
		return nil
	}
	return fmt.Errorf("Invalid rule type: %s", ruleType)
}
