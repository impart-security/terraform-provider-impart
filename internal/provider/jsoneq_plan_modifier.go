package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type conditionalEqualJSONPlanModifier struct {
	AttrPath  path.Path // The attribute path this modifier is applied to
	AttrValue string    // Only apply if the attribute value is equal to this string
	Subset    bool      // If true, check if the subset of the JSON is equal to a value returned by api
}

var _ planmodifier.String = conditionalEqualJSONPlanModifier{}

func (m conditionalEqualJSONPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.PlanValue.IsNull() || req.StateValue.IsNull() {
		return
	}

	// Skip match if attribut path is not set
	if m.AttrPath.String() != "" {
		// Get the value field from config
		var value types.String
		diags := req.Config.GetAttribute(ctx, m.AttrPath, &value)
		resp.Diagnostics.Append(diags...)
		if diags.HasError() || value.IsNull() || value.IsUnknown() {
			return
		}

		if value.ValueString() != m.AttrValue {
			return
		}
	}

	// Compare the JSON values structurally
	planStr := req.PlanValue.ValueString()
	stateStr := req.StateValue.ValueString()

	if m.Subset {
		if subsetEqualJSON(planStr, stateStr) {
			resp.RequiresReplace = false
			resp.PlanValue = req.StateValue
		}
		return
	}

	if equalJSON(planStr, stateStr) {
		resp.RequiresReplace = false
		resp.PlanValue = req.StateValue
	}
}

func (m conditionalEqualJSONPlanModifier) Description(_ context.Context) string {
	return "Skips plan diff if content is JSON and structurally equal to prior state"
}

func (m conditionalEqualJSONPlanModifier) MarkdownDescription(_ context.Context) string {
	return m.Description(context.Background())
}

// Required for terraform-plugin-framework schema.PlanModifier
func (m conditionalEqualJSONPlanModifier) ModifyPlan(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	m.PlanModifyString(ctx, req, resp)
}
