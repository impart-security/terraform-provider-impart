/*
Impart Security v0 REST API

Imparts v0 REST API.

API version: 0.0.0
Contact: support@impart.security
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RuleRecipesAPI interface {

	/*
		CreateRuleRecipe Create a new rule recipe

		Creates a new rule recipe for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@return ApiCreateRuleRecipeRequest
	*/
	CreateRuleRecipe(ctx context.Context, orgId string) ApiCreateRuleRecipeRequest

	// CreateRuleRecipeExecute executes the request
	//  @return RuleRecipe
	CreateRuleRecipeExecute(r ApiCreateRuleRecipeRequest) (*RuleRecipe, *http.Response, error)

	/*
		DeleteRuleRecipe Delete a rule recipe

		Deletes a rule recipe for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@param ruleRecipeId Rule recipe ID
		@return ApiDeleteRuleRecipeRequest
	*/
	DeleteRuleRecipe(ctx context.Context, orgId string, ruleRecipeId string) ApiDeleteRuleRecipeRequest

	// DeleteRuleRecipeExecute executes the request
	DeleteRuleRecipeExecute(r ApiDeleteRuleRecipeRequest) (*http.Response, error)

	/*
		GenerateRuleRecipeScript Generate a rule recipe script

		Generates a rule recipe script for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@return ApiGenerateRuleRecipeScriptRequest
	*/
	GenerateRuleRecipeScript(ctx context.Context, orgId string) ApiGenerateRuleRecipeScriptRequest

	// GenerateRuleRecipeScriptExecute executes the request
	//  @return RuleRecipeScript
	GenerateRuleRecipeScriptExecute(r ApiGenerateRuleRecipeScriptRequest) (*RuleRecipeScript, *http.Response, error)

	/*
		GetRuleRecipe Get a rule recipe

		Gets a rule recipe for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@param ruleRecipeId Rule recipe ID
		@return ApiGetRuleRecipeRequest
	*/
	GetRuleRecipe(ctx context.Context, orgId string, ruleRecipeId string) ApiGetRuleRecipeRequest

	// GetRuleRecipeExecute executes the request
	//  @return RuleRecipe
	GetRuleRecipeExecute(r ApiGetRuleRecipeRequest) (*RuleRecipe, *http.Response, error)

	/*
		GetRuleRecipes Get rule recipes

		Gets a list of rule recipes for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@return ApiGetRuleRecipesRequest
	*/
	GetRuleRecipes(ctx context.Context, orgId string) ApiGetRuleRecipesRequest

	// GetRuleRecipesExecute executes the request
	//  @return RuleRecipes
	GetRuleRecipesExecute(r ApiGetRuleRecipesRequest) (*RuleRecipes, *http.Response, error)

	/*
		UpdateRuleRecipe Update a rule recipe

		Updates a rule recipe for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@param ruleRecipeId Rule recipe ID
		@return ApiUpdateRuleRecipeRequest
	*/
	UpdateRuleRecipe(ctx context.Context, orgId string, ruleRecipeId string) ApiUpdateRuleRecipeRequest

	// UpdateRuleRecipeExecute executes the request
	//  @return RuleRecipe
	UpdateRuleRecipeExecute(r ApiUpdateRuleRecipeRequest) (*RuleRecipe, *http.Response, error)
}

// RuleRecipesAPIService RuleRecipesAPI service
type RuleRecipesAPIService service

type ApiCreateRuleRecipeRequest struct {
	ctx                context.Context
	ApiService         RuleRecipesAPI
	orgId              string
	ruleRecipePostBody *RuleRecipePostBody
}

// Request body for creating a rule recipe.
func (r ApiCreateRuleRecipeRequest) RuleRecipePostBody(ruleRecipePostBody RuleRecipePostBody) ApiCreateRuleRecipeRequest {
	r.ruleRecipePostBody = &ruleRecipePostBody
	return r
}

func (r ApiCreateRuleRecipeRequest) Execute() (*RuleRecipe, *http.Response, error) {
	return r.ApiService.CreateRuleRecipeExecute(r)
}

/*
CreateRuleRecipe Create a new rule recipe

Creates a new rule recipe for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@return ApiCreateRuleRecipeRequest
*/
func (a *RuleRecipesAPIService) CreateRuleRecipe(ctx context.Context, orgId string) ApiCreateRuleRecipeRequest {
	return ApiCreateRuleRecipeRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return RuleRecipe
func (a *RuleRecipesAPIService) CreateRuleRecipeExecute(r ApiCreateRuleRecipeRequest) (*RuleRecipe, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RuleRecipe
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RuleRecipesAPIService.CreateRuleRecipe")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/rule_recipes"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if r.ruleRecipePostBody == nil {
		return localVarReturnValue, nil, reportError("ruleRecipePostBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/security.impart.api.v0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.ruleRecipePostBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v CompilationDiagnostics
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteRuleRecipeRequest struct {
	ctx          context.Context
	ApiService   RuleRecipesAPI
	orgId        string
	ruleRecipeId string
	ifMatch      *string
}

// If Match for ETag lock checks.
func (r ApiDeleteRuleRecipeRequest) IfMatch(ifMatch string) ApiDeleteRuleRecipeRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiDeleteRuleRecipeRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRuleRecipeExecute(r)
}

/*
DeleteRuleRecipe Delete a rule recipe

Deletes a rule recipe for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@param ruleRecipeId Rule recipe ID
	@return ApiDeleteRuleRecipeRequest
*/
func (a *RuleRecipesAPIService) DeleteRuleRecipe(ctx context.Context, orgId string, ruleRecipeId string) ApiDeleteRuleRecipeRequest {
	return ApiDeleteRuleRecipeRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        orgId,
		ruleRecipeId: ruleRecipeId,
	}
}

// Execute executes the request
func (a *RuleRecipesAPIService) DeleteRuleRecipeExecute(r ApiDeleteRuleRecipeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RuleRecipesAPIService.DeleteRuleRecipe")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/rule_recipes/{rule_recipe_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rule_recipe_id"+"}", url.PathEscape(parameterValueToString(r.ruleRecipeId, "ruleRecipeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return nil, reportError("orgId must have less than 36 elements")
	}
	if strlen(r.ruleRecipeId) > 36 {
		return nil, reportError("ruleRecipeId must have less than 36 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/security.impart.api.v0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ifMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGenerateRuleRecipeScriptRequest struct {
	ctx                context.Context
	ApiService         RuleRecipesAPI
	orgId              string
	ruleRecipePostBody *RuleRecipePostBody
}

// Request body for creating a rule recipe.
func (r ApiGenerateRuleRecipeScriptRequest) RuleRecipePostBody(ruleRecipePostBody RuleRecipePostBody) ApiGenerateRuleRecipeScriptRequest {
	r.ruleRecipePostBody = &ruleRecipePostBody
	return r
}

func (r ApiGenerateRuleRecipeScriptRequest) Execute() (*RuleRecipeScript, *http.Response, error) {
	return r.ApiService.GenerateRuleRecipeScriptExecute(r)
}

/*
GenerateRuleRecipeScript Generate a rule recipe script

Generates a rule recipe script for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@return ApiGenerateRuleRecipeScriptRequest
*/
func (a *RuleRecipesAPIService) GenerateRuleRecipeScript(ctx context.Context, orgId string) ApiGenerateRuleRecipeScriptRequest {
	return ApiGenerateRuleRecipeScriptRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return RuleRecipeScript
func (a *RuleRecipesAPIService) GenerateRuleRecipeScriptExecute(r ApiGenerateRuleRecipeScriptRequest) (*RuleRecipeScript, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RuleRecipeScript
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RuleRecipesAPIService.GenerateRuleRecipeScript")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/rule_recipe_generate_script"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if r.ruleRecipePostBody == nil {
		return localVarReturnValue, nil, reportError("ruleRecipePostBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/security.impart.api.v0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.ruleRecipePostBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v CompilationDiagnostics
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetRuleRecipeRequest struct {
	ctx          context.Context
	ApiService   RuleRecipesAPI
	orgId        string
	ruleRecipeId string
}

func (r ApiGetRuleRecipeRequest) Execute() (*RuleRecipe, *http.Response, error) {
	return r.ApiService.GetRuleRecipeExecute(r)
}

/*
GetRuleRecipe Get a rule recipe

Gets a rule recipe for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@param ruleRecipeId Rule recipe ID
	@return ApiGetRuleRecipeRequest
*/
func (a *RuleRecipesAPIService) GetRuleRecipe(ctx context.Context, orgId string, ruleRecipeId string) ApiGetRuleRecipeRequest {
	return ApiGetRuleRecipeRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        orgId,
		ruleRecipeId: ruleRecipeId,
	}
}

// Execute executes the request
//
//	@return RuleRecipe
func (a *RuleRecipesAPIService) GetRuleRecipeExecute(r ApiGetRuleRecipeRequest) (*RuleRecipe, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RuleRecipe
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RuleRecipesAPIService.GetRuleRecipe")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/rule_recipes/{rule_recipe_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rule_recipe_id"+"}", url.PathEscape(parameterValueToString(r.ruleRecipeId, "ruleRecipeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if strlen(r.ruleRecipeId) > 36 {
		return localVarReturnValue, nil, reportError("ruleRecipeId must have less than 36 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/security.impart.api.v0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetRuleRecipesRequest struct {
	ctx        context.Context
	ApiService RuleRecipesAPI
	orgId      string
	page       *int32
	maxResults *int32
	search     *string
}

// The page of results to return
func (r ApiGetRuleRecipesRequest) Page(page int32) ApiGetRuleRecipesRequest {
	r.page = &page
	return r
}

// The max number of results to return
func (r ApiGetRuleRecipesRequest) MaxResults(maxResults int32) ApiGetRuleRecipesRequest {
	r.maxResults = &maxResults
	return r
}

// Search string
func (r ApiGetRuleRecipesRequest) Search(search string) ApiGetRuleRecipesRequest {
	r.search = &search
	return r
}

func (r ApiGetRuleRecipesRequest) Execute() (*RuleRecipes, *http.Response, error) {
	return r.ApiService.GetRuleRecipesExecute(r)
}

/*
GetRuleRecipes Get rule recipes

Gets a list of rule recipes for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@return ApiGetRuleRecipesRequest
*/
func (a *RuleRecipesAPIService) GetRuleRecipes(ctx context.Context, orgId string) ApiGetRuleRecipesRequest {
	return ApiGetRuleRecipesRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return RuleRecipes
func (a *RuleRecipesAPIService) GetRuleRecipesExecute(r ApiGetRuleRecipesRequest) (*RuleRecipes, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RuleRecipes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RuleRecipesAPIService.GetRuleRecipes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/rule_recipes"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.maxResults != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max_results", r.maxResults, "", "")
	} else {
		var defaultValue int32 = 100
		r.maxResults = &defaultValue
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/security.impart.api.v0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateRuleRecipeRequest struct {
	ctx                context.Context
	ApiService         RuleRecipesAPI
	orgId              string
	ruleRecipeId       string
	ruleRecipePostBody *RuleRecipePostBody
	ifMatch            *string
}

// Request body for creating a rule recipe.
func (r ApiUpdateRuleRecipeRequest) RuleRecipePostBody(ruleRecipePostBody RuleRecipePostBody) ApiUpdateRuleRecipeRequest {
	r.ruleRecipePostBody = &ruleRecipePostBody
	return r
}

// If Match for ETag lock checks.
func (r ApiUpdateRuleRecipeRequest) IfMatch(ifMatch string) ApiUpdateRuleRecipeRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiUpdateRuleRecipeRequest) Execute() (*RuleRecipe, *http.Response, error) {
	return r.ApiService.UpdateRuleRecipeExecute(r)
}

/*
UpdateRuleRecipe Update a rule recipe

Updates a rule recipe for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@param ruleRecipeId Rule recipe ID
	@return ApiUpdateRuleRecipeRequest
*/
func (a *RuleRecipesAPIService) UpdateRuleRecipe(ctx context.Context, orgId string, ruleRecipeId string) ApiUpdateRuleRecipeRequest {
	return ApiUpdateRuleRecipeRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        orgId,
		ruleRecipeId: ruleRecipeId,
	}
}

// Execute executes the request
//
//	@return RuleRecipe
func (a *RuleRecipesAPIService) UpdateRuleRecipeExecute(r ApiUpdateRuleRecipeRequest) (*RuleRecipe, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RuleRecipe
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RuleRecipesAPIService.UpdateRuleRecipe")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/rule_recipes/{rule_recipe_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rule_recipe_id"+"}", url.PathEscape(parameterValueToString(r.ruleRecipeId, "ruleRecipeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if strlen(r.ruleRecipeId) > 36 {
		return localVarReturnValue, nil, reportError("ruleRecipeId must have less than 36 elements")
	}
	if r.ruleRecipePostBody == nil {
		return localVarReturnValue, nil, reportError("ruleRecipePostBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/security.impart.api.v0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ifMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "", "")
	}
	// body params
	localVarPostBody = r.ruleRecipePostBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v CompilationDiagnostics
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
