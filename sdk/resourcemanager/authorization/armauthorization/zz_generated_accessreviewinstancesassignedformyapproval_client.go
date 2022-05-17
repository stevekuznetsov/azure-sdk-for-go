//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AccessReviewInstancesAssignedForMyApprovalClient contains the methods for the AccessReviewInstancesAssignedForMyApproval group.
// Don't use this type directly, use NewAccessReviewInstancesAssignedForMyApprovalClient() instead.
type AccessReviewInstancesAssignedForMyApprovalClient struct {
	host string
	pl   runtime.Pipeline
}

// NewAccessReviewInstancesAssignedForMyApprovalClient creates a new instance of AccessReviewInstancesAssignedForMyApprovalClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccessReviewInstancesAssignedForMyApprovalClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessReviewInstancesAssignedForMyApprovalClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AccessReviewInstancesAssignedForMyApprovalClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// GetByID - Get single access review instance assigned for my approval.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-16-preview
// scheduleDefinitionID - The id of the access review schedule definition.
// id - The id of the access review instance.
// options - AccessReviewInstancesAssignedForMyApprovalClientGetByIDOptions contains the optional parameters for the AccessReviewInstancesAssignedForMyApprovalClient.GetByID
// method.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) GetByID(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstancesAssignedForMyApprovalClientGetByIDOptions) (AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, scheduleDefinitionID, id, options)
	if err != nil {
		return AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) getByIDCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstancesAssignedForMyApprovalClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) getByIDHandleResponse(resp *http.Response) (AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse, error) {
	result := AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewInstance); err != nil {
		return AccessReviewInstancesAssignedForMyApprovalClientGetByIDResponse{}, err
	}
	return result, nil
}

// NewListPager - Get access review instances assigned for my approval.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-16-preview
// scheduleDefinitionID - The id of the access review schedule definition.
// options - AccessReviewInstancesAssignedForMyApprovalClientListOptions contains the optional parameters for the AccessReviewInstancesAssignedForMyApprovalClient.List
// method.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) NewListPager(scheduleDefinitionID string, options *AccessReviewInstancesAssignedForMyApprovalClientListOptions) *runtime.Pager[AccessReviewInstancesAssignedForMyApprovalClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessReviewInstancesAssignedForMyApprovalClientListResponse]{
		More: func(page AccessReviewInstancesAssignedForMyApprovalClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessReviewInstancesAssignedForMyApprovalClientListResponse) (AccessReviewInstancesAssignedForMyApprovalClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scheduleDefinitionID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccessReviewInstancesAssignedForMyApprovalClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AccessReviewInstancesAssignedForMyApprovalClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccessReviewInstancesAssignedForMyApprovalClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) listCreateRequest(ctx context.Context, scheduleDefinitionID string, options *AccessReviewInstancesAssignedForMyApprovalClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessReviewInstancesAssignedForMyApprovalClient) listHandleResponse(resp *http.Response) (AccessReviewInstancesAssignedForMyApprovalClientListResponse, error) {
	result := AccessReviewInstancesAssignedForMyApprovalClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewInstanceListResult); err != nil {
		return AccessReviewInstancesAssignedForMyApprovalClientListResponse{}, err
	}
	return result, nil
}
