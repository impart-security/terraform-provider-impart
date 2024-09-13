package provider

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

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
	_ resource.Resource                = &specResource{}
	_ resource.ResourceWithConfigure   = &specResource{}
	_ resource.ResourceWithImportState = &specResource{}
)

// NewSpecResource is a helper function to simplify the provider implementation.
func NewSpecResource() resource.Resource {
	return &specResource{}
}

// specificationResource is the resource implementation.
type specResource struct {
	client *impartAPIClient
}

// specResourceModel maps the resource schema data.
type specResourceModel struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	SourceFile types.String `tfsdk:"source_file"`
	SourceHash types.String `tfsdk:"source_hash"`
}

// Configure adds the provider configured client to the resource.
func (r *specResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
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
func (r *specResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_spec"
}

// Schema defines the schema for the resource.
func (r *specResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage a specification.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for this specification.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name for this specification.",
				Required:    true,
			},

			"source_file": schema.StringAttribute{
				Description: "The specification file.",
				Required:    true,
			},

			"source_hash": schema.StringAttribute{
				Description: "The specification source hash.",
				Optional:    true,
			},
		},
	}
}

func (r *specResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Create a new resource.
func (r *specResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Preparing to create the specification resource")
	// Retrieve values from plan
	var plan specResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := plan.Name.ValueString()
	spec, err := os.ReadFile(plan.SourceFile.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read the specification source file",
			err.Error(),
		)
		return
	}

	specb64 := base64.StdEncoding.EncodeToString(spec)

	// Create new specification
	specRequest := r.client.SpecsAPI.CreateSpec(ctx, r.client.OrgID).
		SpecPostBody(openapiclient.SpecPostBody{
			Name: name,
			Spec: &specb64,
		})

	specResponse, _, err := specRequest.Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to create the specification",
			message,
		)
		return
	}

	// Map response body to model
	plan.ID = types.StringValue(specResponse.Id)
	plan.Name = types.StringValue(specResponse.Name)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Created the specification resource", map[string]any{"success": true})
}

// Read resource information.
func (r *specResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Preparing to read the specification resource")
	// Get current state
	var state specResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	currentHash := state.SourceHash

	specResponse, httpResp, err := r.client.SpecsAPI.GetSpec(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
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
			"Unable to read the specification",
			message,
		)
		return
	}

	// Map response body to model
	state = specResourceModel{
		ID:         types.StringValue(specResponse.Id),
		Name:       types.StringValue(specResponse.Name),
		SourceFile: state.SourceFile,
		SourceHash: state.SourceHash,
	}

	// track hash only if user originally set it
	if !currentHash.IsNull() {
		//state.SourceHash = types.StringValue(httpResp.Header.Get("ETag"))
		bytes, err := base64.StdEncoding.DecodeString(specResponse.Spec)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to base64 decode the specification",
				err.Error(),
			)
		}
		hash, err := calculateSha256(string(bytes))
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to calculate sha256",
				err.Error(),
			)
		}
		state.SourceHash = types.StringValue(hash)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Finished reading the specification resource", map[string]any{"success": true})
}

func (r *specResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Preparing to update the specification resource")
	// Retrieve values from plan
	var plan specResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := plan.Name.ValueString()
	spec, err := os.ReadFile(plan.SourceFile.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read the specification source file",
			err.Error(),
		)
		return
	}
	specb64 := base64.StdEncoding.EncodeToString(spec)

	specRequest := r.client.SpecsAPI.UpdateSpec(ctx, r.client.OrgID, plan.ID.ValueString()).
		SpecPostBody(openapiclient.SpecPostBody{
			Name: name,
			Spec: &specb64,
		})

	// update specification
	specResponse, _, err := specRequest.Execute()

	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to update the specification",
			message,
		)
		return
	}

	// Overwrite specifications with refreshed state
	plan = specResourceModel{
		ID:         types.StringValue(specResponse.Id),
		Name:       types.StringValue(specResponse.Name),
		SourceFile: types.StringValue(plan.SourceFile.ValueString()),
		SourceHash: plan.SourceHash,
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updated the specification resource", map[string]any{"success": true})
}

func (r *specResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Preparing to delete the specification resource")
	// Retrieve values from state
	var state specResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// delete specification
	_, err := r.client.SpecsAPI.DeleteSpec(ctx, r.client.OrgID, state.ID.ValueString()).Execute()
	if err != nil {
		message := err.Error()
		if apiErr, ok := err.(*openapiclient.GenericOpenAPIError); ok {
			message = fmt.Sprintf("%s %s", apiErr.Error(), string(apiErr.Body()))
		}

		resp.Diagnostics.AddError(
			"Unable to delete the specification",
			message,
		)
		return
	}

	tflog.Debug(ctx, "Deleted the specification resource", map[string]any{"success": true})
}
