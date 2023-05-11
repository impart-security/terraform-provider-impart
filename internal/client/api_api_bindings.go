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

type ApiBindingsApi interface {

	/*
		CreateAPIBinding Create an API binding

		Creates an API binding for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@return ApiCreateAPIBindingRequest
	*/
	CreateAPIBinding(ctx context.Context, orgId string) ApiCreateAPIBindingRequest

	// CreateAPIBindingExecute executes the request
	//  @return ApiBinding
	CreateAPIBindingExecute(r ApiCreateAPIBindingRequest) (*ApiBinding, *http.Response, error)

	/*
		DeleteAPIBinding Delete an API binding

		Deletes an API binding from an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@param apiBindingId API Binding ID
		@return ApiDeleteAPIBindingRequest
	*/
	DeleteAPIBinding(ctx context.Context, orgId string, apiBindingId string) ApiDeleteAPIBindingRequest

	// DeleteAPIBindingExecute executes the request
	DeleteAPIBindingExecute(r ApiDeleteAPIBindingRequest) (*http.Response, error)

	/*
		GetAPIBinding Get an API binding

		Gets an API binding for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@param apiBindingId API Binding ID
		@return ApiGetAPIBindingRequest
	*/
	GetAPIBinding(ctx context.Context, orgId string, apiBindingId string) ApiGetAPIBindingRequest

	// GetAPIBindingExecute executes the request
	//  @return ApiBinding
	GetAPIBindingExecute(r ApiGetAPIBindingRequest) (*ApiBinding, *http.Response, error)

	/*
		UpdateAPIBinding Update an API binding

		Updates an API binding for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@param apiBindingId API Binding ID
		@return ApiUpdateAPIBindingRequest
	*/
	UpdateAPIBinding(ctx context.Context, orgId string, apiBindingId string) ApiUpdateAPIBindingRequest

	// UpdateAPIBindingExecute executes the request
	//  @return ApiBinding
	UpdateAPIBindingExecute(r ApiUpdateAPIBindingRequest) (*ApiBinding, *http.Response, error)
}

// ApiBindingsApiService ApiBindingsApi service
type ApiBindingsApiService service

type ApiCreateAPIBindingRequest struct {
	ctx                context.Context
	ApiService         ApiBindingsApi
	orgId              string
	apiBindingPostBody *ApiBindingPostBody
}

// Request body for creating an API binding.
func (r ApiCreateAPIBindingRequest) ApiBindingPostBody(apiBindingPostBody ApiBindingPostBody) ApiCreateAPIBindingRequest {
	r.apiBindingPostBody = &apiBindingPostBody
	return r
}

func (r ApiCreateAPIBindingRequest) Execute() (*ApiBinding, *http.Response, error) {
	return r.ApiService.CreateAPIBindingExecute(r)
}

/*
CreateAPIBinding Create an API binding

Creates an API binding for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@return ApiCreateAPIBindingRequest
*/
func (a *ApiBindingsApiService) CreateAPIBinding(ctx context.Context, orgId string) ApiCreateAPIBindingRequest {
	return ApiCreateAPIBindingRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return ApiBinding
func (a *ApiBindingsApiService) CreateAPIBindingExecute(r ApiCreateAPIBindingRequest) (*ApiBinding, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiBinding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiBindingsApiService.CreateAPIBinding")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/api_bindings"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if r.apiBindingPostBody == nil {
		return localVarReturnValue, nil, reportError("apiBindingPostBody is required and must be specified")
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
	localVarPostBody = r.apiBindingPostBody
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

type ApiDeleteAPIBindingRequest struct {
	ctx          context.Context
	ApiService   ApiBindingsApi
	orgId        string
	apiBindingId string
}

func (r ApiDeleteAPIBindingRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAPIBindingExecute(r)
}

/*
DeleteAPIBinding Delete an API binding

Deletes an API binding from an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@param apiBindingId API Binding ID
	@return ApiDeleteAPIBindingRequest
*/
func (a *ApiBindingsApiService) DeleteAPIBinding(ctx context.Context, orgId string, apiBindingId string) ApiDeleteAPIBindingRequest {
	return ApiDeleteAPIBindingRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        orgId,
		apiBindingId: apiBindingId,
	}
}

// Execute executes the request
func (a *ApiBindingsApiService) DeleteAPIBindingExecute(r ApiDeleteAPIBindingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiBindingsApiService.DeleteAPIBinding")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/api_bindings/{api_binding_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"api_binding_id"+"}", url.PathEscape(parameterValueToString(r.apiBindingId, "apiBindingId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return nil, reportError("orgId must have less than 36 elements")
	}
	if strlen(r.apiBindingId) > 36 {
		return nil, reportError("apiBindingId must have less than 36 elements")
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

type ApiGetAPIBindingRequest struct {
	ctx          context.Context
	ApiService   ApiBindingsApi
	orgId        string
	apiBindingId string
}

func (r ApiGetAPIBindingRequest) Execute() (*ApiBinding, *http.Response, error) {
	return r.ApiService.GetAPIBindingExecute(r)
}

/*
GetAPIBinding Get an API binding

Gets an API binding for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@param apiBindingId API Binding ID
	@return ApiGetAPIBindingRequest
*/
func (a *ApiBindingsApiService) GetAPIBinding(ctx context.Context, orgId string, apiBindingId string) ApiGetAPIBindingRequest {
	return ApiGetAPIBindingRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        orgId,
		apiBindingId: apiBindingId,
	}
}

// Execute executes the request
//
//	@return ApiBinding
func (a *ApiBindingsApiService) GetAPIBindingExecute(r ApiGetAPIBindingRequest) (*ApiBinding, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiBinding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiBindingsApiService.GetAPIBinding")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/api_bindings/{api_binding_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"api_binding_id"+"}", url.PathEscape(parameterValueToString(r.apiBindingId, "apiBindingId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if strlen(r.apiBindingId) > 36 {
		return localVarReturnValue, nil, reportError("apiBindingId must have less than 36 elements")
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

type ApiUpdateAPIBindingRequest struct {
	ctx                context.Context
	ApiService         ApiBindingsApi
	orgId              string
	apiBindingId       string
	apiBindingPostBody *ApiBindingPostBody
}

// Request body for creating an API binding.
func (r ApiUpdateAPIBindingRequest) ApiBindingPostBody(apiBindingPostBody ApiBindingPostBody) ApiUpdateAPIBindingRequest {
	r.apiBindingPostBody = &apiBindingPostBody
	return r
}

func (r ApiUpdateAPIBindingRequest) Execute() (*ApiBinding, *http.Response, error) {
	return r.ApiService.UpdateAPIBindingExecute(r)
}

/*
UpdateAPIBinding Update an API binding

Updates an API binding for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@param apiBindingId API Binding ID
	@return ApiUpdateAPIBindingRequest
*/
func (a *ApiBindingsApiService) UpdateAPIBinding(ctx context.Context, orgId string, apiBindingId string) ApiUpdateAPIBindingRequest {
	return ApiUpdateAPIBindingRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        orgId,
		apiBindingId: apiBindingId,
	}
}

// Execute executes the request
//
//	@return ApiBinding
func (a *ApiBindingsApiService) UpdateAPIBindingExecute(r ApiUpdateAPIBindingRequest) (*ApiBinding, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiBinding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiBindingsApiService.UpdateAPIBinding")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/api_bindings/{api_binding_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"api_binding_id"+"}", url.PathEscape(parameterValueToString(r.apiBindingId, "apiBindingId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if strlen(r.apiBindingId) > 36 {
		return localVarReturnValue, nil, reportError("apiBindingId must have less than 36 elements")
	}
	if r.apiBindingPostBody == nil {
		return localVarReturnValue, nil, reportError("apiBindingPostBody is required and must be specified")
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
	localVarPostBody = r.apiBindingPostBody
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
