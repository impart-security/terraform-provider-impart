package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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

var (
	_ resource.Resource                = &ruleClientIdentifierResource{}
	_ resource.ResourceWithConfigure   = &ruleClientIdentifierResource{}
	_ resource.ResourceWithImportState = &ruleClientIdentifierResource{}
)

// NewRuleClientIdentifierResource is a helper function to simplify the provider implementation.
func NewRuleClientIdentifierResource() resource.Resource {
	return &ruleClientIdentifierResource{}
}

// ruleClientIdentifierResource is the resource implementation.
type ruleClientIdentifierResource struct {
	client *impartAPIClient
}

// Configure adds the provider configured client to the resource.
func (r *ruleClientIdentifierResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *ruleClientIdentifierResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rule_client_identifier"
}

type ruleClientIdentifierModel struct {
	ID          types.String     `tfsdk:"id"`
	Name        types.String     `tfsdk:"name"`
	Description types.String     `tfsdk:"description"`
	StorageID   types.String     `tfsdk:"storage_id"`
	HashFields  []hashFieldModel `tfsdk:"hash_fields"`
}

type hashFieldModel struct {
	Field types.String `tfsdk:"field"`
	Key   types.String `tfsdk:"key"`
}

func (r ruleClientIdentifierResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for the rule client identifier.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name for the rule client identifier.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description for the rule client identifier.",
				Optional:    true,
			},
			"storage_id": schema.StringAttribute{
				Description: "The storage id for the rule client identifier.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"hash_fields": schema.ListNestedAttribute{
				Description: "The hash fields for the rule client identifier.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"field": schema.StringAttribute{
							Required:    true,
							Description: "The hash field.",
						},
						"key": schema.StringAttribute{
							Optional:    true,
							Description: "The hash field key.",
						},
					},
				},
			},
		},
	}
}

func (r *ruleClientIdentifierResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r ruleClientIdentifierResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the client identifier resource")

	var plan ruleClientIdentifierModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody, err := toClientIdentifierPostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create the client identifier resource",
			err.Error(),
		)
		return
	}

	clientIdentifierRequest := r.client.RuleClientIdentifiersAPI.CreateRuleClientIdentifier(ctx, r.client.OrgID).
		RuleClientIdentifierPostBody(postBody)

	clientIdentifierResponse, _, err := clientIdentifierRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the client identifier resource",
			message,
		)
		return
	}

	state := toRuleClientIdentifierModel(clientIdentifierResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Created the client identifier resource", map[string]any{"success": true})
}

func (r ruleClientIdentifierResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the client identifier resource")

	var data ruleClientIdentifierModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clientIdentifierResponse, httpResp, err := r.client.RuleClientIdentifiersAPI.GetRuleClientIdentifier(ctx, r.client.OrgID, data.ID.ValueString()).
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
			"Unable to read the client identifier resource "+data.ID.ValueString(),
			message,
		)
		return
	}

	state := toRuleClientIdentifierModel(clientIdentifierResponse, data)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Finished reading the client identifier resource", map[string]any{"success": true})
}

func (r ruleClientIdentifierResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the client identifier resource")

	var plan ruleClientIdentifierModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody, err := toClientIdentifierPostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update the client identifier resource",
			err.Error(),
		)
		return
	}

	clientIdentifierRequest := r.client.RuleClientIdentifiersAPI.UpdateRuleClientIdentifier(ctx, r.client.OrgID, plan.ID.ValueString()).
		RuleClientIdentifierPostBody(postBody)

	clientIdentifierResponse, _, err := clientIdentifierRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the client identifier resource",
			message,
		)
		return
	}

	state := toRuleClientIdentifierModel(clientIdentifierResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updated the client identifier resource", map[string]any{"success": true})
}

func (r ruleClientIdentifierResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the client identifier resource")
	// Retrieve values from a state
	var state ruleClientIdentifierModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete a client identifier
	_, err := r.client.RuleClientIdentifiersAPI.DeleteRuleClientIdentifier(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the client identifier resource",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the client identifier resource", map[string]any{"success": true})
}

func toClientIdentifierPostBody(plan ruleClientIdentifierModel) (openapiclient.RuleClientIdentifierPostBody, error) {
	postBody := openapiclient.RuleClientIdentifierPostBody{
		Name: plan.Name.ValueString(),
	}

	if !plan.Description.IsNull() {
		description := plan.Description.ValueString()
		postBody.Description = &description
	}

	if !plan.StorageID.IsNull() {
		storageID := plan.StorageID.ValueString()
		postBody.StorageId = &storageID
	}

	if len(plan.HashFields) > 0 {
		hashFields := make([]openapiclient.RuleClientIdentifierHashField, len(plan.HashFields))
		for i, field := range plan.HashFields {
			hashFields[i] = openapiclient.RuleClientIdentifierHashField{
				Field: field.Field.ValueString(),
			}

			if !field.Key.IsNull() {
				key := field.Key.ValueString()
				hashFields[i].Key = &key
			}
		}
		postBody.HashFields = hashFields
	}

	return postBody, nil
}

func toRuleClientIdentifierModel(clientIdentifierResponse *openapiclient.RuleClientIdentifier, plan ruleClientIdentifierModel) ruleClientIdentifierModel {
	ruleClientIdentifier := ruleClientIdentifierModel{
		ID:   types.StringValue(clientIdentifierResponse.Id),
		Name: types.StringValue(clientIdentifierResponse.Name),
	}

	if !plan.Description.IsNull() || clientIdentifierResponse.Description != "" {
		ruleClientIdentifier.Description = types.StringValue(clientIdentifierResponse.Description)
	}

	if clientIdentifierResponse.StorageId != nil {
		ruleClientIdentifier.StorageID = types.StringValue(*clientIdentifierResponse.StorageId)
	}

	if len(clientIdentifierResponse.HashFields) > 0 {
		hashFields := make([]hashFieldModel, len(clientIdentifierResponse.HashFields))
		for i, field := range clientIdentifierResponse.HashFields {
			hashFields[i] = hashFieldModel{
				Field: types.StringValue(field.Field),
			}

			if field.Key != nil {
				hashFields[i].Key = types.StringValue(*field.Key)
			}
		}
		ruleClientIdentifier.HashFields = hashFields
	}

	return ruleClientIdentifier
}

func (r *ruleClientIdentifierResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var plan ruleClientIdentifierModel
	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
