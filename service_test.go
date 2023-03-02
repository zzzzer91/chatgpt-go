package chatgpt

import (
	"fmt"
	"testing"
)

func TestChat(t *testing.T) {
	secretKey := ""
	s := NewService(secretKey)
	resp, err := s.Chat("who are you")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
