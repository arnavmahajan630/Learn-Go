package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func Chat(cleanedInput string) (string, error){

	apikey := os.Getenv("API_KEY")
	if apikey == ""{
		return "", fmt.Errorf("Api Key required")
	}


	uri := "https://openrouter.ai/api/v1"

	client := openai.NewClient(
		option.WithBaseURL(uri),
		option.WithAPIKey(apikey),
	)

	ctx := context.Background()
	messages := []openai.ChatCompletionMessageParamUnion{}
	model := "mistralai/devstral-2512:free"

	messages = append(messages, openai.UserMessage(cleanedInput))

	params := openai.ChatCompletionNewParams{
		Model: model,
		Messages: messages,
	}

	resp , err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", err
	}
	msg := resp.Choices[0].Message.Content
	return msg, nil


}