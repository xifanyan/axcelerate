package searchwebapi

import "context"

func (c *Client) GetCachedSearches(ctx context.Context, projectID, collectionID string) ([]CachedSearchDescription, error) {
	var result []CachedSearchDescription
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "cachedSearches"}, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) DropCachedSearches(ctx context.Context, projectID, collectionID string, opts CachedSearchesDeleteOptions) ([]CachedSearchDescription, error) {
	var result []CachedSearchDescription
	if err := c.doJSON(ctx, httpMethodDelete, []string{"projects", projectID, "collections", collectionID, "cachedSearches"}, opts.values(), nil, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
