package chatgpt

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
	"github.com/zzzzer91/httpgo"
)

type serviceImpl struct {
	client *httpgo.Client
	token  string
	*clientConfig
}

var _ Service = (*serviceImpl)(nil)

func (i *serviceImpl) Chat(msgs []*Message) (*ChatResponse, error) {
	p := ChatRequest{
		Model:       i.modelName,
		Temperature: i.temperature,
		TopP:        i.topP,
	}
	p.Messages = msgs
	url := "https://" + i.host + i.path
	resp, err := i.client.PostJsonWithAuth(url, &p, i.token)
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

func (i *serviceImpl) ChatStream(msgs []*Message, f func(*ChatResponse) error) error {
	p := ChatRequest{
		Model:       i.modelName,
		Temperature: i.temperature,
		TopP:        i.topP,
		Stream:      true,
	}
	p.Messages = msgs
	url := "https://" + i.host + i.path
	header := httpgo.Header{Key: "Accept", Val: "text/event-stream"}
	resp, err := i.client.PostJsonWithAuth(url, &p, i.token, header)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r := bufio.NewReader(resp.Body)
	for {
		buf, err := r.ReadSlice('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return errors.Wrap(err, "ReadBytes error")
		}
		if string(buf) == "\n" {
			continue
		}
		if len(buf) <= 7 {
			return errors.New("length of data is invalid")
		}
		// Remove "data: " from the start of the line
		// Remove "\n" from the end of the line
		buf = buf[6 : len(buf)-1]
		if string(buf) == "[DONE]" {
			break
		}
		m := new(ChatResponse)
		err = json.Unmarshal(buf, m)
		if err != nil {
			return errors.Wrap(err, "json.Unmarshal")
		}

		err = f(m)
		if err != nil {
			return errors.Wrap(err, "execute f() error")
		}
	}

	return nil
}
