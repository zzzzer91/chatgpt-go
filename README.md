# ChatGPT go client

Usage:

```go
func main() {
	secretKey := ""
	cli := chatgpt.NewService(secretKey, chatgpt.WithHost("api.openai.com"), chatgpt.WithTimeout(15*time.Second))
	msgs := []*Message{
		{Role: RoleTypeUser, Content: "who are you"},
	}

	ctx := context.Background()
	resp, err := cli.Chat(ctx, msgs)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Choices[0].Message.Content)

	err = s.ChatStream(ctx, msgs, func(resp *ChatResponse) error {
		fmt.Print(resp.Choices[0].Delta.Content)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
```