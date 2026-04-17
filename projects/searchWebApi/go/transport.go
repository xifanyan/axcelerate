package searchwebapi

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ResponseError struct {
	StatusCode int
	Body       string
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("unexpected HTTP status %d", e.StatusCode)
}

func (c *Client) newURL(segments []string, query url.Values) string {
	u := *c.baseURL
	u.Path = c.joinPath(segments...)
	if query != nil {
		u.RawQuery = query.Encode()
	} else {
		u.RawQuery = ""
	}
	return u.String()
}

func (c *Client) doJSON(ctx context.Context, method string, segments []string, query url.Values, headers http.Header, body any, out any) error {
	var reader io.Reader
	if body != nil {
		buf := &bytes.Buffer{}
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return err
		}
		reader = buf
	}

	req, err := http.NewRequestWithContext(ctx, method, c.newURL(segments, query), reader)
	if err != nil {
		return err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	c.applyHeaders(req, headers)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	c.captureSession(resp)

	if resp.StatusCode >= http.StatusBadRequest {
		data, _ := io.ReadAll(resp.Body)
		return &ResponseError{StatusCode: resp.StatusCode, Body: string(data)}
	}
	if out == nil {
		_, _ = io.Copy(io.Discard, resp.Body)
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (c *Client) doBinary(ctx context.Context, method string, segments []string, query url.Values, headers http.Header) (*BinaryResponse, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.newURL(segments, query), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/octet-stream")
	c.applyHeaders(req, headers)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	c.captureSession(resp)

	if resp.StatusCode >= http.StatusBadRequest {
		defer resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)
		return nil, &ResponseError{StatusCode: resp.StatusCode, Body: string(data)}
	}

	return &BinaryResponse{
		Header:      resp.Header.Clone(),
		ContentType: resp.Header.Get("Content-Type"),
		Body:        resp.Body,
	}, nil
}

func (c *Client) doNDJSON(ctx context.Context, segments []string, query url.Values) (*RecordStream, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.newURL(segments, query), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/x-ndjson")
	c.applyHeaders(req, nil)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	c.captureSession(resp)

	if resp.StatusCode >= http.StatusBadRequest {
		defer resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)
		return nil, &ResponseError{StatusCode: resp.StatusCode, Body: string(data)}
	}

	reader := bufio.NewReader(resp.Body)
	metaLine, err := reader.ReadBytes('\n')
	if err != nil && err != io.EOF {
		resp.Body.Close()
		return nil, err
	}

	metaLine = bytes.TrimSpace(metaLine)
	var meta SearchResult
	if err := json.Unmarshal(metaLine, &meta); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &RecordStream{Meta: meta, reader: reader, body: resp.Body}, nil
}

func (c *Client) doMultipartJSON(ctx context.Context, method string, segments []string, query url.Values, requestField string, requestBody any, files [][]byte, out any) error {
	body, contentType, err := writeMultipartJSONAndFiles(requestField, requestBody, files)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, method, c.newURL(segments, query), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", "application/json")
	c.applyHeaders(req, nil)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	c.captureSession(resp)

	if resp.StatusCode >= http.StatusBadRequest {
		data, _ := io.ReadAll(resp.Body)
		return &ResponseError{StatusCode: resp.StatusCode, Body: string(data)}
	}
	if out == nil {
		_, _ = io.Copy(io.Discard, resp.Body)
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (c *Client) applyHeaders(req *http.Request, headers http.Header) {
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	if c.auth.username != "" || c.auth.password != "" {
		req.SetBasicAuth(c.auth.username, c.auth.password)
	}
	if c.auth.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.auth.bearerToken)
	}
	if sessionID := c.SessionID(); sessionID != "" {
		req.Header.Set("SWA-SESSION", sessionID)
	}
	if c.sessionType != "" {
		req.Header.Set("SWA-SESSION-TYPE", c.sessionType)
	}
	if c.mdcToken != "" {
		req.Header.Set("SWA-MDC-TOKEN", c.mdcToken)
	}
	if c.mdcMethod != "" {
		req.Header.Set("SWA-MDC-METHOD", c.mdcMethod)
	}
}

func (c *Client) captureSession(resp *http.Response) {
	if sessionID := strings.TrimSpace(resp.Header.Get("SWA-SESSION")); sessionID != "" {
		c.SetSessionID(sessionID)
	}
}

func boolPtrString(v *bool) string {
	if v == nil {
		return ""
	}
	return strconv.FormatBool(*v)
}

func intPtrString(v *int) string {
	if v == nil {
		return ""
	}
	return strconv.Itoa(*v)
}

func int64PtrString(v *int64) string {
	if v == nil {
		return ""
	}
	return strconv.FormatInt(*v, 10)
}

func addQuery(values url.Values, key, value string) {
	if strings.TrimSpace(value) != "" {
		values.Set(key, value)
	}
}

func writeMultipartJSONAndFiles(fieldName string, requestBody any, files [][]byte) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormField(fieldName)
	if err != nil {
		return nil, "", err
	}
	if err := json.NewEncoder(part).Encode(requestBody); err != nil {
		return nil, "", err
	}
	for i, data := range files {
		filePart, err := writer.CreateFormFile("binaries", fmt.Sprintf("binary-%d.bin", i))
		if err != nil {
			return nil, "", err
		}
		if _, err := filePart.Write(data); err != nil {
			return nil, "", err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return body, writer.FormDataContentType(), nil
}
