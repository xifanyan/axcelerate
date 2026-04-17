package searchwebapi

import "context"

func (c *Client) WaitForPendingChanges(ctx context.Context, projectID, collectionID string, opts WaitForPendingChangesOptions) (*WaitForPendingChangesResult, error) {
	var result WaitForPendingChangesResult
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "changes", "queue"}, opts.values(), nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
