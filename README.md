ChatGPT go client

# Usage

```go
func main() {
	secretKey := ""
	cli := chatgpt.NewService(secretKey, chatgpt.WithHost("api.openai.com"), chatgpt.WithTimeout(15*time.Second))
	answer, err := cli.Chat("Who are you?")
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}
```