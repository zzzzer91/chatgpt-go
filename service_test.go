package chatgpt

import (
	"fmt"
	"testing"
)

const (
	secretKey = ""
	host      = ""
)

func TestChat(t *testing.T) {
	s := NewService(secretKey, WithHost(host))
	resp, err := s.ChatWithText("who are you")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
