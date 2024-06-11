package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type dateTimeValidator struct {
	name string
}

func (v dateTimeValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("Ensure the '%s' attribute is a valid datetime string and not in the past.", v.name)
}

func (v dateTimeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v dateTimeValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()
	parsedTime, err := time.Parse(time.RFC3339, value)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid DateTime Format",
			fmt.Sprintf("The '%s' attribute '%s' is not a valid RFC 3339 datetime string.\nSource: %s", v.name, value, req.Path.String()),
		)
		return
	}

	if parsedTime.Before(time.Now()) {
		resp.Diagnostics.AddError(
			"DateTime In The Past",
			fmt.Sprintf("The '%s' attribute '%s' is in the past. It must be a future datetime.\nSource: %s", v.name, value, req.Path.String()),
		)
	}
}

func dateTimeNotPast(name string) validator.String {
	return dateTimeValidator{name: name}
}
