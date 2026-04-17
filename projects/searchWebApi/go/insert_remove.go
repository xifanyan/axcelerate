package searchwebapi

import (
	"context"
	"net/url"
)

func (c *Client) InsertRemoveTransaction(ctx context.Context, projectID, collectionID string, searchCacheControl string, request InsertRemoveRequest) (*InsertRemoveResult, error) {
	var result InsertRemoveResult
	query := url.Values{}
	addQuery(query, "SWA-searchCacheControl", searchCacheControl)
	if err := c.doJSON(ctx, httpMethodPost, []string{"projects", projectID, "collections", collectionID, "records", "insertRemoveTransaction"}, query, nil, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) InsertRemoveTransactionMultipart(ctx context.Context, projectID, collectionID string, searchCacheControl string, request InsertRemoveRequest, files [][]byte) (*InsertRemoveResult, error) {
	var result InsertRemoveResult
	query := url.Values{}
	addQuery(query, "SWA-searchCacheControl", searchCacheControl)
	if err := c.doMultipartJSON(ctx, httpMethodPost, []string{"projects", projectID, "collections", collectionID, "records", "insertRemoveTransaction"}, query, "request", request, files, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) StartInsertRemoveTransaction(ctx context.Context, projectID, collectionID string, request StartTransactionRequest) (*StartTransactionResult, error) {
	var result StartTransactionResult
	if err := c.doJSON(ctx, httpMethodPost, []string{"projects", projectID, "collections", collectionID, "records", "bulkInsertRemoveTransaction"}, nil, nil, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) CommitInsertRemoveTransaction(ctx context.Context, projectID, collectionID, indexingBufferID string, request FinishTransactionRequest) (*FinishTransactionResponse, error) {
	var result FinishTransactionResponse
	if err := c.doJSON(ctx, httpMethodPost, []string{"projects", projectID, "collections", collectionID, "records", "bulkInsertRemoveTransaction", indexingBufferID, "end"}, nil, nil, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetFlushJobStatus(ctx context.Context, projectID, collectionID, indexingBufferID, jobID string) (*JobStatusResponse, error) {
	var result JobStatusResponse
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "records", "bulkInsertRemoveTransaction", indexingBufferID, "end", jobID}, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) AddToInsertRemoveTransaction(ctx context.Context, projectID, collectionID, indexingBufferID, searchCacheControl string, request InsertRemoveRequest) (*InsertRemoveResult, error) {
	var result InsertRemoveResult
	query := url.Values{}
	addQuery(query, "SWA-searchCacheControl", searchCacheControl)
	if err := c.doJSON(ctx, httpMethodPost, []string{"projects", projectID, "collections", collectionID, "records", "bulkInsertRemoveTransaction", indexingBufferID, "buffer"}, query, nil, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) AddToInsertRemoveTransactionMultipart(ctx context.Context, projectID, collectionID, indexingBufferID, searchCacheControl string, request InsertRemoveRequest, files [][]byte) (*InsertRemoveResult, error) {
	var result InsertRemoveResult
	query := url.Values{}
	addQuery(query, "SWA-searchCacheControl", searchCacheControl)
	if err := c.doMultipartJSON(ctx, httpMethodPost, []string{"projects", projectID, "collections", collectionID, "records", "bulkInsertRemoveTransaction", indexingBufferID, "buffer"}, query, "request", request, files, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
