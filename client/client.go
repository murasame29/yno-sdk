package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultTimeout = 30 * time.Second
)

type HTTPError struct {
	StatusCode int
	Message    string
	Body       string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("API error: status code %d, message: %s", e.StatusCode, e.Message)
}

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	headers    http.Header
}

type Option func(*Client)

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

func WithHeader(key, value string) Option {
	return func(c *Client) {
		c.headers.Set(key, value)
	}
}

func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.httpClient = hc
	}
}

func WithYnoVersion(version string) Option {
	return func(c *Client) {
		if len(version) != 0 {
			c.headers.Add("X-Yamaha-YNO-MngAPI-Version", version)
		}
	}
}

func WithYnoAPIkey(apiKey string) Option {
	return func(c *Client) {
		c.headers.Add("X-Yamaha-YNO-MngAPI-Key", apiKey)
	}
}

func NewClient(baseURLStr string, opts ...Option) (*Client, error) {
	baseURL, err := url.Parse(baseURLStr)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	c := &Client{
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
		baseURL: baseURL,
		headers: http.Header{},
	}

	c.headers.Set("Content-Type", "application/json")

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

func (c *Client) Do(ctx context.Context, method, path string, requestBody, responseBody any) error {
	rel, err := url.Parse(path)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}
	fullURL := c.baseURL.ResolveReference(rel)

	var bodyReader io.Reader
	if requestBody != nil {
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL.String(), bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header = c.headers.Clone()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return &HTTPError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("request failed with status code %d", resp.StatusCode),
			Body:       string(bodyBytes),
		}
	}

	if responseBody != nil {
		if err := json.NewDecoder(resp.Body).Decode(responseBody); err != nil {
			return fmt.Errorf("failed to decode response body: %w", err)
		}
	}

	return nil
}

func (c *Client) Get(ctx context.Context, path string, responseBody any) error {
	return c.Do(ctx, http.MethodGet, path, nil, responseBody)
}

func (c *Client) Post(ctx context.Context, path string, requestBody, responseBody any) error {
	return c.Do(ctx, http.MethodPost, path, requestBody, responseBody)
}

func (c *Client) Put(ctx context.Context, path string, requestBody, responseBody any) error {
	return c.Do(ctx, http.MethodPut, path, requestBody, responseBody)
}

func (c *Client) Delete(ctx context.Context, path string, responseBody any) error {
	return c.Do(ctx, http.MethodDelete, path, nil, responseBody)
}
