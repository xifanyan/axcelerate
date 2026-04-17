package searchwebapi

import (
	"errors"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
)

type Config struct {
	BaseURL     string
	Username    string
	Password    string
	BearerToken string
	SessionID   string
	SessionType string
	MDCToken    string
	MDCMethod   string
	HTTPClient  *http.Client
}

type Client struct {
	baseURL     *url.URL
	httpClient  *http.Client
	auth        authConfig
	sessionType string
	mdcToken    string
	mdcMethod   string

	mu        sync.RWMutex
	sessionID string
}

func NewClient(cfg Config) (*Client, error) {
	baseURL := strings.TrimSpace(cfg.BaseURL)
	if baseURL == "" {
		return nil, errors.New("base URL is required")
	}

	parsed, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	parsed.Path = normalizeBasePath(parsed.Path)
	parsed.RawQuery = ""
	parsed.Fragment = ""

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		baseURL:    parsed,
		httpClient: httpClient,
		auth: authConfig{
			username:    cfg.Username,
			password:    cfg.Password,
			bearerToken: cfg.BearerToken,
		},
		sessionType: cfg.SessionType,
		mdcToken:    cfg.MDCToken,
		mdcMethod:   cfg.MDCMethod,
		sessionID:   cfg.SessionID,
	}, nil
}

func normalizeBasePath(existing string) string {
	trimmed := strings.TrimRight(existing, "/")
	if trimmed == "" {
		return "/searchWebApi"
	}
	if strings.HasSuffix(trimmed, "/searchWebApi") {
		return trimmed
	}
	return trimmed + "/searchWebApi"
}

func (c *Client) SessionID() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.sessionID
}

func (c *Client) SetSessionID(sessionID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sessionID = sessionID
}

func (c *Client) ClearSessionID() {
	c.SetSessionID("")
}

func (c *Client) joinPath(segments ...string) string {
	parts := make([]string, 0, len(segments))
	for _, segment := range segments {
		if trimmed := strings.Trim(segment, "/"); trimmed != "" {
			parts = append(parts, trimmed)
		}
	}
	if len(parts) == 0 {
		return c.baseURL.Path
	}
	return path.Join(c.baseURL.Path, path.Join(parts...))
}
