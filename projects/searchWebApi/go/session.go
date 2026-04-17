package searchwebapi

import "context"

func (c *Client) Login(ctx context.Context) (*LoginResult, error) {
	var result LoginResult
	if err := c.doJSON(ctx, httpMethodPost, []string{"login"}, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) Logout(ctx context.Context) (*LogoutResult, error) {
	var result LogoutResult
	if err := c.doJSON(ctx, httpMethodDelete, []string{"logout"}, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
