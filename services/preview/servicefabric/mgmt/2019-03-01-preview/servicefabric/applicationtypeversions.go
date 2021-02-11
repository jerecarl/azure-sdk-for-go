package servicefabric

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ApplicationTypeVersionsClient is the service Fabric Management Client
type ApplicationTypeVersionsClient struct {
	BaseClient
}

// NewApplicationTypeVersionsClient creates an instance of the ApplicationTypeVersionsClient client.
func NewApplicationTypeVersionsClient(subscriptionID string) ApplicationTypeVersionsClient {
	return NewApplicationTypeVersionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationTypeVersionsClientWithBaseURI creates an instance of the ApplicationTypeVersionsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewApplicationTypeVersionsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationTypeVersionsClient {
	return ApplicationTypeVersionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create or update a Service Fabric application type version resource with the specified name.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster resource.
// applicationTypeName - the name of the application type name resource.
// version - the application type version.
// parameters - the application type version resource.
func (client ApplicationTypeVersionsClient) Create(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, parameters ApplicationTypeVersionResource) (result ApplicationTypeVersionsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationTypeVersionsClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ApplicationTypeVersionResourceProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ApplicationTypeVersionResourceProperties.AppPackageURL", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("servicefabric.ApplicationTypeVersionsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, clusterName, applicationTypeName, version, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ApplicationTypeVersionsClient) CreatePreparer(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, parameters ApplicationTypeVersionResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationTypeName": autorest.Encode("path", applicationTypeName),
		"clusterName":         autorest.Encode("path", clusterName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"version":             autorest.Encode("path", version),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationTypeVersionsClient) CreateSender(req *http.Request) (future ApplicationTypeVersionsCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ApplicationTypeVersionsClient) (atvr ApplicationTypeVersionResource, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("servicefabric.ApplicationTypeVersionsCreateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		atvr.Response.Response, err = future.GetResult(sender)
		if atvr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsCreateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && atvr.Response.Response.StatusCode != http.StatusNoContent {
			atvr, err = client.CreateResponder(atvr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsCreateFuture", "Result", atvr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ApplicationTypeVersionsClient) CreateResponder(resp *http.Response) (result ApplicationTypeVersionResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a Service Fabric application type version resource with the specified name.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster resource.
// applicationTypeName - the name of the application type name resource.
// version - the application type version.
func (client ApplicationTypeVersionsClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string) (result ApplicationTypeVersionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationTypeVersionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, applicationTypeName, version)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationTypeVersionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationTypeName": autorest.Encode("path", applicationTypeName),
		"clusterName":         autorest.Encode("path", clusterName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"version":             autorest.Encode("path", version),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationTypeVersionsClient) DeleteSender(req *http.Request) (future ApplicationTypeVersionsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ApplicationTypeVersionsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("servicefabric.ApplicationTypeVersionsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationTypeVersionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a Service Fabric application type version resource created or in the process of being created in the Service
// Fabric application type name resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster resource.
// applicationTypeName - the name of the application type name resource.
// version - the application type version.
func (client ApplicationTypeVersionsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string) (result ApplicationTypeVersionResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationTypeVersionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, applicationTypeName, version)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationTypeVersionsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationTypeName": autorest.Encode("path", applicationTypeName),
		"clusterName":         autorest.Encode("path", clusterName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"version":             autorest.Encode("path", version),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationTypeVersionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationTypeVersionsClient) GetResponder(resp *http.Response) (result ApplicationTypeVersionResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all application type version resources created or in the process of being created in the Service Fabric
// application type name resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster resource.
// applicationTypeName - the name of the application type name resource.
func (client ApplicationTypeVersionsClient) List(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string) (result ApplicationTypeVersionResourceList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationTypeVersionsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, clusterName, applicationTypeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationTypeVersionsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationTypeVersionsClient) ListPreparer(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationTypeName": autorest.Encode("path", applicationTypeName),
		"clusterName":         autorest.Encode("path", clusterName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationTypeVersionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationTypeVersionsClient) ListResponder(resp *http.Response) (result ApplicationTypeVersionResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
