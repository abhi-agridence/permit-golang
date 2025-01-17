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

// ElementsDataApiService ElementsDataApi service
type ElementsDataApiService service

type ApiElementsAssignRoleToUserRequest struct {
	ctx                    context.Context
	ApiService             *ElementsDataApiService
	projId                 string
	envId                  string
	elementsConfigId       string
	userId                 string
	elementsUserRoleCreate *models.ElementsUserRoleCreate
}

func (r ApiElementsAssignRoleToUserRequest) ElementsUserRoleCreate(elementsUserRoleCreate models.ElementsUserRoleCreate) ApiElementsAssignRoleToUserRequest {
	r.elementsUserRoleCreate = &elementsUserRoleCreate
	return r
}

func (r ApiElementsAssignRoleToUserRequest) Execute() (*models.RoleAssignmentRead, *http.Response, error) {
	return r.ApiService.ElementsAssignRoleToUserExecute(r)
}

/*
ElementsAssignRoleToUser Assign role to user

Assigns a role to the user within the tenant.

The tenant defines the scope of the assignment. In other words, the role is effective only within the tenant.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @param elementsConfigId Either the unique id of the elements_config, or the URL-friendly key of the elements_config (i.e: the \"slug\").
 @param userId Either the unique id of the user, or the URL-friendly key of the user (i.e: the \"slug\").
 @return ApiElementsAssignRoleToUserRequest
*/
func (a *ElementsDataApiService) ElementsAssignRoleToUser(ctx context.Context, projId string, envId string, elementsConfigId string, userId string) ApiElementsAssignRoleToUserRequest {
	return ApiElementsAssignRoleToUserRequest{
		ApiService:       a,
		ctx:              ctx,
		projId:           projId,
		envId:            envId,
		elementsConfigId: elementsConfigId,
		userId:           userId,
	}
}

// Execute executes the request
//  @return RoleAssignmentRead
func (a *ElementsDataApiService) ElementsAssignRoleToUserExecute(r ApiElementsAssignRoleToUserRequest) (*models.RoleAssignmentRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.RoleAssignmentRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsDataApiService.ElementsAssignRoleToUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}/config/{elements_config_id}/data/users/{user_id}/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"elements_config_id"+"}", url.PathEscape(parameterToString(r.elementsConfigId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.elementsUserRoleCreate == nil {
		return localVarReturnValue, nil, reportError("elementsUserRoleCreate is required and must be specified")
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
	localVarPostBody = r.elementsUserRoleCreate
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

type ApiElementsCreateUserRequest struct {
	ctx                context.Context
	ApiService         *ElementsDataApiService
	projId             string
	envId              string
	elementsConfigId   string
	elementsUserCreate *models.ElementsUserCreate
}

func (r ApiElementsCreateUserRequest) ElementsUserCreate(elementsUserCreate models.ElementsUserCreate) ApiElementsCreateUserRequest {
	r.elementsUserCreate = &elementsUserCreate
	return r
}

func (r ApiElementsCreateUserRequest) Execute() (*models.UserRead, *http.Response, error) {
	return r.ApiService.ElementsCreateUserExecute(r)
}

/*
ElementsCreateUser Create user

Creates a new user inside the Permit.io system, from that point forward
you may run permission checks on that user.

If the user is already created: will return 200 instead of 201,
and will return the existing user object in the response body.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @param elementsConfigId Either the unique id of the elements_config, or the URL-friendly key of the elements_config (i.e: the \"slug\").
 @return ApiElementsCreateUserRequest
*/
func (a *ElementsDataApiService) ElementsCreateUser(ctx context.Context, projId string, envId string, elementsConfigId string) ApiElementsCreateUserRequest {
	return ApiElementsCreateUserRequest{
		ApiService:       a,
		ctx:              ctx,
		projId:           projId,
		envId:            envId,
		elementsConfigId: elementsConfigId,
	}
}

// Execute executes the request
//  @return UserRead
func (a *ElementsDataApiService) ElementsCreateUserExecute(r ApiElementsCreateUserRequest) (*models.UserRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.UserRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsDataApiService.ElementsCreateUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}/config/{elements_config_id}/data/users"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"elements_config_id"+"}", url.PathEscape(parameterToString(r.elementsConfigId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.elementsUserCreate == nil {
		return localVarReturnValue, nil, reportError("elementsUserCreate is required and must be specified")
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
	localVarPostBody = r.elementsUserCreate
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

type ApiElementsDeleteUserRequest struct {
	ctx              context.Context
	ApiService       *ElementsDataApiService
	projId           string
	envId            string
	elementsConfigId string
	userId           string
}

func (r ApiElementsDeleteUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.ElementsDeleteUserExecute(r)
}

/*
ElementsDeleteUser Delete user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @param elementsConfigId Either the unique id of the elements_config, or the URL-friendly key of the elements_config (i.e: the \"slug\").
 @param userId Either the unique id of the user, or the URL-friendly key of the user (i.e: the \"slug\").
 @return ApiElementsDeleteUserRequest
*/
func (a *ElementsDataApiService) ElementsDeleteUser(ctx context.Context, projId string, envId string, elementsConfigId string, userId string) ApiElementsDeleteUserRequest {
	return ApiElementsDeleteUserRequest{
		ApiService:       a,
		ctx:              ctx,
		projId:           projId,
		envId:            envId,
		elementsConfigId: elementsConfigId,
		userId:           userId,
	}
}

// Execute executes the request
func (a *ElementsDataApiService) ElementsDeleteUserExecute(r ApiElementsDeleteUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsDataApiService.ElementsDeleteUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}/config/{elements_config_id}/data/users/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"elements_config_id"+"}", url.PathEscape(parameterToString(r.elementsConfigId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

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

type ApiElementsListRolesRequest struct {
	ctx              context.Context
	ApiService       *ElementsDataApiService
	projId           string
	envId            string
	elementsConfigId string
	search           *string
	page             *int32
	perPage          *int32
}

// Text search for the email field
func (r ApiElementsListRolesRequest) Search(search string) ApiElementsListRolesRequest {
	r.search = &search
	return r
}

// Page number of the results to fetch, starting at 1.
func (r ApiElementsListRolesRequest) Page(page int32) ApiElementsListRolesRequest {
	r.page = &page
	return r
}

// The number of results per page (max 100).
func (r ApiElementsListRolesRequest) PerPage(perPage int32) ApiElementsListRolesRequest {
	r.perPage = &perPage
	return r
}

func (r ApiElementsListRolesRequest) Execute() ([]models.ElementsRoleRead, *http.Response, error) {
	return r.ApiService.ElementsListRolesExecute(r)
}

/*
ElementsListRoles List roles

Lists all the users defined within an environment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @param elementsConfigId Either the unique id of the elements_config, or the URL-friendly key of the elements_config (i.e: the \"slug\").
 @return ApiElementsListRolesRequest
*/
func (a *ElementsDataApiService) ElementsListRoles(ctx context.Context, projId string, envId string, elementsConfigId string) ApiElementsListRolesRequest {
	return ApiElementsListRolesRequest{
		ApiService:       a,
		ctx:              ctx,
		projId:           projId,
		envId:            envId,
		elementsConfigId: elementsConfigId,
	}
}

// Execute executes the request
//  @return []ElementsRoleRead
func (a *ElementsDataApiService) ElementsListRolesExecute(r ApiElementsListRolesRequest) ([]models.ElementsRoleRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []models.ElementsRoleRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsDataApiService.ElementsListRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}/config/{elements_config_id}/data/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"elements_config_id"+"}", url.PathEscape(parameterToString(r.elementsConfigId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
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

type ApiElementsListUsersRequest struct {
	ctx              context.Context
	ApiService       *ElementsDataApiService
	projId           string
	envId            string
	elementsConfigId string
	search           *string
	page             *int32
	perPage          *int32
}

// Text search for the email field
func (r ApiElementsListUsersRequest) Search(search string) ApiElementsListUsersRequest {
	r.search = &search
	return r
}

// Page number of the results to fetch, starting at 1.
func (r ApiElementsListUsersRequest) Page(page int32) ApiElementsListUsersRequest {
	r.page = &page
	return r
}

// The number of results per page (max 100).
func (r ApiElementsListUsersRequest) PerPage(perPage int32) ApiElementsListUsersRequest {
	r.perPage = &perPage
	return r
}

func (r ApiElementsListUsersRequest) Execute() (*models.PaginatedResultUserRead, *http.Response, error) {
	return r.ApiService.ElementsListUsersExecute(r)
}

/*
ElementsListUsers List users

Lists all the users defined within an environment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @param elementsConfigId Either the unique id of the elements_config, or the URL-friendly key of the elements_config (i.e: the \"slug\").
 @return ApiElementsListUsersRequest
*/
func (a *ElementsDataApiService) ElementsListUsers(ctx context.Context, projId string, envId string, elementsConfigId string) ApiElementsListUsersRequest {
	return ApiElementsListUsersRequest{
		ApiService:       a,
		ctx:              ctx,
		projId:           projId,
		envId:            envId,
		elementsConfigId: elementsConfigId,
	}
}

// Execute executes the request
//  @return PaginatedResultUserRead
func (a *ElementsDataApiService) ElementsListUsersExecute(r ApiElementsListUsersRequest) (*models.PaginatedResultUserRead, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.PaginatedResultUserRead
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsDataApiService.ElementsListUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}/config/{elements_config_id}/data/users"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"elements_config_id"+"}", url.PathEscape(parameterToString(r.elementsConfigId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
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

type ApiElementsUnassignRoleFromUserRequest struct {
	ctx                    context.Context
	ApiService             *ElementsDataApiService
	projId                 string
	envId                  string
	elementsConfigId       string
	userId                 string
	elementsUserRoleRemove *models.ElementsUserRoleRemove
}

func (r ApiElementsUnassignRoleFromUserRequest) ElementsUserRoleRemove(elementsUserRoleRemove models.ElementsUserRoleRemove) ApiElementsUnassignRoleFromUserRequest {
	r.elementsUserRoleRemove = &elementsUserRoleRemove
	return r
}

func (r ApiElementsUnassignRoleFromUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.ElementsUnassignRoleFromUserExecute(r)
}

/*
ElementsUnassignRoleFromUser Unassign role from user

Unassigns the role from the user within the tenant.

The tenant defines the scope of the assignment. In other words, the role is effective only within the tenant.

If the role is not actually assigned, will return 404.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @param elementsConfigId Either the unique id of the elements_config, or the URL-friendly key of the elements_config (i.e: the \"slug\").
 @param userId Either the unique id of the user, or the URL-friendly key of the user (i.e: the \"slug\").
 @return ApiElementsUnassignRoleFromUserRequest
*/
func (a *ElementsDataApiService) ElementsUnassignRoleFromUser(ctx context.Context, projId string, envId string, elementsConfigId string, userId string) ApiElementsUnassignRoleFromUserRequest {
	return ApiElementsUnassignRoleFromUserRequest{
		ApiService:       a,
		ctx:              ctx,
		projId:           projId,
		envId:            envId,
		elementsConfigId: elementsConfigId,
		userId:           userId,
	}
}

// Execute executes the request
func (a *ElementsDataApiService) ElementsUnassignRoleFromUserExecute(r ApiElementsUnassignRoleFromUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsDataApiService.ElementsUnassignRoleFromUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}/config/{elements_config_id}/data/users/{user_id}/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"proj_id"+"}", url.PathEscape(parameterToString(r.projId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"env_id"+"}", url.PathEscape(parameterToString(r.envId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"elements_config_id"+"}", url.PathEscape(parameterToString(r.elementsConfigId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.elementsUserRoleRemove == nil {
		return nil, reportError("elementsUserRoleRemove is required and must be specified")
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
	localVarPostBody = r.elementsUserRoleRemove
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

type ApiSetConfigActiveRequest struct {
	ctx              context.Context
	ApiService       *ElementsDataApiService
	elementsConfigId string
	projId           string
	envId            string
}

func (r ApiSetConfigActiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.SetConfigActiveExecute(r)
}

/*
SetConfigActive Set Config Active

Updates the embed_config.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param elementsConfigId Either the unique id of the elements_config, or the URL-friendly key of the elements_config (i.e: the \"slug\").
 @param projId Either the unique id of the project, or the URL-friendly key of the project (i.e: the \"slug\").
 @param envId Either the unique id of the environment, or the URL-friendly key of the environment (i.e: the \"slug\").
 @return ApiSetConfigActiveRequest
*/
func (a *ElementsDataApiService) SetConfigActive(ctx context.Context, elementsConfigId string, projId string, envId string) ApiSetConfigActiveRequest {
	return ApiSetConfigActiveRequest{
		ApiService:       a,
		ctx:              ctx,
		elementsConfigId: elementsConfigId,
		projId:           projId,
		envId:            envId,
	}
}

// Execute executes the request
func (a *ElementsDataApiService) SetConfigActiveExecute(r ApiSetConfigActiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ElementsDataApiService.SetConfigActive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/elements/{proj_id}/{env_id}/config/{elements_config_id}/data/active"
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
