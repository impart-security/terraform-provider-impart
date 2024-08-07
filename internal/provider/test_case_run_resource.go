package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapiclient "github.com/impart-security/terraform-provider-impart/internal/client"
)

var (
	_ resource.Resource                = &ruleTestcaseRunResource{}
	_ resource.ResourceWithConfigure   = &ruleTestcaseRunResource{}
	_ resource.ResourceWithImportState = &ruleTestcaseRunResource{}
)

// NewRuleTestcaseRunResource is a helper function to simplify the provider implementation.
func NewRuleTestcaseRunResource() resource.Resource {
	return &ruleTestcaseRunResource{}
}

// ruleTestcaseRunResource is the resource implementation.
type ruleTestcaseRunResource struct {
	client *impartAPIClient
}

// Configure adds the provider configured client to the resource.
func (r *ruleTestcaseRunResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *ruleTestcaseRunResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rule_test_case_run"
}

type ruleTestCaseRunModel struct {
	RuleIDs     []types.String `tfsdk:"rule_ids"`
	TestCaseIDs []types.String `tfsdk:"test_case_ids"`
}

func (r ruleTestcaseRunResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"rule_ids": schema.ListAttribute{
				Description: "The HTTP response header values.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"test_case_ids": schema.ListAttribute{
				Description: "The HTTP response header values.",
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *ruleTestcaseRunResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r ruleTestcaseRunResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the test case resource")

	var plan ruleTestCaseModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody, err := toRulesTestCasePostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create the rule test case",
			err.Error(),
		)
		return
	}

	ruleTestCaseRequest := r.client.RulesTestCasesAPI.CreateRulesTestCase(ctx, r.client.OrgID).
		RulesTestCasePostBody(postBody)

	ruleTestCaseResponse, _, err := ruleTestCaseRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the rule test case",
			message,
		)
		return
	}

	state := toRuleTestCaseModel(ruleTestCaseResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Created the test case resource", map[string]any{"success": true})
}

func (r ruleTestcaseRunResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the test case resource")

	var data ruleTestCaseModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ruleTestCaseResponse, httpResp, err := r.client.RulesTestCasesAPI.GetRulesTestCase(ctx, r.client.OrgID, data.ID.ValueString()).
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
			"Unable to read the test case",
			message,
		)
		return
	}

	state := toRuleTestCaseModel(ruleTestCaseResponse, data)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Finished reading the test case resource", map[string]any{"success": true})
}

func (r ruleTestcaseRunResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the test case resource")

	var plan ruleTestCaseModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody, err := toRulesTestCasePostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update the rule test case",
			err.Error(),
		)
		return
	}

	ruleTestCaseRequest := r.client.RulesTestCasesAPI.UpdateRulesTestCase(ctx, r.client.OrgID, plan.ID.ValueString()).
		RulesTestCasePostBody(postBody)

	ruleTestCaseResponse, _, err := ruleTestCaseRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the rule test case",
			message,
		)
		return
	}

	state := toRuleTestCaseModel(ruleTestCaseResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updated the test case resource", map[string]any{"success": true})
}

func (r ruleTestcaseRunResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the test case resource")
	// Retrieve values from a state
	var state ruleTestCaseModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete a test case
	_, err := r.client.RulesTestCasesAPI.DeleteRulesTestCase(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the test case",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the test case resource", map[string]any{"success": true})
}
