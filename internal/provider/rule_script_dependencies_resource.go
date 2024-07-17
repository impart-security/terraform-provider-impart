package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapiclient "github.com/impart-security/terraform-provider-impart/internal/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &ruleScriptDependenciesResource{}
	_ resource.ResourceWithConfigure = &ruleScriptDependenciesResource{}
)

// NewRuleScriptResource is a helper function to simplify the provider implementation.
func NewRuleScriptDependenciesResource() resource.Resource {
	return &ruleScriptDependenciesResource{}
}

// ruleScriptDependenciesResource is the resource implementation.
type ruleScriptDependenciesResource struct {
	client *impartAPIClient
}

// ruleScriptDependenciesResourceModel maps the resource schema data.
type ruleScriptDependenciesResourceModel struct {
	Dependencies []ruleScriptDependsOn `tfsdk:"dependencies"`
}

type ruleScriptDependsOn struct {
	RuleScriptID            types.String `tfsdk:"rule_script_id"`
	DependsOnRulesScriptIDs []string     `tfsdk:"depends_on_rule_script_ids"`
}

// Configure adds the provider configured client to the resource.
func (r *ruleScriptDependenciesResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *ruleScriptDependenciesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rule_script_dependencies"
}

// Schema defines the schema for the resource.
func (r *ruleScriptDependenciesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage rule script dependencies. There should only ever be one instance of this resource in a workspace at once, because it manages rule script dependencies at an organization level.",
		Attributes: map[string]schema.Attribute{
			"dependencies": schema.ListNestedAttribute{
				Description: "An array of rule scripts and the other ids of the rules they depend on before executing.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"rule_script_id": schema.StringAttribute{
							Required:    true,
							Description: "The ID of the rule script",
						},
						"depends_on_rule_script_ids": schema.ListAttribute{
							ElementType: types.StringType,
							Required:    true,
							Description: "IDs of the rule script this rule depends on.",
						},
					},
				},
				Validators: []validator.List{
					uniqueValue("rule_script_id"),
				},
			},
		},
	}
}

// Create a new resource.
func (r *ruleScriptDependenciesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the rule script dependencies resource")
	// Retrieve values from plan
	var plan ruleScriptDependenciesResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(plan.Dependencies) == 0 {
		// Allow to create empty dependencies resources so users can see the diff
		diags = resp.State.Set(ctx, plan)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		return
	}

	rulesScriptPutBody := make([]openapiclient.RulesDependenciesPutBodyInner, len(plan.Dependencies))
	for i := range plan.Dependencies {
		rulesScriptPutBody[i] = openapiclient.RulesDependenciesPutBodyInner{
			RuleId:       plan.Dependencies[i].RuleScriptID.ValueString(),
			Dependencies: plan.Dependencies[i].DependsOnRulesScriptIDs,
		}
	}

	// Create new rule dependencies
	ruleDependenciesRequest := r.client.RulesDependenciesAPI.UpdateRulesDependencies(ctx, r.client.OrgID).
		RulesDependenciesPutBodyInner(rulesScriptPutBody)

	ruleDependenciesResponse, _, err := ruleDependenciesRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the rule script dependencies",
			message,
		)
		return
	}

	responseMap := make(map[string]map[string]bool)
	for i := range ruleDependenciesResponse {
		depsMap := make(map[string]bool)
		for j := range ruleDependenciesResponse[i].Dependencies {
			depsMap[ruleDependenciesResponse[i].Dependencies[j]] = false
		}

		responseMap[ruleDependenciesResponse[i].RuleId] = depsMap
	}

	applyDependencyResponseToState(responseMap, &plan)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Created the rule script dependencies resource", map[string]any{"success": true})
}

// Read resource information.
func (r *ruleScriptDependenciesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the rule script dependencies resource")
	// Get current state
	var state ruleScriptDependenciesResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	rulesScriptReq := r.client.RulesScriptsAPI.GetRulesScripts(ctx, r.client.OrgID).
		ExcludeSrc(true).
		ExcludeRevisions(true).
		Type_("custom")

	ruleDependenciesResponse, _, err := rulesScriptReq.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to read the rule script dependencies",
			message,
		)
		return
	}

	responseMap := make(map[string]map[string]bool)
	for i := range ruleDependenciesResponse.Items {
		depsMap := make(map[string]bool)
		for j := range ruleDependenciesResponse.Items[i].Dependencies {
			depsMap[ruleDependenciesResponse.Items[i].Dependencies[j]] = false
		}
		responseMap[ruleDependenciesResponse.Items[i].Id] = depsMap
	}

	applyDependencyResponseToState(responseMap, &state)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Finished reading the rule script dependencies resource", map[string]any{"success": true})
}

func (r *ruleScriptDependenciesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the rule script dependencies resource")
	// Retrieve values from plan
	var plan ruleScriptDependenciesResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ruleScriptDependenciesResourceModel
	diags = req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(plan.Dependencies) == 0 {
		resp.Diagnostics.AddError(
			"Unable to create the rule script dependencies",
			"Dependencies must be present and not empty",
		)
		return
	}

	rulesScriptPutBody := make([]openapiclient.RulesDependenciesPutBodyInner, len(plan.Dependencies))
	for i := range plan.Dependencies {
		rulesScriptPutBody[i] = openapiclient.RulesDependenciesPutBodyInner{
			RuleId:       plan.Dependencies[i].RuleScriptID.ValueString(),
			Dependencies: plan.Dependencies[i].DependsOnRulesScriptIDs,
		}
	}

	ruleDependenciesRequest := r.client.RulesDependenciesAPI.UpdateRulesDependencies(ctx, r.client.OrgID).
		RulesDependenciesPutBodyInner(rulesScriptPutBody)

	ruleDependenciesResponse, _, err := ruleDependenciesRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the rule script dependencies",
			message,
		)
		return
	}

	responseMap := make(map[string]map[string]bool)
	for i := range ruleDependenciesResponse {
		depsMap := make(map[string]bool)
		for j := range ruleDependenciesResponse[i].Dependencies {
			depsMap[ruleDependenciesResponse[i].Dependencies[j]] = false
		}
		responseMap[ruleDependenciesResponse[i].RuleId] = depsMap
	}

	newState := ruleScriptDependenciesResourceModel{
		Dependencies: plan.Dependencies,
	}
	applyDependencyResponseToState(responseMap, &newState)

	// Set refreshed state
	diags = resp.State.Set(ctx, newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updated the rule script dependencies resource", map[string]any{"success": true})
}

func (r *ruleScriptDependenciesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the rule script dependencies resource")
	// Retrieve values from the state
	var state ruleScriptDependenciesResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	getRuleDependenciesResponse, _, err := r.client.RulesScriptsAPI.GetRulesScripts(ctx, r.client.OrgID).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to read the rule script dependencies",
			message,
		)
		return
	}

	rulesScriptPutBody := make([]openapiclient.RulesDependenciesPutBodyInner, len(getRuleDependenciesResponse.Items))
	for i := range getRuleDependenciesResponse.Items {
		rulesScriptPutBody[i] = openapiclient.RulesDependenciesPutBodyInner{
			RuleId:       getRuleDependenciesResponse.Items[i].Id,
			Dependencies: []string{},
		}
	}
	ruleDependenciesRequest := r.client.RulesDependenciesAPI.UpdateRulesDependencies(ctx, r.client.OrgID).
		RulesDependenciesPutBodyInner(rulesScriptPutBody)

	// Delete the rule
	_, _, err = ruleDependenciesRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the rule script dependencies",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the rule script dependencies resource", map[string]any{"success": true})
}

func applyDependencyResponseToState(responseItemsMap map[string]map[string]bool,
	state *ruleScriptDependenciesResourceModel) {

	dependenciesState := []ruleScriptDependsOn{}

	for i := range state.Dependencies {
		stateRuleID := state.Dependencies[i].RuleScriptID.ValueString()
		if deps, ok := responseItemsMap[stateRuleID]; ok {
			ruleIDs := []string{}
			newRuleIDs := []string{}

			for j := range state.Dependencies[i].DependsOnRulesScriptIDs {
				ruleID := state.Dependencies[i].DependsOnRulesScriptIDs[j]

				if _, ok = deps[ruleID]; ok {
					ruleIDs = append(ruleIDs, ruleID)
					deps[ruleID] = true
				}
			}

			// Append new rule dependencies which are not in the state
			for k, v := range deps {
				if !v {
					newRuleIDs = append(newRuleIDs, k)
				}
			}

			dependenciesState = append(dependenciesState, ruleScriptDependsOn{
				RuleScriptID:            state.Dependencies[i].RuleScriptID,
				DependsOnRulesScriptIDs: append(ruleIDs, newRuleIDs...),
			})
			delete(responseItemsMap, stateRuleID)
		}
	}

	// Add new dependencies
	for ruleID, depsMap := range responseItemsMap {
		// Skip adding empty array since dependencies are not in the state
		if len(depsMap) == 0 {
			continue
		}
		deps := make([]string, 0, len(depsMap))
		for key := range depsMap {
			deps = append(deps, key)
		}
		dependenciesState = append(dependenciesState, ruleScriptDependsOn{
			RuleScriptID:            types.StringValue(ruleID),
			DependsOnRulesScriptIDs: deps,
		})
	}

	state.Dependencies = dependenciesState
}
