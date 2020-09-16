// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// QueueOperations contains the methods for the Queue group.
type QueueOperations interface {
	// Create - Creates a new queue with the specified queue name, under the specified account.
	Create(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue StorageQueue) (*StorageQueueResponse, error)
	// Delete - Deletes the queue with the specified queue name, under the specified account if it exists.
	Delete(ctx context.Context, resourceGroupName string, accountName string, queueName string) (*http.Response, error)
	// Get - Gets the queue with the specified queue name, under the specified account if it exists.
	Get(ctx context.Context, resourceGroupName string, accountName string, queueName string) (*StorageQueueResponse, error)
	// List - Gets a list of all the queues under the specified storage account
	List(resourceGroupName string, accountName string, queueListOptions *QueueListOptions) ListQueueResourcePager
	// Update - Creates a new queue with the specified queue name, under the specified account.
	Update(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue StorageQueue) (*StorageQueueResponse, error)
}

// QueueClient implements the QueueOperations interface.
// Don't use this type directly, use NewQueueClient() instead.
type QueueClient struct {
	*Client
	subscriptionID string
}

// NewQueueClient creates a new instance of QueueClient with the specified values.
func NewQueueClient(c *Client, subscriptionID string) QueueOperations {
	return &QueueClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *QueueClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Create - Creates a new queue with the specified queue name, under the specified account.
func (client *QueueClient) Create(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue StorageQueue) (*StorageQueueResponse, error) {
	req, err := client.CreateCreateRequest(ctx, resourceGroupName, accountName, queueName, queue)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateHandleError(resp)
	}
	result, err := client.CreateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCreateRequest creates the Create request.
func (client *QueueClient) CreateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue StorageQueue) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(queue)
}

// CreateHandleResponse handles the Create response.
func (client *QueueClient) CreateHandleResponse(resp *azcore.Response) (*StorageQueueResponse, error) {
	result := StorageQueueResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.StorageQueue)
}

// CreateHandleError handles the Create error response.
func (client *QueueClient) CreateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the queue with the specified queue name, under the specified account if it exists.
func (client *QueueClient) Delete(ctx context.Context, resourceGroupName string, accountName string, queueName string) (*http.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, accountName, queueName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp.Response, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *QueueClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, queueName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *QueueClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the queue with the specified queue name, under the specified account if it exists.
func (client *QueueClient) Get(ctx context.Context, resourceGroupName string, accountName string, queueName string) (*StorageQueueResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, accountName, queueName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *QueueClient) GetCreateRequest(ctx context.Context, resourceGroupName string, accountName string, queueName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *QueueClient) GetHandleResponse(resp *azcore.Response) (*StorageQueueResponse, error) {
	result := StorageQueueResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.StorageQueue)
}

// GetHandleError handles the Get error response.
func (client *QueueClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets a list of all the queues under the specified storage account
func (client *QueueClient) List(resourceGroupName string, accountName string, queueListOptions *QueueListOptions) ListQueueResourcePager {
	return &listQueueResourcePager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, accountName, queueListOptions)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ListQueueResourceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListQueueResource.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *QueueClient) ListCreateRequest(ctx context.Context, resourceGroupName string, accountName string, queueListOptions *QueueListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	if queueListOptions != nil && queueListOptions.Maxpagesize != nil {
		query.Set("$maxpagesize", *queueListOptions.Maxpagesize)
	}
	if queueListOptions != nil && queueListOptions.Filter != nil {
		query.Set("$filter", *queueListOptions.Filter)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *QueueClient) ListHandleResponse(resp *azcore.Response) (*ListQueueResourceResponse, error) {
	result := ListQueueResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListQueueResource)
}

// ListHandleError handles the List error response.
func (client *QueueClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Update - Creates a new queue with the specified queue name, under the specified account.
func (client *QueueClient) Update(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue StorageQueue) (*StorageQueueResponse, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, accountName, queueName, queue)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateHandleError(resp)
	}
	result, err := client.UpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateCreateRequest creates the Update request.
func (client *QueueClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, queueName string, queue StorageQueue) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(queue)
}

// UpdateHandleResponse handles the Update response.
func (client *QueueClient) UpdateHandleResponse(resp *azcore.Response) (*StorageQueueResponse, error) {
	result := StorageQueueResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.StorageQueue)
}

// UpdateHandleError handles the Update error response.
func (client *QueueClient) UpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}