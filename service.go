package chatgpt

import (
	"github.com/zzzzer91/httpgo"
)

type Service interface {
	Chat(msgs []*Message) (*ChatResponse, error)
	ChatStream(msgs []*Message, f func(*ChatResponse) error) error
}

func NewService(secretKey string, opts ...Option) Service {
	conf := newClientConfig()
	for _, o := range opts {
		o(conf)
	}
	return &serviceImpl{
		client:       httpgo.NewClient(conf.timeout, nil),
		token:        "Bearer " + secretKey,
		clientConfig: conf,
	}
}
