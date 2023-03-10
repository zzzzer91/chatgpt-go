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

func (i *serviceImpl) ChatWithMessages(msgs []*Message) (string, error) {
	p := ChatRequest{
		Model:       i.modelName,
		Temperature: i.temperature,
		TopP:        i.topP,
	}
	p.Messages = msgs
	url := "https://" + i.host + chatPath
	resp, err := i.client.PostJsonWithAuth(url, &p, i.secretKey)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var jd ChatResponse
	err = json.NewDecoder(resp.Body).Decode(&jd)
	if err != nil {
		return "", errors.Wrap(err, "json decode error")
	}
	return strings.TrimSpace(jd.Choices[0].Message.Content), nil
}

func (i *serviceImpl) Chat(text string) (string, error) {
	msgs := []*Message{
		{Role: RoleTypeSystem, Content: "You are a helpful assistant."},
		{Role: RoleTypeUser, Content: text},
	}
	return i.ChatWithMessages(msgs)
}
