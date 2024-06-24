package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/netip"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"go4.org/netipx"

	openapiclient "github.com/impart-security/terraform-provider-impart/internal/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &ListResource{}
	_ resource.ResourceWithConfigure   = &ListResource{}
	_ resource.ResourceWithImportState = &ListResource{}
)

// NewListResource is a helper function to simplify the provider implementation.
func NewListResource() resource.Resource {
	return &ListResource{}
}

// ListResource is the resource implementation.
type ListResource struct {
	client *impartAPIClient
}

// listResourceModel maps the resource schema data.
type listResourceModel struct {
	ID      types.String    `tfsdk:"id"`
	Name    types.String    `tfsdk:"name"`
	Kind    types.String    `tfsdk:"kind"`
	Subkind types.String    `tfsdk:"subkind"`
	Items   []listItemModel `tfsdk:"items"`
}

type listItemModel struct {
	Value      types.String `tfsdk:"value"`
	Expiration types.String `tfsdk:"expiration"`
}

// Configure adds the provider configured client to the resource.
func (r *ListResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *ListResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_list"
}

// Schema defines the schema for the resource.
func (r *ListResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage a list.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for this list.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name for this list.",
				Required:    true,
			},
			"kind": schema.StringAttribute{
				Description: "The list kind.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"subkind": schema.StringAttribute{
				Description: "The list subkind.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The list items.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"value": schema.StringAttribute{
							Required:    true,
							Description: "The list item value.",
						},
						"expiration": schema.StringAttribute{
							Optional:    true,
							Description: "The list item expiration.",
							Validators: []validator.String{
								dateTimeNotPast("expiration"),
							},
						},
					},
				},
				Optional: true,
				Validators: []validator.List{
					uniqueValue(),
				},
				PlanModifiers: []planmodifier.List{
					ReplaceWhenStartTrackingItems(),
				},
			},
		},
	}
}

func (r *ListResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Create a new resource.
func (r *ListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the list resource")
	// Retrieve values from plan
	var plan listResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	kind, err := openapiclient.NewListKindFromValue(plan.Kind.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create the list",
			err.Error(),
		)
		return
	}

	postBody := openapiclient.ListPostBody{
		Name: plan.Name.ValueString(),
		Kind: *kind,
	}

	if len(plan.Items) > 0 {
		postBody.Items = []openapiclient.ListItemsInner{}
		for _, item := range plan.Items {
			listItem := openapiclient.NewListItemsInner(item.Value.ValueString())
			if !item.Expiration.IsNull() {
				expiration, err := time.Parse(time.RFC3339, item.Expiration.ValueString())
				if err != nil {
					resp.Diagnostics.AddError(
						"Invalid Expiration Date",
						fmt.Sprintf("The expiration date '%s' is not a valid RFC 3339 date: %s", item.Expiration, err),
					)
					return
				}
				listItem.SetExpiration(expiration)
			}

			postBody.Items = append(postBody.Items, *listItem)
		}
	}

	if !plan.Subkind.IsNull() {
		subkind, err := openapiclient.NewListSubkindFromValue(plan.Subkind.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to create the list",
				err.Error(),
			)
			return
		}
		postBody.Subkind = subkind
	}

	listResponse, _, err := r.client.ListsAPI.CreateList(ctx, r.client.OrgID).ListPostBody(postBody).Execute()

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the list",
			message,
		)
		return
	}

	// Map response body to model
	plan.ID = types.StringValue(listResponse.Id)
	plan.Name = types.StringValue(listResponse.Name)
	plan.Kind = types.StringValue(string(listResponse.Kind))
	if listResponse.Subkind != nil {
		plan.Subkind = types.StringValue(string(*listResponse.Subkind))
	}

	if plan.Items != nil && len(listResponse.Items) > 0 {
		applyResponseToState(listResponse, &plan)
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Created the list resource", map[string]any{"success": true})
}

// Read resource information.
func (r *ListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the list resource")
	// Get current state
	var state listResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	listResponse, httpResp, err := r.client.ListsAPI.GetList(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
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
			"Unable to read the list",
			message,
		)
		return
	}

	// Map response body to model
	state.ID = types.StringValue(listResponse.Id)
	state.Name = types.StringValue(listResponse.Name)
	state.Kind = types.StringValue(string(listResponse.Kind))
	if listResponse.Subkind != nil {
		state.Subkind = types.StringValue(string(*listResponse.Subkind))
	}

	// Because we cannot pull config to check here
	// ReplaceWhenStartTrackingItems plan modifier is used to relacea list resource when items goes from null to set
	if state.Items != nil {
		applyResponseToState(listResponse, &state)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Finished reading the list resource", map[string]any{"success": true})
}

func (r *ListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the list resource")
	// Retrieve values from plan
	var plan listResourceModel
	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state listResourceModel
	diags = req.State.Get(ctx, &state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	patchBody := openapiclient.ListPatchBody{}

	if !plan.Name.IsNull() {
		name := plan.Name.ValueString()
		patchBody.Name = &name
	}

	diffLists(state.Items, plan.Items, &patchBody, resp)
	if resp.Diagnostics.HasError() {
		return
	}

	listRequest := r.client.ListsAPI.UpdateList(ctx, r.client.OrgID, plan.ID.ValueString()).
		ListPatchBody(patchBody)

	// update a list
	listResponse, _, err := listRequest.Execute()

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the list",
			message,
		)
		return
	}

	// Overwrite the list with refreshed state
	newState := listResourceModel{
		ID:    types.StringValue(listResponse.Id),
		Name:  types.StringValue(listResponse.Name),
		Kind:  types.StringValue(string(listResponse.Kind)),
		Items: plan.Items,
	}

	if listResponse.Subkind != nil {
		newState.Subkind = types.StringValue(string(*listResponse.Subkind))
	}

	if plan.Items != nil {
		applyResponseToState(listResponse, &newState)
	}

	// Set the refreshed state
	diags = resp.State.Set(ctx, newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updated the list resource", map[string]any{"success": true})
}

func (r *ListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the list resource")
	// Retrieve values from a state
	var state listResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete a list
	_, err := r.client.ListsAPI.DeleteList(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the list",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the list resource", map[string]any{"success": true})
}

func (r *ListResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var plan listResourceModel
	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !plan.Kind.IsNull() {
		_, err := openapiclient.NewListKindFromValue(plan.Kind.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Configuration Error: Ivalid value",
				err.Error(),
			)
		}
	}

	if !plan.Subkind.IsNull() {
		_, err := openapiclient.NewListSubkindFromValue(plan.Subkind.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"List Configuration Error: Ivalid value",
				err.Error(),
			)
		}

		if !strings.HasSuffix(plan.Subkind.ValueString(), fmt.Sprintf("_%s", plan.Kind.ValueString())) {
			resp.Diagnostics.AddError(
				"List Configuration Error: Invalid value",
				fmt.Sprintf("Subkind %s is not allowed with kind %s.", plan.Subkind.ValueString(), plan.Kind.ValueString()),
			)
		}
	}
}

func diffLists(oldList, newList []listItemModel, patchBody *openapiclient.ListPatchBody, resp *resource.UpdateResponse) {
	oldMap := make(map[string]basetypes.StringValue)
	newMap := make(map[string]basetypes.StringValue)

	// Create a map from old list with value as key and expiration as value
	for _, item := range oldList {
		oldMap[item.Value.ValueString()] = item.Expiration
	}

	// Create a map from new list with value as key and expiration as value
	for _, item := range newList {
		newMap[item.Value.ValueString()] = item.Expiration
	}

	var toAdd []openapiclient.ListItemsInner
	var toRemove []string

	// Find items to remove
	for _, item := range oldList {
		if _, exists := newMap[item.Value.ValueString()]; !exists {
			toRemove = append(toRemove, item.Value.ValueString())
		}
	}

	// Find items to add
	for _, item := range newList {
		if oldExp, exists := oldMap[item.Value.ValueString()]; !exists || !compareStringValues(oldExp, item.Expiration) {
			listItem := openapiclient.NewListItemsInner(item.Value.ValueString())
			if !item.Expiration.IsNull() {
				expiration, err := time.Parse(time.RFC3339, item.Expiration.ValueString())
				if err != nil {
					resp.Diagnostics.AddError(
						"Invalid Expiration Date",
						fmt.Sprintf("The expiration date '%s' is not a valid RFC 3339 date: %s", item.Expiration, err),
					)
					return
				}
				listItem.SetExpiration(expiration)
			}
			toAdd = append(toAdd, *listItem)
		}
	}

	patchBody.Add = toAdd
	patchBody.Remove = toRemove
}

func compareStringValues(a, b types.String) bool {
	// If both are null, they are equal
	if a.IsNull() && b.IsNull() {
		return true
	}
	// If one is null and the other is not, they are not equal
	if a.IsNull() || b.IsNull() {
		return false
	}
	// Compare actual values
	return a.ValueString() == b.ValueString()
}

func applyResponseToState(listResponse *openapiclient.List, state *listResourceModel) {
	responseItemsMap := make(map[string]openapiclient.ListItemsInner)
	for _, item := range listResponse.Items {
		responseItemsMap[item.Value] = item
	}

	valueElements := make([]listItemModel, len(listResponse.Items))
	pos := 0
	for _, item := range state.Items {
		normalized := getListItemRepresentation(state.Kind.ValueString(), item.Value.ValueString())
		val, ok := responseItemsMap[normalized]

		if ok {
			valueElements[pos] = item
			delete(responseItemsMap, val.Value)
			if !val.GetExpiration().IsZero() {
				valueElements[pos].Expiration = types.StringValue(val.GetExpiration().Format(time.RFC3339))
			}
			pos++
		}
	}

	// Append new items
	for _, v := range responseItemsMap {
		valueElements[pos] = listItemModel{
			Value: types.StringValue(v.Value),
		}
		if !v.GetExpiration().IsZero() {
			valueElements[pos].Expiration = types.StringValue(v.GetExpiration().Format(time.RFC3339))
		}
		pos++
	}

	state.Items = valueElements
}

func getListItemRepresentation(kind string, item string) string {
	if kind == "ip" {
		ipRange, ok := parsePrefixRangeOrAddr(item)
		if ok {
			return ipRange.String()
		}
	}

	return item
}

func parsePrefixRangeOrAddr(s string) (ipRange netipx.IPRange, ok bool) {
	switch {
	case strings.IndexByte(s, '-') > 0:
		var err error
		ipRange, err = netipx.ParseIPRange(s)
		if err != nil || !ipRange.IsValid() {
			return ipRange, false
		}

		return ipRange, true
	case strings.LastIndexByte(s, '/') > 0:
		prefix, err := netip.ParsePrefix(s)
		if err != nil || !prefix.IsValid() {
			return ipRange, false
		}

		return netipx.RangeOfPrefix(prefix), true
	default:
		addr, err := netip.ParseAddr(s)
		if err != nil || !addr.IsValid() {
			return ipRange, false
		}

		return netipx.IPRangeFrom(addr, addr), true
	}
}
