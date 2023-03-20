package chatgpt

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	"github.com/zzzzer91/httpgo"
)

type serviceImpl struct {
	client    *httpgo.Client
	secretKey string
	*clientConfig
}

var _ Service = (*serviceImpl)(nil)

func (i *serviceImpl) ChatWithMessages(msgs []*Message) (*ChatResponse, error) {
	p := ChatRequest{
		Model:       i.modelName,
		Temperature: i.temperature,
		TopP:        i.topP,
	}
	p.Messages = msgs
	url := "https://" + i.host + i.path
	resp, err := i.client.PostJsonWithAuth(url, &p, i.secretKey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jd ChatResponse
	err = json.NewDecoder(resp.Body).Decode(&jd)
	if err != nil {
		return nil, errors.Wrap(err, "json decode error")
	}
	return &jd, nil
}

func (i *serviceImpl) Chat(text string) (string, error) {
	msgs := []*Message{
		{Role: RoleTypeSystem, Content: i.defaultSystemMsg},
		{Role: RoleTypeUser, Content: text},
	}
	resp, err := i.ChatWithMessages(msgs)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}
