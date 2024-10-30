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
	_ resource.Resource                = &tagMetadataResource{}
	_ resource.ResourceWithConfigure   = &tagMetadataResource{}
	_ resource.ResourceWithImportState = &tagMetadataResource{}
)

// NewTagMetadataResource is a helper function to simplify the provider implementation.
func NewTagMetadataResource() resource.Resource {
	return &tagMetadataResource{}
}

// tagMetadataResource is the resource implementation.
type tagMetadataResource struct {
	client *impartAPIClient
}

// Configure adds the provider configured client to the resource.
func (r *tagMetadataResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *tagMetadataResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tag_metadata"
}

type tagMetadataModel struct {
	Name          types.String   `tfsdk:"name"`
	Description   types.String   `tfsdk:"description"`
	RiskStatement types.String   `tfsdk:"risk_statement"`
	ExternalURL   types.String   `tfsdk:"external_url"`
	Labels        []types.String `tfsdk:"labels"`
	//Remediations  []remediationModel `tfsdk:"remediations"`
}

// type remediationModel struct {
// 	Action      types.String `tfsdk:"action"`
// 	Description types.String `tfsdk:"description"`
// }

func (r tagMetadataResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "The tag name.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "The description for the tag.",
				Optional:    true,
			},
			"risk_statement": schema.StringAttribute{
				Description: "The risk statement for the tag.",
				Optional:    true,
			},
			"external_url": schema.StringAttribute{
				Description: "The external URL for the tag.",
				Optional:    true,
			},
			"labels": schema.ListAttribute{
				Description: "The applied labels.",
				ElementType: types.StringType,
				Optional:    true,
			},
			// "remediations": schema.ListNestedAttribute{
			// 	Description: "The remediations for the tag.",
			// 	Optional:    true,
			// 	NestedObject: schema.NestedAttributeObject{
			// 		Attributes: map[string]schema.Attribute{
			// 			"action": schema.StringAttribute{
			// 				Description: "Possible remediation action for the tag.",
			// 				Optional:    true,
			// 			},
			// 			"description": schema.StringAttribute{
			// 				Description: "The description of the remediation.",
			// 				Optional:    true,
			// 			},
			// 		},
			// 	},
			// },
		},
	}
}

func (r *tagMetadataResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}

func (r tagMetadataResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the tag metadata resource")

	var plan tagMetadataModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody, err := toTagMetadataPostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create the tag metadata",
			err.Error(),
		)
		return
	}

	tagsRequest := r.client.TagsAPI.CreateTag(ctx, r.client.OrgID).
		TagPostBody(postBody)

	tagsResponse, _, err := tagsRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the tag metadata",
			message,
		)
		return
	}

	state := toTagMetadataModel(tagsResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Created the tag metadata resource", map[string]any{"success": true})
}

func (r tagMetadataResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the tag metadata resource")

	var data tagMetadataModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tagMetadataResponse, httpResp, err := r.client.TagsAPI.GetTag(ctx, r.client.OrgID, data.Name.ValueString()).
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
			"Unable to read the tag metadata",
			message,
		)
		return
	}

	state := toTagMetadataModel(tagMetadataResponse, data)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Finished reading the tag metadata resource", map[string]any{"success": true})
}

func (r tagMetadataResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the tag metadata resource")

	var plan tagMetadataModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	putBody, err := toTagMetadataPutBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update the tag metadata",
			err.Error(),
		)
		return
	}

	tagRequest := r.client.TagsAPI.UpdateTag(ctx, r.client.OrgID, plan.Name.ValueString()).
		TagPutBody(putBody)

	tagResponse, _, err := tagRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the tag metadata",
			message,
		)
		return
	}

	state := toTagMetadataModel(tagResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updated the tag metadata resource", map[string]any{"success": true})
}

func (r tagMetadataResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the tag metadata resource")
	// Retrieve values from a state
	var state tagMetadataModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete a tag metadata
	_, err := r.client.TagsAPI.DeleteTag(ctx, r.client.OrgID, state.Name.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the tag metadata",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the tag metadata resource", map[string]any{"success": true})
}

func toTagMetadataPostBody(plan tagMetadataModel) (openapiclient.TagPostBody, error) {
	postBody := openapiclient.TagPostBody{
		Name: plan.Name.ValueString(),
	}

	if !plan.Description.IsNull() {
		description := plan.Description.ValueString()
		postBody.Description = &description
	}

	if !plan.ExternalURL.IsNull() {
		externalURL := plan.ExternalURL.ValueString()
		postBody.ExternalUrl = &externalURL
	}

	if !plan.RiskStatement.IsNull() {
		riskStatement := plan.RiskStatement.ValueString()
		postBody.RiskStatement = &riskStatement
	}

	if len(plan.Labels) > 0 {
		labels := make([]string, 0, len(plan.Labels))
		for _, label := range plan.Labels {
			labels = append(labels, label.ValueString())
		}
		postBody.Labels = labels
	}

	// if len(plan.Remediations) > 0 {
	// 	remediations := make([]openapiclient.TagRemediation, 0, len(plan.Remediations))
	// 	for _, remediation := range plan.Remediations {
	// 		remediations = append(remediations, openapiclient.TagRemediation{
	// 			Action:      remediation.Action.ValueString(),
	// 			Description: remediation.Description.ValueString(),
	// 		})
	// 	}
	// 	postBody.Remediations = remediations
	// }

	return postBody, nil
}

func toTagMetadataPutBody(plan tagMetadataModel) (openapiclient.TagPutBody, error) {
	putBody := openapiclient.TagPutBody{}

	if !plan.Description.IsNull() {
		description := plan.Description.ValueString()
		putBody.Description = &description
	}

	if !plan.Description.IsNull() {
		description := plan.Description.ValueString()
		putBody.Description = &description
	}

	if !plan.ExternalURL.IsNull() {
		externalURL := plan.ExternalURL.ValueString()
		putBody.ExternalUrl = &externalURL
	}

	if !plan.RiskStatement.IsNull() {
		riskStatement := plan.RiskStatement.ValueString()
		putBody.RiskStatement = &riskStatement
	}

	if len(plan.Labels) > 0 {
		labels := make([]string, 0, len(plan.Labels))
		for _, label := range plan.Labels {
			labels = append(labels, label.ValueString())
		}
		putBody.Labels = labels
	}

	// if len(plan.Remediations) > 0 {
	// 	remediations := make([]openapiclient.TagRemediation, 0, len(plan.Remediations))
	// 	for _, remediation := range plan.Remediations {
	// 		remediations = append(remediations, openapiclient.TagRemediation{
	// 			Action:      remediation.Action.ValueString(),
	// 			Description: remediation.Description.ValueString(),
	// 		})
	// 	}
	// 	putBody.Remediations = remediations
	// }

	return putBody, nil
}

func toTagMetadataModel(labelResponse *openapiclient.Tag, plan tagMetadataModel) tagMetadataModel {
	label := tagMetadataModel{
		Name: types.StringValue(labelResponse.Name),
	}

	if !plan.Description.IsNull() || labelResponse.Description != "" {
		label.Description = types.StringValue(labelResponse.Description)
	}

	if !plan.ExternalURL.IsNull() || labelResponse.ExternalUrl != "" {
		label.ExternalURL = types.StringValue(labelResponse.ExternalUrl)
	}

	if !plan.RiskStatement.IsNull() || labelResponse.RiskStatement != "" {
		label.RiskStatement = types.StringValue(labelResponse.RiskStatement)
	}

	label.Labels = buildStateList(plan.Labels, labelResponse.Labels)

	// if len(labelResponse.Remediations) > 0 {
	// 	remediations := make([]remediationModel, 0, len(labelResponse.Remediations))
	// 	for _, remediation := range labelResponse.Remediations {
	// 		remediations = append(remediations, remediationModel{
	// 			Action:      types.StringValue(remediation.Action),
	// 			Description: types.StringValue(remediation.Description),
	// 		})
	// 	}
	// 	label.Remediations = remediations
	// }

	return label
}
