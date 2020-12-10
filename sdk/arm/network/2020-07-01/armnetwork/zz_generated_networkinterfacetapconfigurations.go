// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// NetworkInterfaceTapConfigurationsClient contains the methods for the NetworkInterfaceTapConfigurations group.
// Don't use this type directly, use NewNetworkInterfaceTapConfigurationsClient() instead.
type NetworkInterfaceTapConfigurationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNetworkInterfaceTapConfigurationsClient creates a new instance of NetworkInterfaceTapConfigurationsClient with the specified values.
func NewNetworkInterfaceTapConfigurationsClient(con *armcore.Connection, subscriptionID string) NetworkInterfaceTapConfigurationsClient {
	return NetworkInterfaceTapConfigurationsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client NetworkInterfaceTapConfigurationsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates or updates a Tap configuration in the specified NetworkInterface.
func (client NetworkInterfaceTapConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters NetworkInterfaceTapConfiguration, options *NetworkInterfaceTapConfigurationsBeginCreateOrUpdateOptions) (NetworkInterfaceTapConfigurationPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, tapConfigurationParameters, options)
	if err != nil {
		return NetworkInterfaceTapConfigurationPollerResponse{}, err
	}
	result := NetworkInterfaceTapConfigurationPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkInterfaceTapConfigurationsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return NetworkInterfaceTapConfigurationPollerResponse{}, err
	}
	poller := &networkInterfaceTapConfigurationPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (NetworkInterfaceTapConfigurationResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new NetworkInterfaceTapConfigurationPoller from the specified resume token.
// token - The value must come from a previous call to NetworkInterfaceTapConfigurationPoller.ResumeToken().
func (client NetworkInterfaceTapConfigurationsClient) ResumeCreateOrUpdate(token string) (NetworkInterfaceTapConfigurationPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkInterfaceTapConfigurationsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &networkInterfaceTapConfigurationPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a Tap configuration in the specified NetworkInterface.
func (client NetworkInterfaceTapConfigurationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters NetworkInterfaceTapConfiguration, options *NetworkInterfaceTapConfigurationsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, tapConfigurationParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client NetworkInterfaceTapConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters NetworkInterfaceTapConfiguration, options *NetworkInterfaceTapConfigurationsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{tapConfigurationName}", url.PathEscape(tapConfigurationName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(tapConfigurationParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client NetworkInterfaceTapConfigurationsClient) createOrUpdateHandleResponse(resp *azcore.Response) (NetworkInterfaceTapConfigurationResponse, error) {
	result := NetworkInterfaceTapConfigurationResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.NetworkInterfaceTapConfiguration)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client NetworkInterfaceTapConfigurationsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified tap configuration from the NetworkInterface.
func (client NetworkInterfaceTapConfigurationsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *NetworkInterfaceTapConfigurationsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkInterfaceTapConfigurationsClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client NetworkInterfaceTapConfigurationsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkInterfaceTapConfigurationsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified tap configuration from the NetworkInterface.
func (client NetworkInterfaceTapConfigurationsClient) delete(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *NetworkInterfaceTapConfigurationsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client NetworkInterfaceTapConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *NetworkInterfaceTapConfigurationsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{tapConfigurationName}", url.PathEscape(tapConfigurationName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client NetworkInterfaceTapConfigurationsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Get the specified tap configuration on a network interface.
func (client NetworkInterfaceTapConfigurationsClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *NetworkInterfaceTapConfigurationsGetOptions) (NetworkInterfaceTapConfigurationResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkInterfaceName, tapConfigurationName, options)
	if err != nil {
		return NetworkInterfaceTapConfigurationResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return NetworkInterfaceTapConfigurationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return NetworkInterfaceTapConfigurationResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return NetworkInterfaceTapConfigurationResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client NetworkInterfaceTapConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *NetworkInterfaceTapConfigurationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{tapConfigurationName}", url.PathEscape(tapConfigurationName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client NetworkInterfaceTapConfigurationsClient) getHandleResponse(resp *azcore.Response) (NetworkInterfaceTapConfigurationResponse, error) {
	result := NetworkInterfaceTapConfigurationResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.NetworkInterfaceTapConfiguration)
	return result, err
}

// getHandleError handles the Get error response.
func (client NetworkInterfaceTapConfigurationsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Get all Tap configurations in a network interface.
func (client NetworkInterfaceTapConfigurationsClient) List(resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceTapConfigurationsListOptions) NetworkInterfaceTapConfigurationListResultPager {
	return &networkInterfaceTapConfigurationListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkInterfaceName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp NetworkInterfaceTapConfigurationListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkInterfaceTapConfigurationListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client NetworkInterfaceTapConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceTapConfigurationsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client NetworkInterfaceTapConfigurationsClient) listHandleResponse(resp *azcore.Response) (NetworkInterfaceTapConfigurationListResultResponse, error) {
	result := NetworkInterfaceTapConfigurationListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.NetworkInterfaceTapConfigurationListResult)
	return result, err
}

// listHandleError handles the List error response.
func (client NetworkInterfaceTapConfigurationsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}