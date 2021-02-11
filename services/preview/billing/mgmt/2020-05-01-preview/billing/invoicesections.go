package billing

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// InvoiceSectionsClient is the billing client provides access to billing resources for Azure subscriptions.
type InvoiceSectionsClient struct {
	BaseClient
}

// NewInvoiceSectionsClient creates an instance of the InvoiceSectionsClient client.
func NewInvoiceSectionsClient(subscriptionID string) InvoiceSectionsClient {
	return NewInvoiceSectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInvoiceSectionsClientWithBaseURI creates an instance of the InvoiceSectionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewInvoiceSectionsClientWithBaseURI(baseURI string, subscriptionID string) InvoiceSectionsClient {
	return InvoiceSectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an invoice section. The operation is supported only for billing accounts with
// agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
// parameters - the new or updated invoice section.
func (client InvoiceSectionsClient) CreateOrUpdate(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, parameters InvoiceSection) (result InvoiceSectionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client InvoiceSectionsClient) CreateOrUpdatePreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, parameters InvoiceSection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client InvoiceSectionsClient) CreateOrUpdateSender(req *http.Request) (future InvoiceSectionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client InvoiceSectionsClient) (is InvoiceSection, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("billing.InvoiceSectionsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		is.Response.Response, err = future.GetResult(sender)
		if is.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && is.Response.Response.StatusCode != http.StatusNoContent {
			is, err = client.CreateOrUpdateResponder(is.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsCreateOrUpdateFuture", "Result", is.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client InvoiceSectionsClient) CreateOrUpdateResponder(resp *http.Response) (result InvoiceSection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets an invoice section by its ID. The operation is supported only for billing accounts with agreement type
// Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
func (client InvoiceSectionsClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result InvoiceSection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client InvoiceSectionsClient) GetPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InvoiceSectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InvoiceSectionsClient) GetResponder(resp *http.Response) (result InvoiceSection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingProfile lists the invoice sections that a user has access to. The operation is supported only for
// billing accounts with agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
func (client InvoiceSectionsClient) ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result InvoiceSectionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.islr.Response.Response != nil {
				sc = result.islr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingProfileNextResults
	req, err := client.ListByBillingProfilePreparer(ctx, billingAccountName, billingProfileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ListByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.islr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ListByBillingProfile", resp, "Failure sending request")
		return
	}

	result.islr, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ListByBillingProfile", resp, "Failure responding to request")
		return
	}
	if result.islr.hasNextLink() && result.islr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByBillingProfilePreparer prepares the ListByBillingProfile request.
func (client InvoiceSectionsClient) ListByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingProfileSender sends the ListByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client InvoiceSectionsClient) ListByBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingProfileResponder handles the response to the ListByBillingProfile request. The method always
// closes the http.Response Body.
func (client InvoiceSectionsClient) ListByBillingProfileResponder(resp *http.Response) (result InvoiceSectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingProfileNextResults retrieves the next set of results, if any.
func (client InvoiceSectionsClient) listByBillingProfileNextResults(ctx context.Context, lastResults InvoiceSectionListResult) (result InvoiceSectionListResult, err error) {
	req, err := lastResults.invoiceSectionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "listByBillingProfileNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "listByBillingProfileNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "listByBillingProfileNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingProfileComplete enumerates all values, automatically crossing page boundaries as required.
func (client InvoiceSectionsClient) ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string) (result InvoiceSectionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingProfile(ctx, billingAccountName, billingProfileName)
	return
}
