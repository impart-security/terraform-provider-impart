/*
Impart Security v0 REST API

Imparts v0 REST API.

API version: 0.0.0
Contact: support@impart.security
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RulesScriptsAPI interface {

	/*
		CreateRulesScript Create a new rules script

		Creates a new rules script for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@return ApiCreateRulesScriptRequest
	*/
	CreateRulesScript(ctx context.Context, orgId string) ApiCreateRulesScriptRequest

	// CreateRulesScriptExecute executes the request
	//  @return RulesScript
	CreateRulesScriptExecute(r ApiCreateRulesScriptRequest) (*RulesScript, *http.Response, error)

	/*
		DeleteRulesScript Delete a rules script

		Deletes a rules script for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@param rulesScriptId Rules script ID
		@return ApiDeleteRulesScriptRequest
	*/
	DeleteRulesScript(ctx context.Context, orgId string, rulesScriptId string) ApiDeleteRulesScriptRequest

	// DeleteRulesScriptExecute executes the request
	DeleteRulesScriptExecute(r ApiDeleteRulesScriptRequest) (*http.Response, error)

	/*
		GetRulesScript Get a rules script

		Gets a rules script for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@param rulesScriptId Rules script ID
		@return ApiGetRulesScriptRequest
	*/
	GetRulesScript(ctx context.Context, orgId string, rulesScriptId string) ApiGetRulesScriptRequest

	// GetRulesScriptExecute executes the request
	//  @return RulesScript
	GetRulesScriptExecute(r ApiGetRulesScriptRequest) (*RulesScript, *http.Response, error)

	/*
		GetRulesScripts Get rules scripts

		Gets a list of rules scripts for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@return ApiGetRulesScriptsRequest
	*/
	GetRulesScripts(ctx context.Context, orgId string) ApiGetRulesScriptsRequest

	// GetRulesScriptsExecute executes the request
	//  @return RulesScripts
	GetRulesScriptsExecute(r ApiGetRulesScriptsRequest) (*RulesScripts, *http.Response, error)

	/*
		RulesTest Test a rule

		Test a rule.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@return ApiRulesTestRequest
	*/
	RulesTest(ctx context.Context, orgId string) ApiRulesTestRequest

	// RulesTestExecute executes the request
	//  @return RulesTest
	RulesTestExecute(r ApiRulesTestRequest) (*RulesTest, *http.Response, error)

	/*
		UpdateRulesScript Update a rules script

		Updates a rules script for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@param rulesScriptId Rules script ID
		@return ApiUpdateRulesScriptRequest
	*/
	UpdateRulesScript(ctx context.Context, orgId string, rulesScriptId string) ApiUpdateRulesScriptRequest

	// UpdateRulesScriptExecute executes the request
	//  @return RulesScript
	UpdateRulesScriptExecute(r ApiUpdateRulesScriptRequest) (*RulesScript, *http.Response, error)
}

// RulesScriptsAPIService RulesScriptsAPI service
type RulesScriptsAPIService service

type ApiCreateRulesScriptRequest struct {
	ctx                 context.Context
	ApiService          RulesScriptsAPI
	orgId               string
	rulesScriptPostBody *RulesScriptPostBody
}

// Request body for creating a rules script.
func (r ApiCreateRulesScriptRequest) RulesScriptPostBody(rulesScriptPostBody RulesScriptPostBody) ApiCreateRulesScriptRequest {
	r.rulesScriptPostBody = &rulesScriptPostBody
	return r
}

func (r ApiCreateRulesScriptRequest) Execute() (*RulesScript, *http.Response, error) {
	return r.ApiService.CreateRulesScriptExecute(r)
}

/*
CreateRulesScript Create a new rules script

Creates a new rules script for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@return ApiCreateRulesScriptRequest
*/
func (a *RulesScriptsAPIService) CreateRulesScript(ctx context.Context, orgId string) ApiCreateRulesScriptRequest {
	return ApiCreateRulesScriptRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return RulesScript
func (a *RulesScriptsAPIService) CreateRulesScriptExecute(r ApiCreateRulesScriptRequest) (*RulesScript, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RulesScript
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesScriptsAPIService.CreateRulesScript")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/rules_scripts"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if r.rulesScriptPostBody == nil {
		return localVarReturnValue, nil, reportError("rulesScriptPostBody is required and must be specified")
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
	localVarPostBody = r.rulesScriptPostBody
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

type ApiDeleteRulesScriptRequest struct {
	ctx           context.Context
	ApiService    RulesScriptsAPI
	orgId         string
	rulesScriptId string
	ifMatch       *string
}

// If Match for ETag lock checks.
func (r ApiDeleteRulesScriptRequest) IfMatch(ifMatch string) ApiDeleteRulesScriptRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiDeleteRulesScriptRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRulesScriptExecute(r)
}

/*
DeleteRulesScript Delete a rules script

Deletes a rules script for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@param rulesScriptId Rules script ID
	@return ApiDeleteRulesScriptRequest
*/
func (a *RulesScriptsAPIService) DeleteRulesScript(ctx context.Context, orgId string, rulesScriptId string) ApiDeleteRulesScriptRequest {
	return ApiDeleteRulesScriptRequest{
		ApiService:    a,
		ctx:           ctx,
		orgId:         orgId,
		rulesScriptId: rulesScriptId,
	}
}

// Execute executes the request
func (a *RulesScriptsAPIService) DeleteRulesScriptExecute(r ApiDeleteRulesScriptRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesScriptsAPIService.DeleteRulesScript")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/rules_scripts/{rules_script_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rules_script_id"+"}", url.PathEscape(parameterValueToString(r.rulesScriptId, "rulesScriptId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return nil, reportError("orgId must have less than 36 elements")
	}
	if strlen(r.rulesScriptId) > 36 {
		return nil, reportError("rulesScriptId must have less than 36 elements")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "")
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

type ApiGetRulesScriptRequest struct {
	ctx           context.Context
	ApiService    RulesScriptsAPI
	orgId         string
	rulesScriptId string
	revision      *int32
}

// The revision to request
func (r ApiGetRulesScriptRequest) Revision(revision int32) ApiGetRulesScriptRequest {
	r.revision = &revision
	return r
}

func (r ApiGetRulesScriptRequest) Execute() (*RulesScript, *http.Response, error) {
	return r.ApiService.GetRulesScriptExecute(r)
}

/*
GetRulesScript Get a rules script

Gets a rules script for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@param rulesScriptId Rules script ID
	@return ApiGetRulesScriptRequest
*/
func (a *RulesScriptsAPIService) GetRulesScript(ctx context.Context, orgId string, rulesScriptId string) ApiGetRulesScriptRequest {
	return ApiGetRulesScriptRequest{
		ApiService:    a,
		ctx:           ctx,
		orgId:         orgId,
		rulesScriptId: rulesScriptId,
	}
}

// Execute executes the request
//
//	@return RulesScript
func (a *RulesScriptsAPIService) GetRulesScriptExecute(r ApiGetRulesScriptRequest) (*RulesScript, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RulesScript
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesScriptsAPIService.GetRulesScript")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/rules_scripts/{rules_script_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rules_script_id"+"}", url.PathEscape(parameterValueToString(r.rulesScriptId, "rulesScriptId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if strlen(r.rulesScriptId) > 36 {
		return localVarReturnValue, nil, reportError("rulesScriptId must have less than 36 elements")
	}

	if r.revision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "revision", r.revision, "")
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

type ApiGetRulesScriptsRequest struct {
	ctx              context.Context
	ApiService       RulesScriptsAPI
	orgId            string
	page             *int32
	maxResults       *int32
	type_            *string
	excludeSrc       *bool
	excludeRevisions *bool
}

// The page of results to return
func (r ApiGetRulesScriptsRequest) Page(page int32) ApiGetRulesScriptsRequest {
	r.page = &page
	return r
}

// The max number of results to return
func (r ApiGetRulesScriptsRequest) MaxResults(maxResults int32) ApiGetRulesScriptsRequest {
	r.maxResults = &maxResults
	return r
}

// Type of rule script to filter results by
func (r ApiGetRulesScriptsRequest) Type_(type_ string) ApiGetRulesScriptsRequest {
	r.type_ = &type_
	return r
}

// Whether to exclude the rule script src from the response
func (r ApiGetRulesScriptsRequest) ExcludeSrc(excludeSrc bool) ApiGetRulesScriptsRequest {
	r.excludeSrc = &excludeSrc
	return r
}

// Whether to exclude the rule script revisions from the response
func (r ApiGetRulesScriptsRequest) ExcludeRevisions(excludeRevisions bool) ApiGetRulesScriptsRequest {
	r.excludeRevisions = &excludeRevisions
	return r
}

func (r ApiGetRulesScriptsRequest) Execute() (*RulesScripts, *http.Response, error) {
	return r.ApiService.GetRulesScriptsExecute(r)
}

/*
GetRulesScripts Get rules scripts

Gets a list of rules scripts for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@return ApiGetRulesScriptsRequest
*/
func (a *RulesScriptsAPIService) GetRulesScripts(ctx context.Context, orgId string) ApiGetRulesScriptsRequest {
	return ApiGetRulesScriptsRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return RulesScripts
func (a *RulesScriptsAPIService) GetRulesScriptsExecute(r ApiGetRulesScriptsRequest) (*RulesScripts, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RulesScripts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesScriptsAPIService.GetRulesScripts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/rules_scripts"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 0
		r.page = &defaultValue
	}
	if r.maxResults != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max_results", r.maxResults, "")
	} else {
		var defaultValue int32 = 100
		r.maxResults = &defaultValue
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.excludeSrc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_src", r.excludeSrc, "")
	} else {
		var defaultValue bool = false
		r.excludeSrc = &defaultValue
	}
	if r.excludeRevisions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_revisions", r.excludeRevisions, "")
	} else {
		var defaultValue bool = false
		r.excludeRevisions = &defaultValue
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

type ApiRulesTestRequest struct {
	ctx               context.Context
	ApiService        RulesScriptsAPI
	orgId             string
	rulesTestPostBody *RulesTestPostBody
}

// Request body for testing rules.
func (r ApiRulesTestRequest) RulesTestPostBody(rulesTestPostBody RulesTestPostBody) ApiRulesTestRequest {
	r.rulesTestPostBody = &rulesTestPostBody
	return r
}

func (r ApiRulesTestRequest) Execute() (*RulesTest, *http.Response, error) {
	return r.ApiService.RulesTestExecute(r)
}

/*
RulesTest Test a rule

Test a rule.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@return ApiRulesTestRequest
*/
func (a *RulesScriptsAPIService) RulesTest(ctx context.Context, orgId string) ApiRulesTestRequest {
	return ApiRulesTestRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return RulesTest
func (a *RulesScriptsAPIService) RulesTestExecute(r ApiRulesTestRequest) (*RulesTest, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RulesTest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesScriptsAPIService.RulesTest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/rules_test"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if r.rulesTestPostBody == nil {
		return localVarReturnValue, nil, reportError("rulesTestPostBody is required and must be specified")
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
	localVarPostBody = r.rulesTestPostBody
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

type ApiUpdateRulesScriptRequest struct {
	ctx                 context.Context
	ApiService          RulesScriptsAPI
	orgId               string
	rulesScriptId       string
	rulesScriptPostBody *RulesScriptPostBody
	ifMatch             *string
}

// Request body for creating a rules script.
func (r ApiUpdateRulesScriptRequest) RulesScriptPostBody(rulesScriptPostBody RulesScriptPostBody) ApiUpdateRulesScriptRequest {
	r.rulesScriptPostBody = &rulesScriptPostBody
	return r
}

// If Match for ETag lock checks.
func (r ApiUpdateRulesScriptRequest) IfMatch(ifMatch string) ApiUpdateRulesScriptRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiUpdateRulesScriptRequest) Execute() (*RulesScript, *http.Response, error) {
	return r.ApiService.UpdateRulesScriptExecute(r)
}

/*
UpdateRulesScript Update a rules script

Updates a rules script for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@param rulesScriptId Rules script ID
	@return ApiUpdateRulesScriptRequest
*/
func (a *RulesScriptsAPIService) UpdateRulesScript(ctx context.Context, orgId string, rulesScriptId string) ApiUpdateRulesScriptRequest {
	return ApiUpdateRulesScriptRequest{
		ApiService:    a,
		ctx:           ctx,
		orgId:         orgId,
		rulesScriptId: rulesScriptId,
	}
}

// Execute executes the request
//
//	@return RulesScript
func (a *RulesScriptsAPIService) UpdateRulesScriptExecute(r ApiUpdateRulesScriptRequest) (*RulesScript, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RulesScript
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RulesScriptsAPIService.UpdateRulesScript")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/rules_scripts/{rules_script_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rules_script_id"+"}", url.PathEscape(parameterValueToString(r.rulesScriptId, "rulesScriptId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if strlen(r.rulesScriptId) > 36 {
		return localVarReturnValue, nil, reportError("rulesScriptId must have less than 36 elements")
	}
	if r.rulesScriptPostBody == nil {
		return localVarReturnValue, nil, reportError("rulesScriptPostBody is required and must be specified")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "")
	}
	// body params
	localVarPostBody = r.rulesScriptPostBody
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
