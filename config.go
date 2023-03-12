package chatgpt

import (
	"time"
)

type clientConfig struct {
	timeout     time.Duration
	host        string
	path        string
	modelName   string
	temperature float64
	topP        float64
}

func newClientConfig() *clientConfig {
	return &clientConfig{
		timeout:     30 * time.Second,
		host:        "api.openai.com",
		path:        "/v1/chat/completions",
		modelName:   "gpt-3.5-turbo",
		temperature: 1,
		topP:        1,
	}
}

type Option func(s *clientConfig)

func WithTimeout(timeout time.Duration) Option {
	return func(s *clientConfig) {
		s.timeout = timeout
	}
}

func WithHost(host string) Option {
	return func(s *clientConfig) {
		s.host = host
	}
}

func WithPath(path string) Option {
	return func(s *clientConfig) {
		s.path = path
	}
}

func WithModelName(modelName string) Option {
	return func(s *clientConfig) {
		s.modelName = modelName
	}
}

// WithTemperature sets temperature.
// What sampling temperature to use, between 0 and 2.
// Higher values like 0.8 will make the output more random,
// while lower values like 0.2 will make it more focused and deterministic.
// We generally recommend altering this or top_p but not both.
func WithTemperature(t float64) Option {
	return func(s *clientConfig) {
		s.temperature = t
	}
}

// WithTopP sets top_p
// An alternative to sampling with temperature, called nucleus sampling,
// where the model considers the results of the tokens with top_p probability mass.
// So 0.1 means only the tokens comprising the top 10% probability mass are considered.
// We generally recommend altering this or temperature but not both.
func WithTopP(t float64) Option {
	return func(s *clientConfig) {
		s.topP = t
	}
}
