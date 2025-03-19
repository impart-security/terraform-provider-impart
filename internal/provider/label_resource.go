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
	_ resource.Resource                = &labelResource{}
	_ resource.ResourceWithConfigure   = &labelResource{}
	_ resource.ResourceWithImportState = &labelResource{}
)

// NewLabelResource is a helper function to simplify the provider implementation.
func NewLabelResource() resource.Resource {
	return &labelResource{}
}

// labelResource is the resource implementation.
type labelResource struct {
	client *impartAPIClient
}

// Configure adds the provider configured client to the resource.
func (r *labelResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *labelResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_label"
}

type labelModel struct {
	Slug        types.String `tfsdk:"slug"`
	DisplyaName types.String `tfsdk:"display_name"`
	Description types.String `tfsdk:"description"`
	Color       types.String `tfsdk:"color"`
}

func (r labelResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"slug": schema.StringAttribute{
				Description: "Slug of the label.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"display_name": schema.StringAttribute{
				Description: "The display name of the label.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the label.",
				Optional:    true,
			},
			"color": schema.StringAttribute{
				Description: "The color of the label.",
				Optional:    true,
			},
		},
	}
}

func (r *labelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("slug"), req, resp)
}

func (r labelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the label resource")

	var plan labelModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody, err := toLabelPostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create the label",
			err.Error(),
		)
		return
	}

	labelRequest := r.client.LabelsAPI.CreateLabel(ctx, r.client.OrgID).
		LabelPostBody(postBody)

	labelResponse, _, err := labelRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the label",
			message,
		)
		return
	}

	state := toLabelModel(labelResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Created the label resource", map[string]any{"success": true})
}

func (r labelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the label resource")

	var data labelModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	labelResponse, httpResp, err := r.client.LabelsAPI.GetLabel(ctx, r.client.OrgID, data.Slug.ValueString()).
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
			"Unable to read the label "+data.Slug.ValueString(),
			message,
		)
		return
	}

	state := toLabelModel(labelResponse, data)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Finished reading the label resource", map[string]any{"success": true})
}

func (r labelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the label resource")

	var plan labelModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody, err := toLabelPutBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update the label",
			err.Error(),
		)
		return
	}

	labelRequest := r.client.LabelsAPI.UpdateLabel(ctx, r.client.OrgID, plan.Slug.ValueString()).
		LabelPutBody(postBody)

	labelResponse, _, err := labelRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the label",
			message,
		)
		return
	}

	state := toLabelModel(labelResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updated the label resource", map[string]any{"success": true})
}

func (r labelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the label resource")
	// Retrieve values from a state
	var state labelModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete a label
	_, err := r.client.LabelsAPI.DeleteLabel(ctx, r.client.OrgID, state.Slug.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the label",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the label resource", map[string]any{"success": true})
}

func toLabelPostBody(plan labelModel) (openapiclient.LabelPostBody, error) {
	postBody := openapiclient.LabelPostBody{
		Slug: plan.Slug.ValueString(),
	}

	if !plan.Description.IsNull() {
		description := plan.Description.ValueString()
		postBody.Description = &description
	}

	if !plan.DisplyaName.IsNull() {
		displayName := plan.DisplyaName.ValueString()
		postBody.DisplayName = &displayName
	}

	if !plan.Color.IsNull() {
		color, err := openapiclient.NewLabelColorFromValue(plan.Color.ValueString())
		if err != nil {
			return postBody, fmt.Errorf("unable to create label color: %w", err)
		}
		postBody.Color = color
	}

	return postBody, nil
}

func toLabelPutBody(plan labelModel) (openapiclient.LabelPutBody, error) {
	putBody := openapiclient.LabelPutBody{}

	if !plan.Description.IsNull() {
		description := plan.Description.ValueString()
		putBody.Description = &description
	}

	if !plan.DisplyaName.IsNull() {
		displayName := plan.DisplyaName.ValueString()
		putBody.DisplayName = &displayName
	}

	if !plan.Color.IsNull() {
		color, err := openapiclient.NewLabelColorFromValue(plan.Color.ValueString())
		if err != nil {
			return putBody, fmt.Errorf("unable to create label color: %w", err)
		}
		putBody.Color = color
	}

	return putBody, nil
}

func toLabelModel(labelResponse *openapiclient.Label, plan labelModel) labelModel {
	label := labelModel{
		Slug: types.StringValue(labelResponse.Slug),
	}

	if !plan.Color.IsNull() || labelResponse.Color != openapiclient.GRAY {
		label.Color = types.StringValue(string(labelResponse.Color))
	}

	if !plan.DisplyaName.IsNull() || labelResponse.DisplayName != labelResponse.Slug {
		label.DisplyaName = types.StringValue(labelResponse.DisplayName)
	}

	if !plan.Description.IsNull() || labelResponse.Description != "" {
		label.Description = types.StringValue(labelResponse.Description)
	}

	return label
}

func (r *labelResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var plan labelModel
	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !plan.Color.IsNull() {
		_, err := openapiclient.NewLabelColorFromValue(plan.Color.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Invalid Color Attribute",
				err.Error(),
			)
		}
	}
}
