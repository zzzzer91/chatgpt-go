package chatgpt

import (
	"context"
	"fmt"
	"testing"
)

const (
	secretKey = ""
	host      = "api.openai.com"
)

var (
	msgs = []*Message{
		{Role: RoleTypeUser, Content: "who are you"},
	}
)

func TestChat(t *testing.T) {
	s := NewService(secretKey, WithHost(host))
	resp, err := s.Chat(context.Background(), msgs)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.Choices[0].Message.Content)
}

func TestChatStream(t *testing.T) {
	s := NewService(secretKey, WithHost(host))
	err := s.ChatStream(context.Background(), msgs, func(resp *ChatResponse) error {
		fmt.Print(resp.Choices[0].Delta.Content)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
