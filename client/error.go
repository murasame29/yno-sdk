package client

import "fmt"

type HTTPError struct {
	StatusCode int
	Message    string
	Body       string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("API error: status code %d, message: %s", e.StatusCode, e.Message)
}
