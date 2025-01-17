/*
Permit.io API

 Authorization as a service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"github.com/permitio/permit-golang/pkg/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// ElementsConfigsApiService ElementsConfigsApi service
type ElementsConfigsApiService service

type ApiCreateElementsConfigRequest struct {
	ctx                  context.Context
	ApiService           *ElementsConfigsApiService
	projId               string
	envId                string
	elementsConfigCreate *models.ElementsConfigCreate
}

func (r ApiCreateElementsConfigRequest) ElementsConfigCreate(elementsConfigCreate models.ElementsConfigCreate) ApiCreateElementsConfigRequest {
	r.elementsConfigCreate = &elementsConfigCreate
	return r
}

func (r ApiCreateElementsConfigRequest) Execute() (*models.ElementsConfigRead, *http.Response, error) {
	return r.ApiService.CreateElementsConfigExecute(r)
}

/*
CreateElementsConfig Create Elements Config

Creates a new elements_config under the active organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @return ApiCreateElementsConfigRequest
*/
func (a *ElementsConfigsApiService) CreateElementsConfig(ctx context.Context, projId string, envId string) ApiCreateElementsConfigRequest {
	return ApiCreateElementsConfigRequest{
		ApiService: a,
		ctx:        ctx,
		projId:     projId,
		envId:      envId,
	}
}

// Execute executes the request
//  @return ElementsConfigRead
func (a *ElementsConfigsApiService) CreateElementsConfigExecute(r ApiCreateElementsConfigRequest) (*models.ElementsConfigRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ElementsConfigRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsConfigsApiService.CreateElementsConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}/config"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.elementsConfigCreate == nil {
		return localVarReturnValue, nil, reportError("elementsConfigCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.elementsConfigCreate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
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

type ApiDeleteElementsConfigRequest struct {
	ctx              context.Context
	ApiService       *ElementsConfigsApiService
	elementsConfigId string
	projId           string
	envId            string
}

func (r ApiDeleteElementsConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteElementsConfigExecute(r)
}

/*
DeleteElementsConfig Delete Elements Config

Deletes the elements_config and all its related data.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param elementsConfigId Either the unique id of the elements_config, or the URL-friendly key of the elements_config (i.e: the \"slug\").
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @return ApiDeleteElementsConfigRequest
*/
func (a *ElementsConfigsApiService) DeleteElementsConfig(ctx context.Context, elementsConfigId string, projId string, envId string) ApiDeleteElementsConfigRequest {
	return ApiDeleteElementsConfigRequest{
		ApiService:       a,
		ctx:              ctx,
		elementsConfigId: elementsConfigId,
		projId:           projId,
		envId:            envId,
	}
}

// Execute executes the request
func (a *ElementsConfigsApiService) DeleteElementsConfigExecute(r ApiDeleteElementsConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsConfigsApiService.DeleteElementsConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}/{elements_config_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"elements_config_id"+"}", url.PathEscape(parameterToString(r.elementsConfigId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
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

type ApiGetElementsConfigRequest struct {
	ctx              context.Context
	ApiService       *ElementsConfigsApiService
	projId           string
	envId            string
	elementsConfigId string
}

func (r ApiGetElementsConfigRequest) Execute() (*models.ElementsConfigRead, *http.Response, error) {
	return r.ApiService.GetElementsConfigExecute(r)
}

/*
GetElementsConfig Get Elements Config

Gets a single elements_config matching the given elements_config_id, if such elements_config exists.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @param elementsConfigId Either the unique id of the elements_config, or the URL-friendly key of the elements_config (i.e: the \"slug\").
 @return ApiGetElementsConfigRequest
*/
func (a *ElementsConfigsApiService) GetElementsConfig(ctx context.Context, projId string, envId string, elementsConfigId string) ApiGetElementsConfigRequest {
	return ApiGetElementsConfigRequest{
		ApiService:       a,
		ctx:              ctx,
		projId:           projId,
		envId:            envId,
		elementsConfigId: elementsConfigId,
	}
}

// Execute executes the request
//  @return ElementsConfigRead
func (a *ElementsConfigsApiService) GetElementsConfigExecute(r ApiGetElementsConfigRequest) (*models.ElementsConfigRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ElementsConfigRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsConfigsApiService.GetElementsConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}/config/{elements_config_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"elements_config_id"+"}", url.PathEscape(parameterToString(r.elementsConfigId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
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

type ApiGetEnvConfigRequest struct {
	ctx        context.Context
	ApiService *ElementsConfigsApiService
	projId     string
	envId      string
}

func (r ApiGetEnvConfigRequest) Execute() (*models.ElementsEnvRead, *http.Response, error) {
	return r.ApiService.GetEnvConfigExecute(r)
}

/*
GetEnvConfig Get Env Config

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @return ApiGetEnvConfigRequest
*/
func (a *ElementsConfigsApiService) GetEnvConfig(ctx context.Context, projId string, envId string) ApiGetEnvConfigRequest {
	return ApiGetEnvConfigRequest{
		ApiService: a,
		ctx:        ctx,
		projId:     projId,
		envId:      envId,
	}
}

// Execute executes the request
//  @return ElementsEnvRead
func (a *ElementsConfigsApiService) GetEnvConfigExecute(r ApiGetEnvConfigRequest) (*models.ElementsEnvRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ElementsEnvRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsConfigsApiService.GetEnvConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
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

type ApiListElementsConfigsRequest struct {
	ctx        context.Context
	ApiService *ElementsConfigsApiService
	projId     string
	envId      string
	page       *int32
	perPage    *int32
}

// Page number of the results to fetch, starting at 1.
func (r ApiListElementsConfigsRequest) Page(page int32) ApiListElementsConfigsRequest {
	r.page = &page
	return r
}

// The number of results per page (max 100).
func (r ApiListElementsConfigsRequest) PerPage(perPage int32) ApiListElementsConfigsRequest {
	r.perPage = &perPage
	return r
}

func (r ApiListElementsConfigsRequest) Execute() (*models.PaginatedResultElementsConfigRead, *http.Response, error) {
	return r.ApiService.ListElementsConfigsExecute(r)
}

/*
ListElementsConfigs List Elements Configs

Lists all the elements_configs under the active organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @return ApiListElementsConfigsRequest
*/
func (a *ElementsConfigsApiService) ListElementsConfigs(ctx context.Context, projId string, envId string) ApiListElementsConfigsRequest {
	return ApiListElementsConfigsRequest{
		ApiService: a,
		ctx:        ctx,
		projId:     projId,
		envId:      envId,
	}
}

// Execute executes the request
//  @return PaginatedResultElementsConfigRead
func (a *ElementsConfigsApiService) ListElementsConfigsExecute(r ApiListElementsConfigsRequest) (*models.PaginatedResultElementsConfigRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.PaginatedResultElementsConfigRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsConfigsApiService.ListElementsConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}/config"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
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

type ApiUpdateElementsConfigRequest struct {
	ctx                  context.Context
	ApiService           *ElementsConfigsApiService
	elementsConfigId     string
	projId               string
	envId                string
	elementsConfigUpdate *models.ElementsConfigUpdate
}

func (r ApiUpdateElementsConfigRequest) ElementsConfigUpdate(elementsConfigUpdate models.ElementsConfigUpdate) ApiUpdateElementsConfigRequest {
	r.elementsConfigUpdate = &elementsConfigUpdate
	return r
}

func (r ApiUpdateElementsConfigRequest) Execute() (*models.ElementsConfigRead, *http.Response, error) {
	return r.ApiService.UpdateElementsConfigExecute(r)
}

/*
UpdateElementsConfig Update Elements Config

Updates the elements_config.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param elementsConfigId Either the unique id of the elements_config, or the URL-friendly key of the elements_config (i.e: the \"slug\").
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @return ApiUpdateElementsConfigRequest
*/
func (a *ElementsConfigsApiService) UpdateElementsConfig(ctx context.Context, elementsConfigId string, projId string, envId string) ApiUpdateElementsConfigRequest {
	return ApiUpdateElementsConfigRequest{
		ApiService:       a,
		ctx:              ctx,
		elementsConfigId: elementsConfigId,
		projId:           projId,
		envId:            envId,
	}
}

// Execute executes the request
//  @return ElementsConfigRead
func (a *ElementsConfigsApiService) UpdateElementsConfigExecute(r ApiUpdateElementsConfigRequest) (*models.ElementsConfigRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ElementsConfigRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsConfigsApiService.UpdateElementsConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}/config/{elements_config_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"elements_config_id"+"}", url.PathEscape(parameterToString(r.elementsConfigId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.elementsConfigUpdate == nil {
		return localVarReturnValue, nil, reportError("elementsConfigUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.elementsConfigUpdate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
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

type ApiUpdateElementsEnvRequest struct {
	ctx               context.Context
	ApiService        *ElementsConfigsApiService
	projId            string
	envId             string
	elementsEnvUpdate *models.ElementsEnvUpdate
}

func (r ApiUpdateElementsEnvRequest) ElementsEnvUpdate(elementsEnvUpdate models.ElementsEnvUpdate) ApiUpdateElementsEnvRequest {
	r.elementsEnvUpdate = &elementsEnvUpdate
	return r
}

func (r ApiUpdateElementsEnvRequest) Execute() (*models.ElementsEnvRead, *http.Response, error) {
	return r.ApiService.UpdateElementsEnvExecute(r)
}

/*
UpdateElementsEnv Update Elements Env

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @return ApiUpdateElementsEnvRequest
*/
func (a *ElementsConfigsApiService) UpdateElementsEnv(ctx context.Context, projId string, envId string) ApiUpdateElementsEnvRequest {
	return ApiUpdateElementsEnvRequest{
		ApiService: a,
		ctx:        ctx,
		projId:     projId,
		envId:      envId,
	}
}

// Execute executes the request
//  @return ElementsEnvRead
func (a *ElementsConfigsApiService) UpdateElementsEnvExecute(r ApiUpdateElementsEnvRequest) (*models.ElementsEnvRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ElementsEnvRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsConfigsApiService.UpdateElementsEnv")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.elementsEnvUpdate == nil {
		return localVarReturnValue, nil, reportError("elementsEnvUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.elementsEnvUpdate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v models.HTTPValidationError
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
