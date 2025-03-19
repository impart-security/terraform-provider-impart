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
	_ resource.Resource                = &externalLinkResource{}
	_ resource.ResourceWithConfigure   = &externalLinkResource{}
	_ resource.ResourceWithImportState = &externalLinkResource{}
)

// NewExternalLinkResource is a helper function to simplify the provider implementation.
func NewExternalLinkResource() resource.Resource {
	return &externalLinkResource{}
}

// externalLinkResource is the resource implementation.
type externalLinkResource struct {
	client *impartAPIClient
}

// Configure adds the provider configured client to the resource.
func (r *externalLinkResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *externalLinkResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_external_link"
}

type externalLinkModel struct {
	ID              types.String   `tfsdk:"id"`
	Name            types.String   `tfsdk:"name"`
	Description     types.String   `tfsdk:"description"`
	SpecsIDs        []types.String `tfsdk:"spec_ids"`
	Entity          types.String   `tfsdk:"entity"`
	JSONPathElement types.String   `tfsdk:"json_path_element"`
	URL             types.String   `tfsdk:"url"`
	Vendor          types.String   `tfsdk:"vendor"`
}

func (r externalLinkResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for this list.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the external link.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the external link.",
				Optional:    true,
			},
			"entity": schema.StringAttribute{
				Description: "The entity to which the links should be applied.",
				Required:    true,
			},
			"json_path_element": schema.StringAttribute{
				Description: "A JSONPath to the element for which this link should apply (e.g. $.client_ip.address).",
				Required:    true,
			},
			"url": schema.StringAttribute{
				Description: "The external URL template with JSONPath element variables.",
				Required:    true,
			},
			"vendor": schema.StringAttribute{
				Description: "The vendor for the external link.",
				Required:    true,
			},
			"spec_ids": schema.ListAttribute{
				Description: "A list of spec IDs this external link applies to (empty means all).",
				ElementType: types.StringType,
				Optional:    true,
			},
		},
	}
}

func (r *externalLinkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r externalLinkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the external link resource")

	var plan externalLinkModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody, err := toExternalLinkPostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create the external link",
			err.Error(),
		)
		return
	}

	externalLinkRequest := r.client.ExternalLinksAPI.CreateExternalLink(ctx, r.client.OrgID).
		ExternalLinkPostBody(postBody)

	externalLinkResponse, _, err := externalLinkRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the external link",
			message,
		)
		return
	}

	state := toExternalLinkModel(externalLinkResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Created the external link resource", map[string]any{"success": true})
}

func (r externalLinkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the external link resource")

	var data externalLinkModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	externalLinkResponse, httpResp, err := r.client.ExternalLinksAPI.GetExternalLink(ctx, r.client.OrgID, data.ID.ValueString()).
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
			"Unable to read the external link ",
			message,
		)
		return
	}

	state := toExternalLinkModel(externalLinkResponse, data)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Finished reading the external link resource", map[string]any{"success": true})
}

func (r externalLinkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the external link resource")

	var plan externalLinkModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody, err := toExternalLinkPostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update the external link",
			err.Error(),
		)
		return
	}

	externalLinkRequest := r.client.ExternalLinksAPI.UpdateExternalLink(ctx, r.client.OrgID, plan.ID.ValueString()).
		ExternalLinkPostBody(postBody)

	externalLinkResponse, _, err := externalLinkRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the external link",
			message,
		)
		return
	}

	state := toExternalLinkModel(externalLinkResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updated the external link resource", map[string]any{"success": true})
}

func (r externalLinkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the external link resource")
	// Retrieve values from a state
	var state externalLinkModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete the external link
	_, err := r.client.ExternalLinksAPI.DeleteExternalLink(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the external link",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the external link resource", map[string]any{"success": true})
}

func toExternalLinkPostBody(plan externalLinkModel) (openapiclient.ExternalLinkPostBody, error) {
	postBody := openapiclient.ExternalLinkPostBody{
		Name:            plan.Name.ValueString(),
		JsonPathElement: plan.JSONPathElement.ValueString(),
		Url:             plan.URL.ValueString(),
		Vendor:          plan.Vendor.ValueString(),
	}

	if !plan.Description.IsNull() {
		description := plan.Description.ValueString()
		postBody.Description = &description
	}

	if len(plan.SpecsIDs) > 0 {
		specIDs := make([]string, 0, len(plan.SpecsIDs))
		for _, specID := range plan.SpecsIDs {
			specIDs = append(specIDs, specID.ValueString())
		}
		postBody.SpecIds = specIDs
	}

	if !plan.Entity.IsNull() {
		postBody.Entity = plan.Entity.ValueString()
	}

	return postBody, nil
}

func toExternalLinkModel(externalLinkResource *openapiclient.ExternalLink, plan externalLinkModel) externalLinkModel {
	externalLink := externalLinkModel{
		ID:              types.StringValue(externalLinkResource.Id),
		Name:            types.StringValue(externalLinkResource.Name),
		Entity:          types.StringValue(externalLinkResource.Entity),
		JSONPathElement: types.StringValue(externalLinkResource.JsonPathElement),
		URL:             types.StringValue(externalLinkResource.Url),
		Vendor:          types.StringValue(externalLinkResource.Vendor),
	}

	if !plan.Description.IsNull() || externalLinkResource.Description != "" {
		externalLink.Description = types.StringValue(externalLinkResource.Description)
	}

	externalLink.SpecsIDs = buildStateList(plan.SpecsIDs, externalLinkResource.SpecIds)

	return externalLink
}

func (r *externalLinkResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var plan externalLinkModel
	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
