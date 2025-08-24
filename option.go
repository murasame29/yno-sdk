package yno

import (
	"time"

	"github.com/murasame29/yno-sdk/client"
)

type OptionFunc func([]client.Option) []client.Option

func WithMngAPIVersion(version string) OptionFunc {
	return func(o []client.Option) []client.Option {
		return append(o, client.WithHeader("X-Yamaha-YNO-MngAPI-Version", version))
	}
}

func WithExtraHeader(key, val string) OptionFunc {
	return func(o []client.Option) []client.Option {
		return append(o, client.WithHeader(key, val))
	}
}

func WithTimeout(timeout time.Duration) OptionFunc {
	return func(o []client.Option) []client.Option {
		return append(o, client.WithTimeout(timeout))
	}
}
