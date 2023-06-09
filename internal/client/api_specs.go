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

type SpecsApi interface {

	/*
		CreateSpec Create a spec

		Creates a spec for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@return ApiCreateSpecRequest
	*/
	CreateSpec(ctx context.Context, orgId string) ApiCreateSpecRequest

	// CreateSpecExecute executes the request
	//  @return Spec
	CreateSpecExecute(r ApiCreateSpecRequest) (*Spec, *http.Response, error)

	/*
		DeleteSpec Delete a spec

		Deletes a spec for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@param specId Spec ID
		@return ApiDeleteSpecRequest
	*/
	DeleteSpec(ctx context.Context, orgId string, specId string) ApiDeleteSpecRequest

	// DeleteSpecExecute executes the request
	DeleteSpecExecute(r ApiDeleteSpecRequest) (*http.Response, error)

	/*
		GetSpec Get a spec

		Gets a spec for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@param specId Spec ID
		@return ApiGetSpecRequest
	*/
	GetSpec(ctx context.Context, orgId string, specId string) ApiGetSpecRequest

	// GetSpecExecute executes the request
	//  @return Spec
	GetSpecExecute(r ApiGetSpecRequest) (*Spec, *http.Response, error)

	/*
		UpdateSpec Update a spec

		Updates a spec for an organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Organization ID
		@param specId Spec ID
		@return ApiUpdateSpecRequest
	*/
	UpdateSpec(ctx context.Context, orgId string, specId string) ApiUpdateSpecRequest

	// UpdateSpecExecute executes the request
	//  @return Spec
	UpdateSpecExecute(r ApiUpdateSpecRequest) (*Spec, *http.Response, error)
}

// SpecsApiService SpecsApi service
type SpecsApiService service

type ApiCreateSpecRequest struct {
	ctx          context.Context
	ApiService   SpecsApi
	orgId        string
	specPostBody *SpecPostBody
}

// Request body for creating a new spec.
func (r ApiCreateSpecRequest) SpecPostBody(specPostBody SpecPostBody) ApiCreateSpecRequest {
	r.specPostBody = &specPostBody
	return r
}

func (r ApiCreateSpecRequest) Execute() (*Spec, *http.Response, error) {
	return r.ApiService.CreateSpecExecute(r)
}

/*
CreateSpec Create a spec

Creates a spec for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@return ApiCreateSpecRequest
*/
func (a *SpecsApiService) CreateSpec(ctx context.Context, orgId string) ApiCreateSpecRequest {
	return ApiCreateSpecRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return Spec
func (a *SpecsApiService) CreateSpecExecute(r ApiCreateSpecRequest) (*Spec, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Spec
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SpecsApiService.CreateSpec")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/specs"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if r.specPostBody == nil {
		return localVarReturnValue, nil, reportError("specPostBody is required and must be specified")
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
	localVarPostBody = r.specPostBody
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

type ApiDeleteSpecRequest struct {
	ctx        context.Context
	ApiService SpecsApi
	orgId      string
	specId     string
	ifMatch    *string
}

// If Match for ETag lock checks.
func (r ApiDeleteSpecRequest) IfMatch(ifMatch string) ApiDeleteSpecRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiDeleteSpecRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSpecExecute(r)
}

/*
DeleteSpec Delete a spec

Deletes a spec for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@param specId Spec ID
	@return ApiDeleteSpecRequest
*/
func (a *SpecsApiService) DeleteSpec(ctx context.Context, orgId string, specId string) ApiDeleteSpecRequest {
	return ApiDeleteSpecRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		specId:     specId,
	}
}

// Execute executes the request
func (a *SpecsApiService) DeleteSpecExecute(r ApiDeleteSpecRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SpecsApiService.DeleteSpec")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/specs/{spec_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"spec_id"+"}", url.PathEscape(parameterValueToString(r.specId, "specId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return nil, reportError("orgId must have less than 36 elements")
	}
	if strlen(r.specId) > 36 {
		return nil, reportError("specId must have less than 36 elements")
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

type ApiGetSpecRequest struct {
	ctx        context.Context
	ApiService SpecsApi
	orgId      string
	specId     string
	revision   *int32
}

// The revision to request
func (r ApiGetSpecRequest) Revision(revision int32) ApiGetSpecRequest {
	r.revision = &revision
	return r
}

func (r ApiGetSpecRequest) Execute() (*Spec, *http.Response, error) {
	return r.ApiService.GetSpecExecute(r)
}

/*
GetSpec Get a spec

Gets a spec for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@param specId Spec ID
	@return ApiGetSpecRequest
*/
func (a *SpecsApiService) GetSpec(ctx context.Context, orgId string, specId string) ApiGetSpecRequest {
	return ApiGetSpecRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		specId:     specId,
	}
}

// Execute executes the request
//
//	@return Spec
func (a *SpecsApiService) GetSpecExecute(r ApiGetSpecRequest) (*Spec, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Spec
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SpecsApiService.GetSpec")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/specs/{spec_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"spec_id"+"}", url.PathEscape(parameterValueToString(r.specId, "specId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if strlen(r.specId) > 36 {
		return localVarReturnValue, nil, reportError("specId must have less than 36 elements")
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

type ApiGetSpecAnalysesRequest struct {
	ctx          context.Context
	ApiService   SpecsApi
	orgId        string
	maxResults   *int32
	specId       *[]string
	collectionId *[]string
}

// The max number of results to return
func (r ApiGetSpecAnalysesRequest) MaxResults(maxResults int32) ApiGetSpecAnalysesRequest {
	r.maxResults = &maxResults
	return r
}

// Spec ID
func (r ApiGetSpecAnalysesRequest) SpecId(specId []string) ApiGetSpecAnalysesRequest {
	r.specId = &specId
	return r
}

// Collection ID
func (r ApiGetSpecAnalysesRequest) CollectionId(collectionId []string) ApiGetSpecAnalysesRequest {
	r.collectionId = &collectionId
	return r
}

/*
GetSpecAnalyses Get specs analysis

Gets a list of specs analysis for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@return ApiGetSpecAnalysesRequest
*/
func (a *SpecsApiService) GetSpecAnalyses(ctx context.Context, orgId string) ApiGetSpecAnalysesRequest {
	return ApiGetSpecAnalysesRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}


type ApiUpdateSpecRequest struct {
	ctx          context.Context
	ApiService   SpecsApi
	orgId        string
	specId       string
	specPostBody *SpecPostBody
	ifMatch      *string
}

// Request body for creating a new spec.
func (r ApiUpdateSpecRequest) SpecPostBody(specPostBody SpecPostBody) ApiUpdateSpecRequest {
	r.specPostBody = &specPostBody
	return r
}

// If Match for ETag lock checks.
func (r ApiUpdateSpecRequest) IfMatch(ifMatch string) ApiUpdateSpecRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiUpdateSpecRequest) Execute() (*Spec, *http.Response, error) {
	return r.ApiService.UpdateSpecExecute(r)
}

/*
UpdateSpec Update a spec

Updates a spec for an organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Organization ID
	@param specId Spec ID
	@return ApiUpdateSpecRequest
*/
func (a *SpecsApiService) UpdateSpec(ctx context.Context, orgId string, specId string) ApiUpdateSpecRequest {
	return ApiUpdateSpecRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		specId:     specId,
	}
}

// Execute executes the request
//
//	@return Spec
func (a *SpecsApiService) UpdateSpecExecute(r ApiUpdateSpecRequest) (*Spec, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Spec
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SpecsApiService.UpdateSpec")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/specs/{spec_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"spec_id"+"}", url.PathEscape(parameterValueToString(r.specId, "specId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) > 36 {
		return localVarReturnValue, nil, reportError("orgId must have less than 36 elements")
	}
	if strlen(r.specId) > 36 {
		return localVarReturnValue, nil, reportError("specId must have less than 36 elements")
	}
	if r.specPostBody == nil {
		return localVarReturnValue, nil, reportError("specPostBody is required and must be specified")
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
	localVarPostBody = r.specPostBody
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
