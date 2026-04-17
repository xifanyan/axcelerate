package searchwebapi

import "context"

func (c *Client) ListProjects(ctx context.Context) (*ProjectsResult, error) {
	var result ProjectsResult
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects"}, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetProjectResources(ctx context.Context, projectID string) (*ProjectResourcesResult, error) {
	var result ProjectResourcesResult
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID}, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
