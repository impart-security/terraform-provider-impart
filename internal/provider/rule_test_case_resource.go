package provider

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	openapiclient "github.com/impart-security/terraform-provider-impart/internal/apiclient"
)

var (
	_ resource.Resource                = &ruleTestcaseResource{}
	_ resource.ResourceWithConfigure   = &ruleTestcaseResource{}
	_ resource.ResourceWithImportState = &ruleTestcaseResource{}
)

type AssertionType string

const (
	AssertionTypeOutput     AssertionType = "output"
	AssertionTypeBlock      AssertionType = "block"
	AssertionTypeStatusCode AssertionType = "status_code"
	AssertionTypeTags       AssertionType = "tags"
)

// NewRuleTestcaseResource is a helper function to simplify the provider implementation.
func NewRuleTestcaseResource() resource.Resource {
	return &ruleTestcaseResource{}
}

// ruleTestcaseResource is the resource implementation.
type ruleTestcaseResource struct {
	client *impartAPIClient
}

// Configure adds the provider configured client to the resource.
func (r *ruleTestcaseResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *ruleTestcaseResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rule_test_case"
}

type ruleTestCaseModel struct {
	ID          types.String     `tfsdk:"id"`
	Name        types.String     `tfsdk:"name"`
	Description types.String     `tfsdk:"description"`
	Required    types.Bool       `tfsdk:"required"`
	Messages    []messageModel   `tfsdk:"messages"`
	Assertions  []assertionModel `tfsdk:"assertions"`
}

type messageModel struct {
	Req       reqModel    `tfsdk:"req"`
	Res       resModel    `tfsdk:"res"`
	Count     types.Int32 `tfsdk:"count"`
	Delay     types.Int32 `tfsdk:"delay"`
	PostDelay types.Int32 `tfsdk:"post_delay"`
}

type reqModel struct {
	URL           types.String   `tfsdk:"url"`
	Method        types.String   `tfsdk:"method"`
	TruncatedBody types.Bool     `tfsdk:"truncated_body"`
	Body          types.String   `tfsdk:"body"`
	HeaderKeys    []types.String `tfsdk:"header_keys"`
	HeaderValues  []types.String `tfsdk:"header_values"`
	CookieKeys    []types.String `tfsdk:"cookie_keys"`
	CookieValues  []types.String `tfsdk:"cookie_values"`
	RemoteAddr    types.String   `tfsdk:"remote_addr"`
}

type resModel struct {
	TruncatedBody types.Bool     `tfsdk:"truncated_body"`
	Body          types.String   `tfsdk:"body"`
	HeaderKeys    []types.String `tfsdk:"header_keys"`
	HeaderValues  []types.String `tfsdk:"header_values"`
	StatusCode    types.Int32    `tfsdk:"status_code"`
}

type assertionModel struct {
	MessageIndexes []types.Int32 `tfsdk:"message_indexes"`
	AssertionType  types.String  `tfsdk:"assertion_type"`
	Location       types.String  `tfsdk:"location"`
	Condition      types.String  `tfsdk:"condition"`
	Expected       types.String  `tfsdk:"expected"`
}

func (r ruleTestcaseResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier of the test case.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the test case.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the test case.",
				Optional:    true,
			},
			"required": schema.BoolAttribute{
				Description: "Sets if test case required to pass on update.",
				Optional:    true,
			},
			"messages": schema.ListNestedAttribute{
				Description: "The messages of the test case.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"req": schema.SingleNestedAttribute{
							Description: "A payload sent to the inspector to inspect an HTTP request.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"url": schema.StringAttribute{
									Description: "The URL of the request.",
									Required:    true,
								},
								"method": schema.StringAttribute{
									Description: "The method of the request.",
									Required:    true,
								},
								"truncated_body": schema.BoolAttribute{
									Description: "Indicates whether the request body was truncated.",
									Optional:    true,
								},
								"body": schema.StringAttribute{
									Description: "The base64 encoded HTTP request body.",
									Optional:    true,
								},
								"header_keys": schema.ListAttribute{
									Description: "The HTTP request header keys.",
									Optional:    true,
									ElementType: types.StringType,
									Validators: []validator.List{
										listvalidator.SizeAtLeast(1),
									},
								},
								"header_values": schema.ListAttribute{
									Description: "The HTTP request header values.",
									Optional:    true,
									ElementType: types.StringType,
									Validators: []validator.List{
										listvalidator.SizeAtLeast(1),
									},
								},
								"cookie_keys": schema.ListAttribute{
									Description: "The HTTP request cookie keys.",
									Optional:    true,
									ElementType: types.StringType,
									Validators: []validator.List{
										listvalidator.SizeAtLeast(1),
									},
								},
								"cookie_values": schema.ListAttribute{
									Description: "The HTTP request cookie values.",
									Optional:    true,
									ElementType: types.StringType,
									Validators: []validator.List{
										listvalidator.SizeAtLeast(1),
									},
								},
								"remote_addr": schema.StringAttribute{
									Description: "The remote address of the request.",
									Optional:    true,
								},
							},
						},
						"res": schema.SingleNestedAttribute{
							Required:    true,
							Description: "A payload sent to the inspector to inspect an HTTP response.",
							Attributes: map[string]schema.Attribute{
								"truncated_body": schema.BoolAttribute{
									Description: "Indicates whether the response body was truncated.",
									Optional:    true,
								},
								"body": schema.StringAttribute{
									Description: "The base64 encoded HTTP response body.",
									Optional:    true,
								},
								"header_keys": schema.ListAttribute{
									Description: "The HTTP response header keys.",
									Optional:    true,
									ElementType: types.StringType,
									Validators: []validator.List{
										listvalidator.SizeAtLeast(1),
									},
								},
								"header_values": schema.ListAttribute{
									Description: "The HTTP response header values.",
									Optional:    true,
									ElementType: types.StringType,
									Validators: []validator.List{
										listvalidator.SizeAtLeast(1),
									},
								},
								"status_code": schema.Int32Attribute{
									Description: "The HTTP response status code.",
									Required:    true,
								},
							},
						},
						"count": schema.Int32Attribute{
							Description: "The number of times to include the message in the test case.",
							Default:     int32default.StaticInt32(1),
							Computed:    true,
							Optional:    true,
							Validators: []validator.Int32{
								int32validator.AtLeast(1),
							},
						},
						"delay": schema.Int32Attribute{
							Description: "The delay in milliseconds between message iterations.",
							Default:     int32default.StaticInt32(0),
							Computed:    true,
							Optional:    true,
							Validators: []validator.Int32{
								int32validator.AtLeast(0),
							},
						},
						"post_delay": schema.Int32Attribute{
							Description: "The delay in milliseconds after a set of message iterations.",
							Default:     int32default.StaticInt32(0),
							Computed:    true,
							Optional:    true,
							Validators: []validator.Int32{
								int32validator.AtLeast(0),
							},
						},
					},
				},
			},
			"assertions": schema.ListNestedAttribute{
				Description: "The assertions of the test case.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"message_indexes": schema.ListAttribute{
							Description: "The indexes of the messages in the test case the assertion applies to.",
							ElementType: types.Int32Type,
							Required:    true,
							Validators: []validator.List{
								listvalidator.UniqueValues(),
							},
						},
						"assertion_type": schema.StringAttribute{
							Description: "The assertion type of the request.",
							Required:    true,
							Validators: []validator.String{stringvalidator.OneOf(string(AssertionTypeOutput),
								string(AssertionTypeStatusCode),
								string(AssertionTypeTags),
								string(AssertionTypeBlock))},
						},
						"location": schema.StringAttribute{
							Description: "The location of the assertion. Allowed values: req, res. Not applicable for assertion type output.",
							Optional:    true,
							Validators:  []validator.String{stringvalidator.OneOf("req", "res")},
						},
						"condition": schema.StringAttribute{
							Description: "The condition of the request.",
							MarkdownDescription: `
The condition of the assertion.
Accepted values per assertion type:
**output**: contains, not_contains.
**tags**: contains, not_contains.
**status_code**: equal, not_equal, greater_than, less_than, one_of.
**block**: N/A`,
							Optional: true,
						},
						"expected": schema.StringAttribute{
							Description: "The expected value of the assertion.",
							MarkdownDescription: `
The expected value of the assertion.
It is a string value, and the format it must satisfy depends on the assertion type:
**output**: A string.
**tags**: A string.
**status_code**: An integer value (e.g., "200") or comma-separated list of integers for one_of condition (e.g., "200,404,500").
**block**: A boolean value represented as "true" or "false".`,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func (r *ruleTestcaseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r ruleTestcaseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the test case resource")

	var plan ruleTestCaseModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody, err := toRulesTestCasePostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create the rule test case",
			err.Error(),
		)
		return
	}

	ruleTestCaseRequest := r.client.RulesTestCasesAPI.CreateRulesTestCase(ctx, r.client.OrgID).
		RulesTestCasePostBody(postBody)

	ruleTestCaseResponse, _, err := ruleTestCaseRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the rule test case",
			message,
		)
		return
	}

	state := toRuleTestCaseModel(ruleTestCaseResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Created the test case resource", map[string]any{"success": true})
}

func (r ruleTestcaseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the test case resource")

	var data ruleTestCaseModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ruleTestCaseResponse, httpResp, err := r.client.RulesTestCasesAPI.GetRulesTestCase(ctx, r.client.OrgID, data.ID.ValueString()).
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
			"Unable to read the test case",
			message,
		)
		return
	}

	state := toRuleTestCaseModel(ruleTestCaseResponse, data)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Finished reading the test case resource", map[string]any{"success": true})
}

func (r ruleTestcaseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the test case resource")

	var plan ruleTestCaseModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postBody, err := toRulesTestCasePostBody(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update the rule test case",
			err.Error(),
		)
		return
	}

	ruleTestCaseRequest := r.client.RulesTestCasesAPI.UpdateRulesTestCase(ctx, r.client.OrgID, plan.ID.ValueString()).
		RulesTestCasePostBody(postBody)

	ruleTestCaseResponse, _, err := ruleTestCaseRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the rule test case",
			message,
		)
		return
	}

	state := toRuleTestCaseModel(ruleTestCaseResponse, plan)

	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updated the test case resource", map[string]any{"success": true})
}

func (r ruleTestcaseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the test case resource")
	// Retrieve values from a state
	var state ruleTestCaseModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete a test case
	_, err := r.client.RulesTestCasesAPI.DeleteRulesTestCase(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the test case",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the test case resource", map[string]any{"success": true})
}

func fromArray(strings []string) []types.String {
	var result []types.String
	for _, s := range strings {
		result = append(result, types.StringValue(s))
	}
	return result
}

func fromInt32Array(items []int32) []types.Int32 {
	var result []types.Int32
	for _, item := range items {
		result = append(result, types.Int32Value(item))
	}
	return result
}

func toInt32Array(items []types.Int32) []int32 {
	var result []int32
	for _, item := range items {
		result = append(result, item.ValueInt32())
	}
	return result
}

func toRulesTestCasePostBody(plan ruleTestCaseModel) (openapiclient.RulesTestCasePostBody, error) {
	postBody := openapiclient.RulesTestCasePostBody{
		Name:     plan.Name.ValueString(),
		Messages: make([]openapiclient.RulesTestCaseMessagesInner, 0, len(plan.Messages)),
	}

	if !plan.Description.IsNull() {
		description := plan.Description.ValueString()
		postBody.Description = &description
	}

	if !plan.Required.IsNull() {
		required := plan.Required.ValueBool()
		postBody.Required = &required
	}

	if len(plan.Messages) == 0 {
		return postBody, nil
	}

	for _, message := range plan.Messages {
		rtMessage := openapiclient.RulesTestCaseMessagesInner{
			Req: openapiclient.InspectorReqMsg{
				Url:    message.Req.URL.ValueString(),
				Method: message.Req.Method.ValueString(),
			},
			Res: openapiclient.InspectorResMsg{
				StatusCode: message.Res.StatusCode.ValueInt32(),
			},
			Count:     message.Count.ValueInt32(),
			Delay:     message.Delay.ValueInt32(),
			PostDelay: message.PostDelay.ValueInt32(),
		}

		if !message.Req.TruncatedBody.IsNull() {
			truncated := message.Req.TruncatedBody.ValueBool()
			rtMessage.Req.TruncatedBody = &truncated
		}

		if !message.Req.Body.IsNull() {
			body := message.Req.Body.ValueString()
			rtMessage.Req.Body = &body
		}

		if !message.Req.RemoteAddr.IsNull() {
			remoteAddr := message.Req.RemoteAddr.ValueString()
			rtMessage.Req.RemoteAddr = &remoteAddr
		}

		if len(message.Req.HeaderKeys) > 0 {
			rtMessage.Req.HeaderKeys = make([]string, 0, len(message.Req.HeaderKeys))
			for _, key := range message.Req.HeaderKeys {
				rtMessage.Req.HeaderKeys = append(rtMessage.Req.HeaderKeys, key.ValueString())
			}
		}

		if len(message.Req.HeaderValues) > 0 {
			rtMessage.Req.HeaderValues = make([]string, 0, len(message.Req.HeaderValues))
			for _, value := range message.Req.HeaderValues {
				rtMessage.Req.HeaderValues = append(rtMessage.Req.HeaderValues, value.ValueString())
			}
		}

		if len(message.Req.CookieKeys) > 0 {
			rtMessage.Req.CookieKeys = make([]string, 0, len(message.Req.CookieKeys))
			for _, key := range message.Req.CookieKeys {
				rtMessage.Req.CookieKeys = append(rtMessage.Req.CookieKeys, key.ValueString())
			}
		}

		if len(message.Req.CookieValues) > 0 {
			rtMessage.Req.CookieValues = make([]string, 0, len(message.Req.CookieValues))
			for _, value := range message.Req.CookieValues {
				rtMessage.Req.CookieValues = append(rtMessage.Req.CookieValues, value.ValueString())
			}
		}

		if !message.Res.Body.IsNull() {
			body := message.Res.Body.ValueString()
			rtMessage.Res.Body = &body
		}

		if !message.Res.TruncatedBody.IsNull() {
			truncated := message.Res.TruncatedBody.ValueBool()
			rtMessage.Res.TruncatedBody = &truncated
		}

		if len(message.Res.HeaderKeys) > 0 {
			rtMessage.Res.HeaderKeys = make([]string, 0, len(message.Res.HeaderKeys))
			for _, key := range message.Res.HeaderKeys {
				rtMessage.Res.HeaderKeys = append(rtMessage.Res.HeaderKeys, key.ValueString())
			}
		}

		if len(message.Res.HeaderValues) > 0 {
			rtMessage.Res.HeaderValues = make([]string, 0, len(message.Res.HeaderValues))
			for _, value := range message.Res.HeaderValues {
				rtMessage.Res.HeaderValues = append(rtMessage.Res.HeaderValues, value.ValueString())
			}
		}
		postBody.Messages = append(postBody.Messages, rtMessage)
	}

	for _, assertion := range plan.Assertions {
		var a openapiclient.RulesTestCaseAssertion
		if assertion.AssertionType.ValueString() == string(AssertionTypeBlock) {
			a = openapiclient.RulesTestCaseAssertion{
				RulesTestCaseAssertionBlock: &openapiclient.RulesTestCaseAssertionBlock{
					MessageIndexes: toInt32Array(assertion.MessageIndexes),
					AssertionType:  assertion.AssertionType.ValueString(),
					Location:       assertion.Location.ValueString(),
					Expected:       assertion.Expected.ValueString() == "true",
				},
			}
		} else if assertion.AssertionType.ValueString() == string(AssertionTypeOutput) {
			condition, err := openapiclient.NewRulesTestCaseAssertionConditionPresenceFromValue(assertion.Condition.ValueString())
			if err != nil {
				return postBody, fmt.Errorf("unable to create assertion condition: %w", err)
			}

			a = openapiclient.RulesTestCaseAssertion{
				RulesTestCaseAssertionOutput: &openapiclient.RulesTestCaseAssertionOutput{
					MessageIndexes: toInt32Array(assertion.MessageIndexes),
					AssertionType:  assertion.AssertionType.ValueString(),
					Condition:      *condition,
					Expected:       assertion.Expected.ValueString(),
				},
			}
		} else if assertion.AssertionType.ValueString() == string(AssertionTypeStatusCode) {
			stringValues := strings.Split(assertion.Expected.ValueString(), ",")
			intValues := make([]int32, len(stringValues))
			for i, v := range stringValues {
				value, err := strconv.Atoi(v)
				if err != nil { //this should not happened as we are validating the resource
					return postBody, fmt.Errorf("unable to convert string to int: %w", err)
				}
				intValues[i] = int32(value)
			}

			condition, err := openapiclient.NewRulesTestCaseAssertionConditionDefaultFromValue(assertion.Condition.ValueString())
			if err != nil {
				return postBody, fmt.Errorf("unable to create assertion condition: %w", err)
			}

			a = openapiclient.RulesTestCaseAssertion{
				RulesTestCaseAssertionStatusCode: &openapiclient.RulesTestCaseAssertionStatusCode{
					MessageIndexes: toInt32Array(assertion.MessageIndexes),
					AssertionType:  assertion.AssertionType.ValueString(),
					Location:       assertion.Location.ValueString(),
					Condition:      *condition,
					Expected:       intValues,
				},
			}
		} else if assertion.AssertionType.ValueString() == string(AssertionTypeTags) {
			condition, err := openapiclient.NewRulesTestCaseAssertionConditionPresenceFromValue(assertion.Condition.ValueString())
			if err != nil {
				return postBody, fmt.Errorf("unable to create assertion condition: %w", err)
			}

			a = openapiclient.RulesTestCaseAssertion{
				RulesTestCaseAssertionTags: &openapiclient.RulesTestCaseAssertionTags{
					MessageIndexes: toInt32Array(assertion.MessageIndexes),
					AssertionType:  assertion.AssertionType.ValueString(),
					Location:       assertion.Location.ValueString(),
					Condition:      *condition,
					Expected:       assertion.Expected.ValueString(),
				},
			}
		}

		postBody.Assertions = append(postBody.Assertions, a)
	}

	return postBody, nil
}

func toRuleTestCaseModel(ruleTestCaseResponse *openapiclient.RulesTestCase, plan ruleTestCaseModel) ruleTestCaseModel {
	testCaseModel := ruleTestCaseModel{
		ID:       types.StringValue(ruleTestCaseResponse.Id),
		Name:     types.StringValue(ruleTestCaseResponse.Name),
		Messages: make([]messageModel, 0, len(ruleTestCaseResponse.Messages)),
	}

	if !plan.Required.IsNull() || ruleTestCaseResponse.Required {
		testCaseModel.Required = plan.Required
	}

	if !plan.Description.IsNull() || ruleTestCaseResponse.Description != "" {
		testCaseModel.Description = types.StringValue(ruleTestCaseResponse.Description)
	}

	for i, message := range ruleTestCaseResponse.Messages {
		messageModel := messageModel{
			Count:     types.Int32Value(message.Count),
			Delay:     types.Int32Value(message.Delay),
			PostDelay: types.Int32Value(message.PostDelay),
			Req: reqModel{
				URL:          types.StringValue(message.Req.Url),
				Method:       types.StringValue(message.Req.Method),
				HeaderKeys:   fromArray(message.Req.HeaderKeys),
				HeaderValues: fromArray(message.Req.HeaderValues),
				CookieKeys:   fromArray(message.Req.CookieKeys),
				CookieValues: fromArray(message.Req.CookieValues),
			},
			Res: resModel{
				StatusCode:   types.Int32Value(message.Res.StatusCode),
				HeaderKeys:   fromArray(message.Res.HeaderKeys),
				HeaderValues: fromArray(message.Res.HeaderValues),
			},
		}

		if message.Req.Body != nil {
			messageModel.Req.Body = types.StringValue(*message.Req.Body)
		}

		// api doesn't return false so we need to set if plan has it
		if message.Req.TruncatedBody != nil {
			messageModel.Req.TruncatedBody = types.BoolValue(*message.Req.TruncatedBody)
		} else if i < len(plan.Messages) && !plan.Messages[i].Req.TruncatedBody.IsNull() {
			messageModel.Req.TruncatedBody = plan.Messages[i].Req.TruncatedBody
		}

		if message.Req.RemoteAddr != nil {
			messageModel.Req.RemoteAddr = types.StringValue(*message.Req.RemoteAddr)
		}

		if message.Res.Body != nil {
			messageModel.Res.Body = types.StringValue(*message.Res.Body)
		}

		// api doesn't return false so we need to set if plan has it
		if message.Res.TruncatedBody != nil {
			messageModel.Res.TruncatedBody = types.BoolValue(*message.Res.TruncatedBody)
		} else if i < len(plan.Messages) && !plan.Messages[i].Res.TruncatedBody.IsNull() {
			messageModel.Res.TruncatedBody = plan.Messages[i].Res.TruncatedBody
		}

		testCaseModel.Messages = append(testCaseModel.Messages, messageModel)
	}

	// Handle empty array in state
	if len(ruleTestCaseResponse.Assertions) == 0 && plan.Assertions != nil {
		testCaseModel.Assertions = make([]assertionModel, 0)
	}

	for _, assertion := range ruleTestCaseResponse.Assertions {
		var aModel assertionModel
		if assertion.RulesTestCaseAssertionBlock != nil {
			aModel = assertionModel{
				MessageIndexes: fromInt32Array(assertion.RulesTestCaseAssertionBlock.MessageIndexes),
				AssertionType:  types.StringValue(assertion.RulesTestCaseAssertionBlock.AssertionType),
				Location:       types.StringValue(assertion.RulesTestCaseAssertionBlock.Location),
				Expected:       types.StringValue(fmt.Sprintf("%t", assertion.RulesTestCaseAssertionBlock.Expected)),
			}
		} else if assertion.RulesTestCaseAssertionOutput != nil {
			aModel = assertionModel{
				MessageIndexes: fromInt32Array(assertion.RulesTestCaseAssertionOutput.MessageIndexes),
				AssertionType:  types.StringValue(assertion.RulesTestCaseAssertionOutput.AssertionType),
				Condition:      types.StringValue(string(assertion.RulesTestCaseAssertionOutput.Condition)),
				Expected:       types.StringValue(assertion.RulesTestCaseAssertionOutput.Expected),
			}
		} else if assertion.RulesTestCaseAssertionStatusCode != nil {
			aModel = assertionModel{
				MessageIndexes: fromInt32Array(assertion.RulesTestCaseAssertionStatusCode.MessageIndexes),
				AssertionType:  types.StringValue(assertion.RulesTestCaseAssertionStatusCode.AssertionType),
				Location:       types.StringValue(assertion.RulesTestCaseAssertionStatusCode.Location),
				Condition:      types.StringValue(string(assertion.RulesTestCaseAssertionStatusCode.Condition)),
				Expected:       types.StringValue(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(assertion.RulesTestCaseAssertionStatusCode.Expected)), ","), "[]")), // convert []int32 to comma separated string
			}
		} else if assertion.RulesTestCaseAssertionTags != nil {
			aModel = assertionModel{
				MessageIndexes: fromInt32Array(assertion.RulesTestCaseAssertionTags.MessageIndexes),
				AssertionType:  types.StringValue(assertion.RulesTestCaseAssertionTags.AssertionType),
				Location:       types.StringValue(assertion.RulesTestCaseAssertionTags.Location),
				Condition:      types.StringValue(string(assertion.RulesTestCaseAssertionTags.Condition)),
				Expected:       types.StringValue(assertion.RulesTestCaseAssertionTags.Expected),
			}
		}

		testCaseModel.Assertions = append(testCaseModel.Assertions, aModel)
	}

	return testCaseModel
}

func (r *ruleTestcaseResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var plan ruleTestCaseModel
	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	conditions := map[string][]string{
		string(AssertionTypeOutput):     {"contains", "not_contains"},
		string(AssertionTypeStatusCode): {"equal", "not_equal", "greater_than", "less_than", "one_of"},
		string(AssertionTypeTags):       {"contains", "not_contains"},
		string(AssertionTypeBlock):      nil,
	}

	for _, assertion := range plan.Assertions {
		conditions, ok := conditions[assertion.AssertionType.ValueString()]
		if !ok {
			continue
		}

		if len(conditions) == 0 {
			if !assertion.Condition.IsNull() {
				resp.Diagnostics.AddError(
					"Invalid Assertion Condition",
					fmt.Sprintf("The condition attribute is not appplicable for assertion type %q", assertion.AssertionType.ValueString()),
				)
			}
		} else {
			if assertion.Condition.IsNull() {
				resp.Diagnostics.AddError(
					"Missing Assertion Condition",
					fmt.Sprintf("The condition attribute is required for assertion type %q", assertion.AssertionType.ValueString()),
				)
			} else if !slices.Contains(conditions, assertion.Condition.ValueString()) {
				resp.Diagnostics.AddError(
					"Invalid Assertion Condition",
					fmt.Sprintf("The condition attribute must be one of %q for assertion type %q", strings.Join(conditions, ", "), assertion.AssertionType.ValueString()),
				)
			}
		}

		// Validate expected value
		if assertion.AssertionType.ValueString() == string(AssertionTypeStatusCode) {
			stringValues := strings.Split(assertion.Expected.ValueString(), ",")
			for _, v := range stringValues {
				if _, err := strconv.Atoi(v); err != nil {
					resp.Diagnostics.AddError(
						"Invalid Assertion Expected Value",
						fmt.Sprintf("The expected value must be a comma-separated list of integers for assertion type %q", assertion.AssertionType.ValueString()),
					)
					break
				}
			}
		} else if assertion.AssertionType.ValueString() == string(AssertionTypeBlock) {
			if assertion.Expected.ValueString() != "true" && assertion.Expected.ValueString() != "false" {
				resp.Diagnostics.AddError(
					"Invalid Assertion Expected Value",
					fmt.Sprintf("The expected value must be true or false for assertion type %q", assertion.AssertionType.ValueString()),
				)
			}
		} else if assertion.AssertionType.ValueString() == string(AssertionTypeOutput) {
			if !assertion.Location.IsNull() {
				resp.Diagnostics.AddError(
					"Invalid Assertion Location Attribute",
					fmt.Sprintf("The location attribute is not applicable for assertion type %q", assertion.AssertionType.ValueString()),
				)
			}
		}
	}
}
