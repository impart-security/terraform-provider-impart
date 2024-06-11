package provider

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

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
	_ resource.Resource                = &ruleScriptResource{}
	_ resource.ResourceWithConfigure   = &ruleScriptResource{}
	_ resource.ResourceWithImportState = &ruleScriptResource{}
)

// NewRuleScriptResource is a helper function to simplify the provider implementation.
func NewRuleScriptResource() resource.Resource {
	return &ruleScriptResource{}
}

// ruleScriptResource is the resource implementation.
type ruleScriptResource struct {
	client *impartAPIClient
}

// ruleScriptResourceModel maps the resource schema data.
type ruleScriptResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Disabled    types.Bool   `tfsdk:"disabled"`
	SourceFile  types.String `tfsdk:"source_file"`
	SourceHash  types.String `tfsdk:"source_hash"`
	Content     types.String `tfsdk:"content"`
}

// Configure adds the provider configured client to the resource.
func (r *ruleScriptResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *ruleScriptResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rule_script"
}

// Schema defines the schema for the resource.
func (r *ruleScriptResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage a rule script.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for the rule script.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name for this rule script.",
				Required:    true,
			},

			"description": schema.StringAttribute{
				Description: "The description for this rule script.",
				Optional:    true,
			},

			"disabled": schema.BoolAttribute{
				Description: "Set true to disable the rule script.",
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
			},
		},
	}
}

func (r *ruleScriptResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Create a new resource.
func (r *ruleScriptResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the rule script resource")
	// Retrieve values from plan
	var plan ruleScriptResourceModel
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
			"Both 'source_hash' can only be set when `source_file` is set.",
		)
		return
	}

	var ruleBytes []byte
	if !plan.SourceFile.IsNull() {
		bytes, err := os.ReadFile(plan.SourceFile.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to read the rule script source file",
				err.Error(),
			)
			return
		}
		ruleBytes = bytes
	} else {
		ruleBytes = []byte(plan.Content.ValueString())
	}

	ruleb64 := base64.StdEncoding.EncodeToString(ruleBytes)

	rulesScriptPostBody := openapiclient.RulesScriptPostBody{
		Name:     plan.Name.ValueString(),
		Src:      ruleb64,
		Disabled: plan.Disabled.ValueBool(),
	}

	if !plan.Description.IsNull() {
		description := plan.Description.ValueString()
		rulesScriptPostBody.Description = &description
	}

	// Create new rule
	ruleRequest := r.client.RulesScriptsApi.CreateRulesScript(ctx, r.client.OrgID).
		RulesScriptPostBody(rulesScriptPostBody)

	ruleResponse, _, err := ruleRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the rule script",
			message,
		)
		return
	}

	// Map response body to model
	plan.ID = types.StringValue(ruleResponse.Id)
	plan.Name = types.StringValue(ruleResponse.Name)
	plan.Description = types.StringValue(ruleResponse.Description)
	plan.Disabled = types.BoolValue(ruleResponse.Disabled)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Created the rule script resource", map[string]any{"success": true})
}

// Read resource information.
func (r *ruleScriptResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the rule script resource")
	// Get current state
	var state ruleScriptResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ruleResponse, httpResp, err := r.client.RulesScriptsApi.GetRulesScript(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
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
			"Unable to read the rule script",
			message,
		)
		return
	}

	// Map response body to model
	newState := ruleScriptResourceModel{
		ID:          types.StringValue(ruleResponse.Id),
		Name:        types.StringValue(ruleResponse.Name),
		SourceFile:  state.SourceFile,
		Description: types.StringValue(ruleResponse.Description),
		Disabled:    types.BoolValue(ruleResponse.Disabled),
	}

	// track hash only if user originally set source_hash or content
	if !state.SourceHash.IsNull() || !state.Content.IsNull() {
		bytes, err := base64.StdEncoding.DecodeString(ruleResponse.Src)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to base64 decode the rule script",
				err.Error(),
			)
		}
		hash, err := calculateSha256(string(bytes))
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to calculate sha256",
				err.Error(),
			)
		}

		if !state.SourceHash.IsNull() {
			newState.SourceHash = types.StringValue(hash)
		}

		if !state.Content.IsNull() {
			newState.Content = types.StringValue(string(bytes))
		}
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Finished reading the rule script resource", map[string]any{"success": true})
}

func (r *ruleScriptResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the rule script resource")
	// Retrieve values from plan
	var plan ruleScriptResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	rule, err := os.ReadFile(plan.SourceFile.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read the rule script source file",
			err.Error(),
		)
		return
	}
	ruleb64 := base64.StdEncoding.EncodeToString(rule)

	rulesScriptPostBody := openapiclient.RulesScriptPostBody{
		Name:     plan.Name.ValueString(),
		Src:      ruleb64,
		Disabled: plan.Disabled.ValueBool(),
	}

	if !plan.Description.IsNull() {
		description := plan.Description.ValueString()
		rulesScriptPostBody.Description = &description
	}

	ruleRequest := r.client.RulesScriptsApi.UpdateRulesScript(ctx, r.client.OrgID, plan.ID.ValueString()).
		RulesScriptPostBody(rulesScriptPostBody)

	// update rule
	ruleResponse, _, err := ruleRequest.Execute()

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the rule script",
			message,
		)
		return
	}

	// Overwrite the rules with refreshed state
	state := ruleScriptResourceModel{
		ID:         types.StringValue(ruleResponse.Id),
		Name:       types.StringValue(ruleResponse.Name),
		SourceFile: plan.SourceFile,
		Disabled:   types.BoolValue(ruleResponse.Disabled),
		SourceHash: plan.SourceHash,
	}

	if !plan.Description.IsNull() {
		state.Description = types.StringValue(ruleResponse.Description)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updated the rule script resource", map[string]any{"success": true})
}

func (r *ruleScriptResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the rule script resource")
	// Retrieve values from the state
	var state ruleScriptResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// delete the rule
	_, err := r.client.RulesScriptsApi.DeleteRulesScript(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the rule script",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the rule script resource", map[string]any{"success": true})
}

func (r *ruleScriptResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var plan ruleScriptResourceModel
	diags := req.Config.Get(ctx, &plan)
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
			"Both 'source_hash' can only be set when `source_file` is set.",
		)
		return
	}
}
