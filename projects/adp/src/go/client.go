package adp

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type rawTaskRequest struct {
	TaskType          string         `json:"taskType"`
	TaskConfiguration map[string]any `json:"taskConfiguration,omitempty"`
	TaskDescription   string         `json:"taskDescription,omitempty"`
	TaskDisplayName   string         `json:"taskDisplayName,omitempty"`
}

type statusRequest struct {
	ExecutionID string `json:"executionId"`
}

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
	baseURL := strings.TrimSpace(cfg.BaseURL)
	if baseURL == "" {
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

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: cfg.Insecure}

	return &Client{
		baseURL:    strings.TrimRight(baseURL, "/"),
		username:   cfg.Username,
		password:   cfg.Password,
		httpClient: &http.Client{Timeout: timeout, Transport: transport},
		debug:      cfg.Debug,
		debugOut:   debugOut,
	}, nil
}

func (c *Client) execute(ctx context.Context, endpoint string, req rawTaskRequest) (*TaskResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	if c.debug {
		fmt.Fprintf(c.debugOut, "request body: %s\n", body)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, c.baseURL+endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Auth-Username", c.username)
	httpReq.Header.Set("Auth-Password", c.password)

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if c.debug {
		fmt.Fprintf(c.debugOut, "response body: %s\n", respBody)
	}

	if httpResp.StatusCode < http.StatusOK || httpResp.StatusCode >= http.StatusMultipleChoices {
		message := fmt.Sprintf("unexpected HTTP status %s", httpResp.Status)
		if bodyText := strings.TrimSpace(string(respBody)); bodyText != "" {
			message += ": " + bodyText
		}
		return nil, errors.New(message)
	}

	var resp TaskResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if resp.ExecutionStatus == "failed" {
		return nil, &TaskExecutionError{
			ExecutionID:       resp.ExecutionID,
			TaskType:          resp.TaskType,
			ExecutionStatus:   resp.ExecutionStatus,
			ErrorMessage:      resp.ErrorMessage,
			ExecutionMetaData: resp.ExecutionMetaData,
		}
	}

	return &resp, nil
}

func (c *Client) Poll(ctx context.Context, executionID string) (*TaskResponse, error) {
	body, err := json.Marshal(statusRequest{ExecutionID: executionID})
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	if c.debug {
		fmt.Fprintf(c.debugOut, "request body: %s\n", body)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, c.baseURL+"/statusAndProgress", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Auth-Username", c.username)
	httpReq.Header.Set("Auth-Password", c.password)

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if c.debug {
		fmt.Fprintf(c.debugOut, "response body: %s\n", respBody)
	}

	if httpResp.StatusCode < http.StatusOK || httpResp.StatusCode >= http.StatusMultipleChoices {
		message := fmt.Sprintf("unexpected HTTP status %s", httpResp.Status)
		if bodyText := strings.TrimSpace(string(respBody)); bodyText != "" {
			message += ": " + bodyText
		}
		return nil, errors.New(message)
	}

	var resp TaskResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	if resp.ExecutionStatus == "" {
		return nil, errors.New("invalid polling response: missing executionStatus")
	}

	if resp.ExecutionStatus == "failed" {
		return nil, &TaskExecutionError{
			ExecutionID:       resp.ExecutionID,
			TaskType:          resp.TaskType,
			ExecutionStatus:   resp.ExecutionStatus,
			ErrorMessage:      resp.ErrorMessage,
			ExecutionMetaData: resp.ExecutionMetaData,
		}
	}

	return &resp, nil
}

func (c *Client) Wait(ctx context.Context, executionID string, interval time.Duration) (*TaskResponse, error) {
	if interval <= 0 {
		interval = time.Second
	}

	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		resp, err := c.Poll(ctx, executionID)
		if err != nil {
			return nil, err
		}
		if resp.ExecutionStatus == "success" {
			return resp, nil
		}

		timer := time.NewTimer(interval)
		select {
		case <-ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}
			return nil, ctx.Err()
		case <-timer.C:
		}
	}
}
