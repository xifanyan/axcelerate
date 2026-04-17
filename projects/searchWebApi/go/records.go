package searchwebapi

import (
	"context"
	"net/url"
	"strconv"
)

func (c *Client) SearchRecords(ctx context.Context, projectID, collectionID string, opts SearchRecordsOptions) (*SearchResult, error) {
	var result SearchResult
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "records"}, opts.values(), nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) SearchRecordsStream(ctx context.Context, projectID, collectionID string, opts SearchRecordsOptions) (*RecordStream, error) {
	return c.doNDJSON(ctx, []string{"projects", projectID, "collections", collectionID, "records"}, opts.values())
}

func (c *Client) ChangeAllInSearchResult(ctx context.Context, projectID, collectionID string, opts SearchRecordsOptions, requests []ChangeRequest) (*ChangeResult, error) {
	var result ChangeResult
	if err := c.doJSON(ctx, httpMethodPut, []string{"projects", projectID, "collections", collectionID, "records"}, opts.values(), nil, requests, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetSearchResultToken(ctx context.Context, projectID, collectionID string, opts SearchTokenOptions) (*SearchResultTokenResponse, error) {
	var result SearchResultTokenResponse
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "searchToken"}, opts.values(), nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) DeleteSearchResultToken(ctx context.Context, projectID, collectionID string, token SearchResultToken) error {
	return c.doJSON(ctx, httpMethodDelete, []string{"projects", projectID, "collections", collectionID, "searchToken"}, nil, nil, token, nil)
}

func (c *Client) TouchSearchResultToken(ctx context.Context, projectID, collectionID string, token SearchResultToken) (*SearchResultTokenResponse, error) {
	var result SearchResultTokenResponse
	if err := c.doJSON(ctx, httpMethodPut, []string{"projects", projectID, "collections", collectionID, "searchToken"}, nil, nil, token, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) CreateSortOrderSnapshot(ctx context.Context, projectID, collectionID string, opts SortOrderSnapshotOptions, token SearchResultToken) error {
	return c.doJSON(ctx, httpMethodPost, []string{"projects", projectID, "collections", collectionID, "searchToken", "sortOrderSnapshot"}, opts.values(), nil, token, nil)
}

func (c *Client) GetHighlightingForSearchResult(ctx context.Context, projectID, collectionID string, opts GetHighlightExpressionsOptions) (*SearchResultHighlightingResult, error) {
	var result SearchResultHighlightingResult
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "search", "highlightExpression"}, opts.values(), nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetRecordResources(ctx context.Context, projectID, collectionID, recordID string) (*RecordResourcesResult, error) {
	var result RecordResourcesResult
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "records", recordID}, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) FetchRecordContent(ctx context.Context, projectID, collectionID, recordID string, opts FetchRecordContentOptions) (*Record, error) {
	var result Record
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "records", recordID, "content"}, opts.values(), nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) ChangeRecordContent(ctx context.Context, projectID, collectionID, recordID string, opts ChangeOptions, requests []ChangeRequest) (*ChangeResult, error) {
	var result ChangeResult
	if err := c.doJSON(ctx, httpMethodPut, []string{"projects", projectID, "collections", collectionID, "records", recordID, "content"}, opts.values(), nil, requests, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) ChangeRecordContentMultipart(ctx context.Context, projectID, collectionID, recordID string, opts ChangeOptions, requests []ChangeRequest, files [][]byte) (*ChangeResult, error) {
	var result ChangeResult
	requestBody := map[string]any{"request": requests}
	if err := c.doMultipartJSON(ctx, httpMethodPut, []string{"projects", projectID, "collections", collectionID, "records", recordID, "content"}, opts.values(), "request", requestBody, files, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) SearchInDocumentText(ctx context.Context, projectID, collectionID, recordID string, opts InDocumentSearchOptions) (*HighlightedWordResult, error) {
	var result HighlightedWordResult
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "records", recordID, "inDocumentSearch"}, opts.values(), nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) ChangeRecordInDocumentContext(ctx context.Context, projectID, collectionID, recordID string, opts ChangeOptions, requests []ChangeRequest) (*ChangeResult, error) {
	var result ChangeResult
	if err := c.doJSON(ctx, httpMethodPut, []string{"projects", projectID, "collections", collectionID, "records", recordID, "inDocumentSearch"}, opts.values(), nil, requests, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) ChangeRecordInDocumentContextMultipart(ctx context.Context, projectID, collectionID, recordID string, opts ChangeOptions, requests []ChangeRequest, files [][]byte) (*ChangeResult, error) {
	var result ChangeResult
	requestBody := map[string]any{"request": requests}
	if err := c.doMultipartJSON(ctx, httpMethodPut, []string{"projects", projectID, "collections", collectionID, "records", recordID, "inDocumentSearch"}, opts.values(), "request", requestBody, files, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

const (
	httpMethodGet    = "GET"
	httpMethodPost   = "POST"
	httpMethodPut    = "PUT"
	httpMethodDelete = "DELETE"
)

func addFolderProperties(values url.Values, raw string) {
	if raw != "" {
		values.Set("folderFieldsWithProperties", raw)
	}
}

func addOptionalInt(values url.Values, key string, v *int) {
	if v != nil {
		values.Set(key, strconv.Itoa(*v))
	}
}
