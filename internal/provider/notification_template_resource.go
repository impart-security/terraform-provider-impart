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

	openapiclient "github.com/impart-security/terraform-provider-impart/internal/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &notificationTemplateResource{}
	_ resource.ResourceWithConfigure   = &notificationTemplateResource{}
	_ resource.ResourceWithImportState = &notificationTemplateResource{}
)

// NewNotificationTemplateResource is a helper function to simplify the provider implementation.
func NewNotificationTemplateResource() resource.Resource {
	return &notificationTemplateResource{}
}

// notificationTemplateResource is the resource implementation.
type notificationTemplateResource struct {
	client *impartAPIClient
}

// notificationTemplateResourceModel maps the resource schema data.
type notificationTemplateResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Payload     types.String `tfsdk:"payload"`
	Subject     types.String `tfsdk:"subject"`
	Destination []string     `tfsdk:"destination"`
	ConnectorID types.String `tfsdk:"connector_id"`
}

// Configure adds the provider configured client to the resource.
func (r *notificationTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *notificationTemplateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_notification_template"
}

// Schema defines the schema for the resource.
func (r *notificationTemplateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage an notification template.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for this notification template.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name for this notification template.",
				Required:    true,
			},
			"payload": schema.StringAttribute{
				Description: "The payload message that will be sent to the Third Party API.",
				Required:    true,
			},
			"subject": schema.StringAttribute{
				Description: "The subject message that will be sent to the Third Party API.",
				Required:    true,
			},
			"destination": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "An array of destination ids to which the payloads will be sent.",
				Required:    true,
			},
			"connector_id": schema.StringAttribute{
				Description: "The connector id.",
				Required:    true,
			},
		},
	}
}

func (r *notificationTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Create a new resource.
func (r *notificationTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the notification template resource")

	// Retrieve values from plan
	var plan notificationTemplateResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody := openapiclient.NotificationTemplatePostBody{
		ConnectorId: plan.ConnectorID.ValueString(),
		Name:        plan.Name.ValueString(),
		Payload:     plan.Payload.ValueString(),
		Subject:     plan.Subject.ValueString(),
		Destination: plan.Destination,
	}

	notificationTemplateResponse, _, err := r.client.NotificationTemplatesAPI.CreateNotificationTemplate(ctx, r.client.OrgID).NotificationTemplatePostBody(postBody).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the notification template",
			message,
		)
		return
	}

	// Map response body to model
	plan.ID = types.StringValue(notificationTemplateResponse.Id)
	plan.Name = types.StringValue(notificationTemplateResponse.Name)
	plan.Payload = types.StringValue(notificationTemplateResponse.Payload)
	plan.Subject = types.StringValue(notificationTemplateResponse.Subject)
	plan.Destination = notificationTemplateResponse.Destination

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Created the notification template", map[string]any{"success": true})
}

// Read resource information.
func (r *notificationTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the notification template resource")
	// Get current state
	var state notificationTemplateResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	notificationTemplateResponse, httpResp, err := r.client.NotificationTemplatesAPI.GetNotificationTemplate(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
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
			"Unable to read the notification template",
			message,
		)
		return
	}

	// Map response body to model
	state.ID = types.StringValue(notificationTemplateResponse.Id)
	state.Name = types.StringValue(notificationTemplateResponse.Name)
	state.Payload = types.StringValue(notificationTemplateResponse.Payload)
	state.Subject = types.StringValue(notificationTemplateResponse.Subject)
	state.Destination = notificationTemplateResponse.Destination

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Finished reading the notification template resource", map[string]any{"success": true})
}

func (r *notificationTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the notification template resource")
	// Retrieve values from plan
	var plan notificationTemplateResourceModel
	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody := openapiclient.NotificationTemplatePostBody{
		ConnectorId: plan.ConnectorID.ValueString(),
		Name:        plan.Name.ValueString(),
		Payload:     plan.Payload.ValueString(),
		Subject:     plan.Subject.ValueString(),
		Destination: plan.Destination,
	}

	notificationTemplateRequest := r.client.NotificationTemplatesAPI.UpdateNotificationTemplate(ctx, r.client.OrgID, plan.ID.ValueString()).
		NotificationTemplatePostBody(postBody)

	// update notification template
	notificationTemplateResponse, _, err := notificationTemplateRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the notification template",
			message,
		)
		return
	}

	// Overwrite the notification template with refreshed state
	state := notificationTemplateResourceModel{
		ID:          types.StringValue(notificationTemplateResponse.Id),
		Name:        types.StringValue(notificationTemplateResponse.Name),
		Payload:     types.StringValue(notificationTemplateResponse.Payload),
		Subject:     types.StringValue(notificationTemplateResponse.Subject),
		Destination: notificationTemplateResponse.Destination,
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updated the notification template resource", map[string]any{"success": true})
}

func (r *notificationTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the notification template resource")
	// Retrieve values from state
	var state notificationTemplateResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// delete notification template
	_, err := r.client.NotificationTemplatesAPI.DeleteNotificationTemplate(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the notification template",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the notification template resource", map[string]any{"success": true})
}
