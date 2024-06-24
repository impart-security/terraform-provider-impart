package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// ReplaceWhenStartTrackingItems replaces the planned value with the state value if the
// user sets items which previously were set as null.
// This means that items are will be tracked in the state.
func ReplaceWhenStartTrackingItems() planmodifier.List {
	return replaceWhenStartTrackingItems{}
}

// replaceWhenStartTrackingItems implements the plan modifier.
type replaceWhenStartTrackingItems struct{}

// Description returns a human-readable description of the plan modifier.
func (m replaceWhenStartTrackingItems) Description(_ context.Context) string {
	return "Whem items goes from null to set a list resource will be replaced."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m replaceWhenStartTrackingItems) MarkdownDescription(_ context.Context) string {
	return "Whem items goes from null to set a list resource will be replaced."
}

// PlanModifyString implements the plan modification logic.
func (m replaceWhenStartTrackingItems) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	// Do nothing if there is an unknown configuration value, otherwise interpolation gets messed up.
	if req.ConfigValue.IsUnknown() {
		return
	}

	if !req.ConfigValue.IsNull() && req.StateValue.IsNull() {
		resp.RequiresReplace = true
	}
}
