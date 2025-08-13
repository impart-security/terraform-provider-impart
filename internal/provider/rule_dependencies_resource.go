package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapiclient "github.com/impart-security/terraform-provider-impart/internal/apiclient"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &ruleDependenciesResource{}
	_ resource.ResourceWithConfigure = &ruleDependenciesResource{}
)

// NewRuleDependenciesResource is a helper function to simplify the provider implementation.
func NewRuleDependenciesResource() resource.Resource {
	return &ruleDependenciesResource{}
}

// ruleDependenciesResource is the resource implementation.
type ruleDependenciesResource struct {
	client *impartAPIClient
}

// ruleDependenciesResourceModel maps the resource schema data.
type ruleDependenciesResourceModel struct {
	Dependencies []ruleDependsOn `tfsdk:"dependencies"`
}

type ruleDependsOn struct {
	RuleID    types.String `tfsdk:"rule_id"`
	DependsOn []string     `tfsdk:"depends_on"`
}

// Configure adds the provider configured client to the resource.
func (r *ruleDependenciesResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *ruleDependenciesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rule_dependencies"
}

// Schema defines the schema for the resource.
func (r *ruleDependenciesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage rule dependencies. There should only ever be one instance of this resource in a workspace at once, because it manages rule dependencies at an organization level.",
		Attributes: map[string]schema.Attribute{
			"dependencies": schema.ListNestedAttribute{
				Description: "An array of rules and the other ids of the rules they depend on before executing.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"rule_id": schema.StringAttribute{
							Required:    true,
							Description: "The ID of the rule",
						},
						"depends_on": schema.ListAttribute{
							ElementType: types.StringType,
							Required:    true,
							Description: "IDs of the rule this rule depends on.",
						},
					},
				},
				Validators: []validator.List{
					uniqueValue("rule_id"),
				},
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

// Create a new resource.
func (r *ruleDependenciesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the rule dependencies resource")
	// Retrieve values from plan
	var plan ruleDependenciesResourceModel
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
			RuleId:       plan.Dependencies[i].RuleID.ValueString(),
			Dependencies: plan.Dependencies[i].DependsOn,
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
			"Unable to create the rule dependencies",
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

	applyRuleDependencyResponseToState(responseMap, &plan)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Created the rule dependencies resource", map[string]any{"success": true})
}

// Read resource information.
func (r *ruleDependenciesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the rule dependencies resource")
	// Get current state
	var state ruleDependenciesResourceModel
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
			"Unable to read the rule dependencies",
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

	applyRuleDependencyResponseToState(responseMap, &state)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Finished reading the rule dependencies resource", map[string]any{"success": true})
}

func (r *ruleDependenciesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the rule dependencies resource")
	// Retrieve values from plan
	var plan ruleDependenciesResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ruleDependenciesResourceModel
	diags = req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(plan.Dependencies) == 0 {
		resp.Diagnostics.AddError(
			"Unable to create the rule dependencies",
			"Dependencies must be present and not empty",
		)
		return
	}

	rulesScriptPutBody := make([]openapiclient.RulesDependenciesPutBodyInner, len(plan.Dependencies))
	for i := range plan.Dependencies {
		rulesScriptPutBody[i] = openapiclient.RulesDependenciesPutBodyInner{
			RuleId:       plan.Dependencies[i].RuleID.ValueString(),
			Dependencies: plan.Dependencies[i].DependsOn,
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
			"Unable to update the rule dependencies",
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

	newState := ruleDependenciesResourceModel{
		Dependencies: plan.Dependencies,
	}
	applyRuleDependencyResponseToState(responseMap, &newState)

	// Set refreshed state
	diags = resp.State.Set(ctx, newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updated the rule dependencies resource", map[string]any{"success": true})
}

func (r *ruleDependenciesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the rule dependencies resource")
	// Retrieve values from the state
	var state ruleDependenciesResourceModel
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
			"Unable to read the rule dependencies",
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
			"Unable to delete the rule dependencies",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the rule dependencies resource", map[string]any{"success": true})
}

func applyRuleDependencyResponseToState(responseItemsMap map[string]map[string]bool,
	state *ruleDependenciesResourceModel) {

	dependenciesState := []ruleDependsOn{}

	for i := range state.Dependencies {
		stateRuleID := state.Dependencies[i].RuleID.ValueString()
		if deps, ok := responseItemsMap[stateRuleID]; ok {
			ruleIDs := []string{}
			newRuleIDs := []string{}

			for j := range state.Dependencies[i].DependsOn {
				ruleID := state.Dependencies[i].DependsOn[j]

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

			dependenciesState = append(dependenciesState, ruleDependsOn{
				RuleID:    state.Dependencies[i].RuleID,
				DependsOn: append(ruleIDs, newRuleIDs...),
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
		dependenciesState = append(dependenciesState, ruleDependsOn{
			RuleID:    types.StringValue(ruleID),
			DependsOn: deps,
		})
	}

	state.Dependencies = dependenciesState
}
