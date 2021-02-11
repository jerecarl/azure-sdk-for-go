package digitaltwins

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

// EndpointClient is the azure Digital Twins Client for managing DigitalTwinsInstance
type EndpointClient struct {
	BaseClient
}

// NewEndpointClient creates an instance of the EndpointClient client.
func NewEndpointClient(subscriptionID string) EndpointClient {
	return NewEndpointClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEndpointClientWithBaseURI creates an instance of the EndpointClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEndpointClientWithBaseURI(baseURI string, subscriptionID string) EndpointClient {
	return EndpointClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update DigitalTwinsInstance endpoint.
// Parameters:
// resourceGroupName - the name of the resource group that contains the DigitalTwinsInstance.
// resourceName - the name of the DigitalTwinsInstance.
// endpointName - name of Endpoint Resource.
// endpointDescription - the DigitalTwinsInstance endpoint metadata and security metadata.
func (client EndpointClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, endpointDescription EndpointResource) (result EndpointCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: endpointName,
			Constraints: []validation.Constraint{{Target: "endpointName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "endpointName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "endpointName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-._]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("digitaltwins.EndpointClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, resourceName, endpointName, endpointDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.EndpointClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.EndpointClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client EndpointClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, endpointDescription EndpointResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints/{endpointName}", pathParameters),
		autorest.WithJSON(endpointDescription),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointClient) CreateOrUpdateSender(req *http.Request) (future EndpointCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client EndpointClient) (er EndpointResource, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "digitaltwins.EndpointCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("digitaltwins.EndpointCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		er.Response.Response, err = future.GetResult(sender)
		if er.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "digitaltwins.EndpointCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && er.Response.Response.StatusCode != http.StatusNoContent {
			er, err = client.CreateOrUpdateResponder(er.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "digitaltwins.EndpointCreateOrUpdateFuture", "Result", er.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client EndpointClient) CreateOrUpdateResponder(resp *http.Response) (result EndpointResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a DigitalTwinsInstance endpoint.
// Parameters:
// resourceGroupName - the name of the resource group that contains the DigitalTwinsInstance.
// resourceName - the name of the DigitalTwinsInstance.
// endpointName - name of Endpoint Resource.
func (client EndpointClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, endpointName string) (result EndpointDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: endpointName,
			Constraints: []validation.Constraint{{Target: "endpointName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "endpointName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "endpointName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-._]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("digitaltwins.EndpointClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName, endpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.EndpointClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.EndpointClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client EndpointClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string, endpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints/{endpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointClient) DeleteSender(req *http.Request) (future EndpointDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client EndpointClient) (er EndpointResource, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "digitaltwins.EndpointDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("digitaltwins.EndpointDeleteFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		er.Response.Response, err = future.GetResult(sender)
		if er.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "digitaltwins.EndpointDeleteFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && er.Response.Response.StatusCode != http.StatusNoContent {
			er, err = client.DeleteResponder(er.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "digitaltwins.EndpointDeleteFuture", "Result", er.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client EndpointClient) DeleteResponder(resp *http.Response) (result EndpointResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get DigitalTwinsInstances Endpoint.
// Parameters:
// resourceGroupName - the name of the resource group that contains the DigitalTwinsInstance.
// resourceName - the name of the DigitalTwinsInstance.
// endpointName - name of Endpoint Resource.
func (client EndpointClient) Get(ctx context.Context, resourceGroupName string, resourceName string, endpointName string) (result EndpointResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: endpointName,
			Constraints: []validation.Constraint{{Target: "endpointName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "endpointName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "endpointName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-._]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("digitaltwins.EndpointClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName, endpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.EndpointClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "digitaltwins.EndpointClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.EndpointClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client EndpointClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string, endpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints/{endpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EndpointClient) GetResponder(resp *http.Response) (result EndpointResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get DigitalTwinsInstance Endpoints.
// Parameters:
// resourceGroupName - the name of the resource group that contains the DigitalTwinsInstance.
// resourceName - the name of the DigitalTwinsInstance.
func (client EndpointClient) List(ctx context.Context, resourceGroupName string, resourceName string) (result EndpointResourceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointClient.List")
		defer func() {
			sc := -1
			if result.erlr.Response.Response != nil {
				sc = result.erlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("digitaltwins.EndpointClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.EndpointClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.erlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "digitaltwins.EndpointClient", "List", resp, "Failure sending request")
		return
	}

	result.erlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.EndpointClient", "List", resp, "Failure responding to request")
		return
	}
	if result.erlr.hasNextLink() && result.erlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client EndpointClient) ListPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client EndpointClient) ListResponder(resp *http.Response) (result EndpointResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client EndpointClient) listNextResults(ctx context.Context, lastResults EndpointResourceListResult) (result EndpointResourceListResult, err error) {
	req, err := lastResults.endpointResourceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "digitaltwins.EndpointClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "digitaltwins.EndpointClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "digitaltwins.EndpointClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client EndpointClient) ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result EndpointResourceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, resourceName)
	return
}
