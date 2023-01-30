package containerregistry

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
)

// TagClient is the metadata API definition for the Azure Container Registry runtime
type TagClient struct {
	BaseClient
}

// NewTagClient creates an instance of the TagClient client.
func NewTagClient(loginURI string) TagClient {
	return TagClient{New(loginURI)}
}

// Delete delete tag
// Parameters:
// name - name of the image (including the namespace)
// reference - tag name
func (client TagClient) Delete(ctx context.Context, name string, reference string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, name, reference)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.TagClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "containerregistry.TagClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.TagClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client TagClient) DeletePreparer(ctx context.Context, name string, reference string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"url": client.LoginURI,
	}

	pathParameters := map[string]interface{}{
		"name":      autorest.Encode("path", name),
		"reference": autorest.Encode("path", reference),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPathParameters("/acr/v1/{name}/_tags/{reference}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client TagClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client TagClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetAttributes get tag attributes by tag
// Parameters:
// name - name of the image (including the namespace)
// reference - tag name
func (client TagClient) GetAttributes(ctx context.Context, name string, reference string) (result TagAttributes, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagClient.GetAttributes")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAttributesPreparer(ctx, name, reference)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.TagClient", "GetAttributes", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAttributesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.TagClient", "GetAttributes", resp, "Failure sending request")
		return
	}

	result, err = client.GetAttributesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.TagClient", "GetAttributes", resp, "Failure responding to request")
		return
	}

	return
}

// GetAttributesPreparer prepares the GetAttributes request.
func (client TagClient) GetAttributesPreparer(ctx context.Context, name string, reference string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"url": client.LoginURI,
	}

	pathParameters := map[string]interface{}{
		"name":      autorest.Encode("path", name),
		"reference": autorest.Encode("path", reference),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPathParameters("/acr/v1/{name}/_tags/{reference}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAttributesSender sends the GetAttributes request. The method will close the
// http.Response Body if it receives an error.
func (client TagClient) GetAttributesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAttributesResponder handles the response to the GetAttributes request. The method always
// closes the http.Response Body.
func (client TagClient) GetAttributesResponder(resp *http.Response) (result TagAttributes, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList list tags of a repository
// Parameters:
// name - name of the image (including the namespace)
// last - query parameter for the last item in previous query. Result set will include values lexically after
// last.
// n - query parameter for max number of items
// orderby - orderby query parameter
// digest - filter by digest
func (client TagClient) GetList(ctx context.Context, name string, last string, n *int32, orderby string, digest string) (result TagList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, name, last, n, orderby, digest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.TagClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.TagClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.TagClient", "GetList", resp, "Failure responding to request")
		return
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client TagClient) GetListPreparer(ctx context.Context, name string, last string, n *int32, orderby string, digest string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"url": client.LoginURI,
	}

	pathParameters := map[string]interface{}{
		"name": autorest.Encode("path", name),
	}

	queryParameters := map[string]interface{}{}
	if len(last) > 0 {
		queryParameters["last"] = autorest.Encode("query", last)
	}
	if n != nil {
		queryParameters["n"] = autorest.Encode("query", *n)
	}
	if len(orderby) > 0 {
		queryParameters["orderby"] = autorest.Encode("query", orderby)
	}
	if len(digest) > 0 {
		queryParameters["digest"] = autorest.Encode("query", digest)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPathParameters("/acr/v1/{name}/_tags", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client TagClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client TagClient) GetListResponder(resp *http.Response) (result TagList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateAttributes update tag attributes
// Parameters:
// name - name of the image (including the namespace)
// reference - tag name
// value - repository attribute value
func (client TagClient) UpdateAttributes(ctx context.Context, name string, reference string, value *ChangeableAttributes) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagClient.UpdateAttributes")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateAttributesPreparer(ctx, name, reference, value)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.TagClient", "UpdateAttributes", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateAttributesSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "containerregistry.TagClient", "UpdateAttributes", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateAttributesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.TagClient", "UpdateAttributes", resp, "Failure responding to request")
		return
	}

	return
}

// UpdateAttributesPreparer prepares the UpdateAttributes request.
func (client TagClient) UpdateAttributesPreparer(ctx context.Context, name string, reference string, value *ChangeableAttributes) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"url": client.LoginURI,
	}

	pathParameters := map[string]interface{}{
		"name":      autorest.Encode("path", name),
		"reference": autorest.Encode("path", reference),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPathParameters("/acr/v1/{name}/_tags/{reference}", pathParameters))
	if value != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(value))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateAttributesSender sends the UpdateAttributes request. The method will close the
// http.Response Body if it receives an error.
func (client TagClient) UpdateAttributesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateAttributesResponder handles the response to the UpdateAttributes request. The method always
// closes the http.Response Body.
func (client TagClient) UpdateAttributesResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
