package gpt

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kackerx/interview/pkg/util"
)

const (
	API    = "https://api.gpt.ge/v1/chat/completions"
	prompt = "帮我翻译单引号内的内容, 只返回翻译后的文本: '%s'"
)

func Trans(ctx context.Context, content string) (string, error) {
	resp, err := util.Post(ctx, API, &GPTRequest{
		Model: "gpt-4o",
		Messages: []Message{
			{
				Role:    "user",
				Content: fmt.Sprintf(prompt, content),
			},
		},
		MaxTokens:   1688,
		Temperature: 0.5,
		Stream:      false,
	})
	if err != nil {
		return "", err
	}

	ret := new(GPTResponse)
	if err = json.Unmarshal(resp, ret); err != nil {
		return "", err
	}

	return ret.Choices[0].Message.Content, nil
}
