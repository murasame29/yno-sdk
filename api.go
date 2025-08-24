package yno

import (
	"github.com/murasame29/yno-sdk/client"
)

type ynoClient struct {
	APIKey        string
	MngAPIVersion string
	client        *client.Client
}

const YNO_BASE_URL = "https://yno-mngapi.netvolante.jp"

func NewClient(baseURL, apiKey string, opts ...client.Option) (*ynoClient, error) {
	opts = append(opts, client.WithHeader("X-Yamaha-YNO-MngAPI-Key", apiKey))

	client, err := client.NewClient(baseURL, opts...)
	if err != nil {
		return nil, err
	}

	return &ynoClient{
		APIKey: apiKey,
		client: client,
	}, nil
}
