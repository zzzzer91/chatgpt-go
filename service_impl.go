package chatgpt

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/zzzzer91/httpgo"
)

type serviceImpl struct {
	client      *httpgo.Client
	secretKey   string
	temperature float64
	topP        float64
}

var _ Service = (*serviceImpl)(nil)

func (s *serviceImpl) ChatWithMessages(msgs []Message) (string, error) {
	p := ChatRequest{
		Model:       modelName,
		Temperature: s.temperature,
		TopP:        s.topP,
	}
	p.Messages = msgs
	resp, err := s.client.PostJsonWithAuth(chatUrl, &p, s.secretKey)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var jd ChatResponse
	err = json.NewDecoder(resp.Body).Decode(&jd)
	if err != nil {
		return "", errors.Wrap(err, "json decode error")
	}
	return jd.Choices[0].Message.Content, nil
}

func (s *serviceImpl) Chat(text string) (string, error) {
	msgs := []Message{
		{Role: "system", Content: "You are a helpful assistant."},
		{Role: "user", Content: text},
	}
	return s.ChatWithMessages(msgs)
}
