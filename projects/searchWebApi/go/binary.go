package searchwebapi

import "context"

func (c *Client) GetBinary(ctx context.Context, projectID, collectionID string, opts BinarySearchOptions) (*BinaryResponse, error) {
	return c.doBinary(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "binary"}, opts.values(), nil)
}

func (c *Client) GetBinaryByRecordID(ctx context.Context, projectID, collectionID, recordID, field string) (*BinaryResponse, error) {
	query := BinarySearchOptions{Field: field}.values()
	return c.doBinary(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "binary", recordID, "content"}, query, nil)
}
