package chatgpt

import (
	"github.com/zzzzer91/httpgo"
)

type Service interface {
	ChatWithMessages(msgs []*Message) (*ChatResponse, error)
	Chat(text string) (string, error)
}

func NewService(secretKey string, opts ...Option) Service {
	conf := newClientConfig()
	for _, o := range opts {
		o(conf)
	}
	return &serviceImpl{
		client:       httpgo.NewClient(conf.timeout, nil),
		secretKey:    secretKey,
		clientConfig: conf,
	}
}
