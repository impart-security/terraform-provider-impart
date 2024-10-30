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

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &monitorResource{}
	_ resource.ResourceWithConfigure   = &monitorResource{}
	_ resource.ResourceWithImportState = &monitorResource{}
)

// NewMonitorResource is a helper function to simplify the provider implementation.
func NewMonitorResource() resource.Resource {
	return &monitorResource{}
}

// monitorResource is the resource implementation.
type monitorResource struct {
	client *impartAPIClient
}

// monitorResourceModel maps the resource schema data.
type monitorResourceModel struct {
	ID                      types.String     `tfsdk:"id"`
	Name                    types.String     `tfsdk:"name"`
	Description             types.String     `tfsdk:"description"`
	Conditions              []conditionModel `tfsdk:"conditions"`
	NotificationTemplateIDs []string         `tfsdk:"notification_template_ids"`
	Labels                  []types.String   `tfsdk:"labels"`
}

type conditionModel struct {
	Threshold  types.Int64           `tfsdk:"threshold"`
	Comparator types.String          `tfsdk:"comparator"`
	TimePeriod types.Int64           `tfsdk:"time_period"`
	Delay      types.Int64           `tfsdk:"delay"`
	Details    conditionDetailsModel `tfsdk:"details"`
}

type conditionDetailsModel struct {
	Type        types.String `tfsdk:"type"`
	Action      *string      `tfsdk:"action"`
	SubjectType *string      `tfsdk:"subject_type"`
	ActorType   *string      `tfsdk:"actor_type"`
	Tag         *string      `tfsdk:"tag"`
}

// Configure adds the provider configured client to the resource.
func (r *monitorResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *monitorResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_monitor"
}

// Schema defines the schema for the resource.
func (r *monitorResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage a monitor.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for this monitor.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name for this monitor.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description for this monitor.",
				Required:    true,
			},
			"conditions": schema.ListNestedAttribute{
				Description: "An array of conditions for which the monitor will trigger.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"threshold": schema.Int64Attribute{
							Required:    true,
							Description: "Number of occurrences that need to execute to have this condition be true.",
						},
						"comparator": schema.StringAttribute{
							Required:    true,
							Description: "Greater than, equal to, or less than (should be one of 'gt', 'lt', or 'eq')",
						},
						"time_period": schema.Int64Attribute{
							Required:    true,
							Description: "In milliseconds, the time span from now until when we should be counting events (for example, 60000 is all events in the last minute).",
						},
						"delay": schema.Int64Attribute{
							Required:    true,
							Description: "In milliseconds, the offset from now() for the time window.",
						},
						"details": schema.SingleNestedAttribute{
							Required: true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Required:    true,
									Description: "The type of monitor (should be one of 'event' or 'metric'",
								},

								// Event monitor fields
								"action": schema.StringAttribute{
									Optional:    true,
									Description: "Strictly for event type monitors. A slug of the action the monitor is tracking.",
								},
								"subject_type": schema.StringAttribute{
									Optional:    true,
									Description: "Strictly for event type monitors. A slug of the subject type the monitor is tracking.",
								},
								"actor_type": schema.StringAttribute{
									Optional:    true,
									Description: "Strictly for event type monitors. A slug of the actor type the monitor is tracking.",
								},

								// Metric monitor fields
								"tag": schema.StringAttribute{
									Optional:    true,
									Description: "Strictly for metric type monitors. The tag the monitor is tracking.",
								},
							},
						},
					},
				},
			},
			"notification_template_ids": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "An array of notification template ids for the templates that will send notifications to their respective connectors.",
				Required:    true,
			},
			"labels": schema.ListAttribute{
				Description: "The applied labels.",
				ElementType: types.StringType,
				Optional:    true,
			},
		},
	}
}

func (r *monitorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Create a new resource.
func (r *monitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the monitor resource")

	// Retrieve values from plan
	var plan monitorResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	conditions := make([]openapiclient.EventMonitorConditionPostBody, len(plan.Conditions))
	for i := range plan.Conditions {
		var eventDetails *openapiclient.EventMonitorConditionEvent
		if plan.Conditions[i].Details.Type.ValueString() == "event" {
			if plan.Conditions[i].Details.Tag != nil {
				resp.Diagnostics.AddError(
					fmt.Sprintf("conditions[%d].details.tag", i),
					"field is strictly for metric type monitors",
				)
				return
			}
			eventDetails = &openapiclient.EventMonitorConditionEvent{
				Type:        plan.Conditions[i].Details.Type.ValueString(),
				Action:      plan.Conditions[i].Details.Action,
				ActorType:   plan.Conditions[i].Details.ActorType,
				SubjectType: plan.Conditions[i].Details.SubjectType,
			}
		}

		var metricDetails *openapiclient.EventMonitorConditionMetric
		if plan.Conditions[i].Details.Type.ValueString() == "metric" {
			if plan.Conditions[i].Details.Action != nil {
				resp.Diagnostics.AddError(
					fmt.Sprintf("conditions[%d].details.action", i),
					"field is strictly for event type monitors",
				)
				return
			}

			if plan.Conditions[i].Details.SubjectType != nil {
				resp.Diagnostics.AddError(
					fmt.Sprintf("conditions[%d].details.subject_type", i),
					"field is strictly for event type monitors",
				)
				return
			}

			if plan.Conditions[i].Details.ActorType != nil {
				resp.Diagnostics.AddError(
					fmt.Sprintf("conditions[%d].details.actor_type", i),
					"field is strictly for event type monitors",
				)
				return
			}
			metricDetails = &openapiclient.EventMonitorConditionMetric{
				Type: plan.Conditions[i].Details.Type.ValueString(),
				Tag:  plan.Conditions[i].Details.Tag,
			}
		}

		conditions[i] = openapiclient.EventMonitorConditionPostBody{
			Threshold:  int32(plan.Conditions[i].Threshold.ValueInt64()),
			Comparator: plan.Conditions[i].Comparator.ValueString(),
			Delay:      plan.Conditions[i].Delay.ValueInt64(),
			TimePeriod: plan.Conditions[i].TimePeriod.ValueInt64(),
			Details: openapiclient.EventMonitorConditionDetails{
				EventMonitorConditionEvent:  eventDetails,
				EventMonitorConditionMetric: metricDetails,
			},
		}
	}

	postBody := openapiclient.EventMonitorPostBody{
		Name:                    plan.Name.ValueString(),
		Description:             plan.Description.ValueString(),
		Conditions:              conditions,
		NotificationTemplateIds: plan.NotificationTemplateIDs,
	}

	if len(plan.Labels) > 0 {
		labels := make([]string, len(plan.Labels))
		for i, label := range plan.Labels {
			labels[i] = label.ValueString()
		}
		postBody.Labels = labels
	}

	monitorResponse, _, err := r.client.EventMonitorsAPI.CreateEventMonitor(ctx, r.client.OrgID).EventMonitorPostBody(postBody).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the monitor",
			message,
		)
		return
	}

	responseConditions := make([]conditionModel, len(monitorResponse.Conditions))
	for i := range monitorResponse.Conditions {
		var monitorType string
		var action *string
		var actorType *string
		var subjectType *string
		var tag *string
		if monitorResponse.Conditions[i].Details.EventMonitorConditionEvent != nil {
			monitorType = monitorResponse.Conditions[i].Details.EventMonitorConditionEvent.Type
			action = monitorResponse.Conditions[i].Details.EventMonitorConditionEvent.Action
			actorType = monitorResponse.Conditions[i].Details.EventMonitorConditionEvent.ActorType
			subjectType = monitorResponse.Conditions[i].Details.EventMonitorConditionEvent.SubjectType
		} else {
			monitorType = monitorResponse.Conditions[i].Details.EventMonitorConditionMetric.Type
			tag = monitorResponse.Conditions[i].Details.EventMonitorConditionMetric.Tag
		}

		responseConditions[i] = conditionModel{
			Threshold:  types.Int64Value(int64(monitorResponse.Conditions[i].Threshold)),
			Comparator: types.StringValue(monitorResponse.Conditions[i].Comparator),
			Delay:      types.Int64Value(monitorResponse.Conditions[i].Delay),
			TimePeriod: types.Int64Value(monitorResponse.Conditions[i].TimePeriod),
			Details: conditionDetailsModel{
				Type:        types.StringValue(monitorType),
				Action:      action,
				ActorType:   actorType,
				SubjectType: subjectType,
				Tag:         tag,
			},
		}
	}
	// Map response body to model
	plan.ID = types.StringValue(monitorResponse.Id)
	plan.Name = types.StringValue(monitorResponse.Name)
	plan.Conditions = responseConditions
	plan.NotificationTemplateIDs = monitorResponse.NotificationTemplateIds

	plan.Labels = buildStateList(plan.Labels, monitorResponse.Labels)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Created the monitor", map[string]any{"success": true})
}

// Read resource information.
func (r *monitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the monitor resource")
	// Get current state
	var state monitorResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	monitorResponse, httpResp, err := r.client.EventMonitorsAPI.GetEventMonitor(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
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
			"Unable to read the monitor",
			message,
		)
		return
	}

	responseConditions := make([]conditionModel, len(monitorResponse.Conditions))
	for i := range monitorResponse.Conditions {
		var monitorType string
		var action *string
		var actorType *string
		var subjectType *string
		var tag *string
		if monitorResponse.Conditions[i].Details.EventMonitorConditionEvent != nil {
			monitorType = monitorResponse.Conditions[i].Details.EventMonitorConditionEvent.Type
			action = monitorResponse.Conditions[i].Details.EventMonitorConditionEvent.Action
			actorType = monitorResponse.Conditions[i].Details.EventMonitorConditionEvent.ActorType
			subjectType = monitorResponse.Conditions[i].Details.EventMonitorConditionEvent.SubjectType
		} else {
			monitorType = monitorResponse.Conditions[i].Details.EventMonitorConditionMetric.Type
			tag = monitorResponse.Conditions[i].Details.EventMonitorConditionMetric.Tag
		}

		responseConditions[i] = conditionModel{
			Threshold:  types.Int64Value(int64(monitorResponse.Conditions[i].Threshold)),
			Comparator: types.StringValue(monitorResponse.Conditions[i].Comparator),
			Delay:      types.Int64Value(monitorResponse.Conditions[i].Delay),
			TimePeriod: types.Int64Value(monitorResponse.Conditions[i].TimePeriod),
			Details: conditionDetailsModel{
				Type:        types.StringValue(monitorType),
				Action:      action,
				ActorType:   actorType,
				SubjectType: subjectType,
				Tag:         tag,
			},
		}
	}

	// Map response body to model
	state.ID = types.StringValue(monitorResponse.Id)
	state.Name = types.StringValue(monitorResponse.Name)
	state.Conditions = responseConditions
	state.NotificationTemplateIDs = monitorResponse.NotificationTemplateIds
	state.Labels = buildStateList(state.Labels, monitorResponse.Labels)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Finished reading the monitor resource", map[string]any{"success": true})
}

func (r *monitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the monitor resource")
	// Retrieve values from plan
	var plan monitorResourceModel
	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	conditions := make([]openapiclient.EventMonitorConditionPostBody, len(plan.Conditions))
	for i := range plan.Conditions {
		var eventDetails *openapiclient.EventMonitorConditionEvent
		if plan.Conditions[i].Details.Type.ValueString() == "event" {
			eventDetails = &openapiclient.EventMonitorConditionEvent{
				Type:        plan.Conditions[i].Details.Type.ValueString(),
				Action:      plan.Conditions[i].Details.Action,
				ActorType:   plan.Conditions[i].Details.ActorType,
				SubjectType: plan.Conditions[i].Details.SubjectType,
			}
		}

		var metricDetails *openapiclient.EventMonitorConditionMetric
		if plan.Conditions[i].Details.Type.ValueString() == "metric" {
			metricDetails = &openapiclient.EventMonitorConditionMetric{
				Type: plan.Conditions[i].Details.Type.ValueString(),
				Tag:  plan.Conditions[i].Details.Tag,
			}
		}
		conditions[i] = openapiclient.EventMonitorConditionPostBody{
			Threshold:  int32(plan.Conditions[i].Threshold.ValueInt64()),
			Comparator: plan.Conditions[i].Comparator.ValueString(),
			Delay:      plan.Conditions[i].Delay.ValueInt64(),
			TimePeriod: plan.Conditions[i].TimePeriod.ValueInt64(),
			Details: openapiclient.EventMonitorConditionDetails{
				EventMonitorConditionEvent:  eventDetails,
				EventMonitorConditionMetric: metricDetails,
			},
		}
	}

	postBody := openapiclient.EventMonitorPostBody{
		Name:                    plan.Name.ValueString(),
		Description:             plan.Description.ValueString(),
		Conditions:              conditions,
		NotificationTemplateIds: plan.NotificationTemplateIDs,
	}

	if len(plan.Labels) > 0 {
		labels := make([]string, len(plan.Labels))
		for i, label := range plan.Labels {
			labels[i] = label.ValueString()
		}
		postBody.Labels = labels
	}

	monitorRequest := r.client.EventMonitorsAPI.UpdateEventMonitor(ctx, r.client.OrgID, plan.ID.ValueString()).
		EventMonitorPostBody(postBody)

	// update monitor
	monitorResponse, _, err := monitorRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the monitor",
			message,
		)
		return
	}

	responseConditions := make([]conditionModel, len(monitorResponse.Conditions))
	for i := range monitorResponse.Conditions {
		var monitorType string
		var action *string
		var actorType *string
		var subjectType *string
		var tag *string
		if monitorResponse.Conditions[i].Details.EventMonitorConditionEvent != nil {
			monitorType = monitorResponse.Conditions[i].Details.EventMonitorConditionEvent.Type
			action = monitorResponse.Conditions[i].Details.EventMonitorConditionEvent.Action
			actorType = monitorResponse.Conditions[i].Details.EventMonitorConditionEvent.ActorType
			subjectType = monitorResponse.Conditions[i].Details.EventMonitorConditionEvent.SubjectType
		} else {
			monitorType = monitorResponse.Conditions[i].Details.EventMonitorConditionMetric.Type
			tag = monitorResponse.Conditions[i].Details.EventMonitorConditionMetric.Tag
		}

		responseConditions[i] = conditionModel{
			Threshold:  types.Int64Value(int64(monitorResponse.Conditions[i].Threshold)),
			Comparator: types.StringValue(monitorResponse.Conditions[i].Comparator),
			Delay:      types.Int64Value(monitorResponse.Conditions[i].Delay),
			TimePeriod: types.Int64Value(monitorResponse.Conditions[i].TimePeriod),
			Details: conditionDetailsModel{
				Type:        types.StringValue(monitorType),
				Action:      action,
				ActorType:   actorType,
				SubjectType: subjectType,
				Tag:         tag,
			},
		}
	}

	// Overwrite the monitor with refreshed state
	state := monitorResourceModel{
		ID:                      types.StringValue(monitorResponse.Id),
		Name:                    types.StringValue(monitorResponse.Name),
		Description:             types.StringValue(monitorResponse.Description),
		NotificationTemplateIDs: monitorResponse.NotificationTemplateIds,
		Conditions:              responseConditions,
	}
	state.Labels = buildStateList(plan.Labels, monitorResponse.Labels)

	// Set refreshed state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updated the monitor resource", map[string]any{"success": true})
}

func (r *monitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the monitor resource")
	// Retrieve values from state
	var state monitorResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// delete monitor
	_, err := r.client.EventMonitorsAPI.DeleteEventMonitor(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the monitor",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the monitor resource", map[string]any{"success": true})
}
