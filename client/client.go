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

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

func (c *Client) clone() *Client {
	newHTTPClient := &http.Client{
		Transport:     c.httpClient.Transport,
		CheckRedirect: c.httpClient.CheckRedirect,
		Jar:           c.httpClient.Jar,
		Timeout:       c.httpClient.Timeout,
	}

	return &Client{
		httpClient: newHTTPClient,
		baseURL:    c.baseURL,
		headers:    c.headers.Clone(),
	}
}

func (c *Client) do(ctx context.Context, method, path string, requestBody, responseBody any) error {
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

func (c *Client) Do(ctx context.Context, method, path string, requestBody, responseBody any, opts ...Option) error {
	if len(opts) == 0 {
		return c.do(ctx, method, path, requestBody, responseBody)
	}

	clonedClient := c.clone()
	for _, opt := range opts {
		if opt != nil {
			opt(clonedClient)
		}
	}
	return clonedClient.do(ctx, method, path, requestBody, responseBody)
}

func (c *Client) Get(ctx context.Context, path string, query map[string]string, responseBody any, clientOpts ...Option) error {
	u, err := url.Parse(path)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}
	q := u.Query()

	for k, v := range query {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()

	return c.Do(ctx, http.MethodGet, u.String(), nil, responseBody, clientOpts...)
}

func (c *Client) Post(ctx context.Context, path string, requestBody, responseBody any, clientOpts ...Option) error {
	clientOpts = append(clientOpts, WithHeader("Content-Type", "application/json"))
	return c.Do(ctx, http.MethodPost, path, requestBody, responseBody, clientOpts...)
}

func (c *Client) Put(ctx context.Context, path string, requestBody, responseBody any, clientOpts ...Option) error {
	clientOpts = append(clientOpts, WithHeader("Content-Type", "application/json"))
	return c.Do(ctx, http.MethodPut, path, requestBody, responseBody, clientOpts...)
}

func (c *Client) Delete(ctx context.Context, path string, responseBody any, clientOpts ...Option) error {
	return c.Do(ctx, http.MethodDelete, path, nil, responseBody, clientOpts...)
}
