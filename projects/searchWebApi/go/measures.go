package searchwebapi

import "context"

func (c *Client) GetMeasureCube(ctx context.Context, projectID, collectionID string, opts MeasureOptions, dimensions []DimensionRequest) (*MeasureCube, error) {
	var result MeasureCube
	if err := c.doJSON(ctx, httpMethodPost, []string{"projects", projectID, "collections", collectionID, "measures"}, opts.values(), nil, dimensions, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
