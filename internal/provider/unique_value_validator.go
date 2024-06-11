package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type uniqueValueValidator struct{}

func (v uniqueValueValidator) Description(ctx context.Context) string {
	return "Ensure that all 'value' attributes in the list are unique."
}

func (v uniqueValueValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v uniqueValueValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	var list []basetypes.ObjectValue
	diags := req.ConfigValue.ElementsAs(ctx, &list, false)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	valueSet := make(map[string]struct{})
	for _, item := range list {
		valueAttr := item.Attributes()["value"]
		if valueAttr.IsNull() || valueAttr.IsUnknown() {
			continue
		}

		valueString, ok := valueAttr.(types.String)

		if !ok {
			resp.Diagnostics.AddError(
				"Unexpected Type",
				fmt.Sprintf("Unxpected 'value' type '%T'. This is likely a bug in terraform provider.\nSource: %s", valueString, req.Path.String()),
			)
			return
		}

		if _, exists := valueSet[valueString.ValueString()]; exists {
			resp.Diagnostics.AddError(
				"Duplicate 'value' Attribute",
				fmt.Sprintf("The 'value' attribute '%s' is not unique.\nSource: %s", valueString.ValueString(), req.Path.String()),
			)
			return
		}
		valueSet[valueString.ValueString()] = struct{}{}
	}
}

func uniqueValue() validator.List {
	return uniqueValueValidator{}
}
