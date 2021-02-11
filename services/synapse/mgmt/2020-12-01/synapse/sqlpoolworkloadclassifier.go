package synapse

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

// SQLPoolWorkloadClassifierClient is the azure Synapse Analytics Management Client
type SQLPoolWorkloadClassifierClient struct {
	BaseClient
}

// NewSQLPoolWorkloadClassifierClient creates an instance of the SQLPoolWorkloadClassifierClient client.
func NewSQLPoolWorkloadClassifierClient(subscriptionID string) SQLPoolWorkloadClassifierClient {
	return NewSQLPoolWorkloadClassifierClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSQLPoolWorkloadClassifierClientWithBaseURI creates an instance of the SQLPoolWorkloadClassifierClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewSQLPoolWorkloadClassifierClientWithBaseURI(baseURI string, subscriptionID string) SQLPoolWorkloadClassifierClient {
	return SQLPoolWorkloadClassifierClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create Or Update workload classifier for a Sql pool's workload group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// workloadGroupName - the name of the workload group.
// workloadClassifierName - the name of the workload classifier.
// parameters - the properties of the workload classifier.
func (client SQLPoolWorkloadClassifierClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier) (result SQLPoolWorkloadClassifierCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolWorkloadClassifierClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.WorkloadClassifierProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.WorkloadClassifierProperties.MemberName", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("synapse.SQLPoolWorkloadClassifierClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, workloadGroupName, workloadClassifierName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SQLPoolWorkloadClassifierClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workloadClassifierName": autorest.Encode("path", workloadClassifierName),
		"workloadGroupName":      autorest.Encode("path", workloadGroupName),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolWorkloadClassifierClient) CreateOrUpdateSender(req *http.Request) (future SQLPoolWorkloadClassifierCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client SQLPoolWorkloadClassifierClient) (wc WorkloadClassifier, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("synapse.SQLPoolWorkloadClassifierCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		wc.Response.Response, err = future.GetResult(sender)
		if wc.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && wc.Response.Response.StatusCode != http.StatusNoContent {
			wc, err = client.CreateOrUpdateResponder(wc.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierCreateOrUpdateFuture", "Result", wc.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SQLPoolWorkloadClassifierClient) CreateOrUpdateResponder(resp *http.Response) (result WorkloadClassifier, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete remove workload classifier of a Sql pool's workload group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// workloadGroupName - the name of the workload group.
// workloadClassifierName - the name of the workload classifier.
func (client SQLPoolWorkloadClassifierClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string, workloadClassifierName string) (result SQLPoolWorkloadClassifierDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolWorkloadClassifierClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.SQLPoolWorkloadClassifierClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, workloadGroupName, workloadClassifierName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SQLPoolWorkloadClassifierClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string, workloadClassifierName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workloadClassifierName": autorest.Encode("path", workloadClassifierName),
		"workloadGroupName":      autorest.Encode("path", workloadGroupName),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolWorkloadClassifierClient) DeleteSender(req *http.Request) (future SQLPoolWorkloadClassifierDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client SQLPoolWorkloadClassifierClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("synapse.SQLPoolWorkloadClassifierDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SQLPoolWorkloadClassifierClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a workload classifier of Sql pool's workload group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// workloadGroupName - the name of the workload group.
// workloadClassifierName - the name of the workload classifier.
func (client SQLPoolWorkloadClassifierClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string, workloadClassifierName string) (result WorkloadClassifier, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolWorkloadClassifierClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.SQLPoolWorkloadClassifierClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, workloadGroupName, workloadClassifierName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SQLPoolWorkloadClassifierClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string, workloadClassifierName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workloadClassifierName": autorest.Encode("path", workloadClassifierName),
		"workloadGroupName":      autorest.Encode("path", workloadGroupName),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolWorkloadClassifierClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SQLPoolWorkloadClassifierClient) GetResponder(resp *http.Response) (result WorkloadClassifier, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get list of  Sql pool's workload classifier for workload groups.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// workloadGroupName - the name of the workload group.
func (client SQLPoolWorkloadClassifierClient) List(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string) (result WorkloadClassifierListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolWorkloadClassifierClient.List")
		defer func() {
			sc := -1
			if result.wclr.Response.Response != nil {
				sc = result.wclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.SQLPoolWorkloadClassifierClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, workloadGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.wclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierClient", "List", resp, "Failure sending request")
		return
	}

	result.wclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierClient", "List", resp, "Failure responding to request")
		return
	}
	if result.wclr.hasNextLink() && result.wclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client SQLPoolWorkloadClassifierClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sqlPoolName":       autorest.Encode("path", SQLPoolName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workloadGroupName": autorest.Encode("path", workloadGroupName),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}/workloadClassifiers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolWorkloadClassifierClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SQLPoolWorkloadClassifierClient) ListResponder(resp *http.Response) (result WorkloadClassifierListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SQLPoolWorkloadClassifierClient) listNextResults(ctx context.Context, lastResults WorkloadClassifierListResult) (result WorkloadClassifierListResult, err error) {
	req, err := lastResults.workloadClassifierListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifierClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLPoolWorkloadClassifierClient) ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string) (result WorkloadClassifierListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolWorkloadClassifierClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, workspaceName, SQLPoolName, workloadGroupName)
	return
}
