package searchwebapi

import "context"

func (c *Client) ListCollections(ctx context.Context, projectID string) (*CollectionsResult, error) {
	var result CollectionsResult
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections"}, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetCollectionResources(ctx context.Context, projectID, collectionID string) (*CollectionResourcesResult, error) {
	var result CollectionResourcesResult
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID}, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetFields(ctx context.Context, projectID, collectionID string) (*FieldsResult, error) {
	var result FieldsResult
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "fields"}, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetFolderFields(ctx context.Context, projectID, collectionID string) (*FieldsResult, error) {
	var result FieldsResult
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "filters"}, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetFolderFieldResources(ctx context.Context, projectID, collectionID, fieldID string) (*FolderFieldResourcesResult, error) {
	var result FolderFieldResourcesResult
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "filters", fieldID}, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetFolderValues(ctx context.Context, projectID, collectionID, fieldID string, opts FolderValuesOptions) (*FolderValuesResult, error) {
	var result FolderValuesResult
	if err := c.doJSON(ctx, httpMethodGet, []string{"projects", projectID, "collections", collectionID, "filters", fieldID, "values"}, opts.values(), nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
