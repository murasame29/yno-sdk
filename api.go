package yno

import (
	"github.com/murasame29/yno-sdk/client"
)

type YNOClient struct {
	APIKey        string
	MngAPIVersion string
	client        *client.Client
}

func NewClient(apiKey string, MngAPIVersion string, opts ...client.Option) (*YNOClient, error) {
	opts = append(opts, client.WithYnoAPIkey(apiKey), client.WithYnoVersion(MngAPIVersion))

	client, err := client.NewClient("https://yno-mngapi.netvolante.jp", opts...)
	if err != nil {
		return nil, err
	}

	return &YNOClient{
		APIKey:        apiKey,
		MngAPIVersion: MngAPIVersion,
		client:        client,
	}, nil
}
