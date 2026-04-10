package adp

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

type ClientConfig struct {
	BaseURL  string
	Username string
	Password string
	Insecure bool
	Timeout  time.Duration
	Debug    bool
	DebugOut io.Writer
}

type Client struct {
	baseURL    string
	username   string
	password   string
	httpClient *http.Client
	debug      bool
	debugOut   io.Writer
}

func NewClient(cfg ClientConfig) (*Client, error) {
	if strings.TrimSpace(cfg.BaseURL) == "" {
		return nil, errors.New("base URL is required")
	}
	if strings.TrimSpace(cfg.Username) == "" {
		return nil, errors.New("username is required")
	}
	if strings.TrimSpace(cfg.Password) == "" {
		return nil, errors.New("password is required")
	}

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	debugOut := cfg.DebugOut
	if debugOut == nil {
		debugOut = io.Discard
	}

	return &Client{
		baseURL:    strings.TrimRight(cfg.BaseURL, "/"),
		username:   cfg.Username,
		password:   cfg.Password,
		httpClient: &http.Client{Timeout: timeout},
		debug:      cfg.Debug,
		debugOut:   debugOut,
	}, nil
}
