package ai

import (
	"context"
	"log"

	"github.com/AbhaySingh002/Pollen/internal/config"
	"github.com/AbhaySingh002/Pollen/internal/types"


	"google.golang.org/genai"
)

func IntentChecking(ctx context.Context,query string,out chan<- types.StreamMessage) {

	defer close(out)


	var appConfig config.Config
	appConfig = appConfig.Loader()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  config.Gemini_apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	temp := float32(0.7)
	thinkingBudget := int32(-1)

	generateConfig := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(appConfig.Prompt.IntentPrompt, genai.RoleUser),
		ThinkingConfig: &genai.ThinkingConfig{
			ThinkingBudget:  &thinkingBudget,
		},
		Temperature: &temp,
	}

	stream := client.Models.GenerateContentStream(
		ctx,
		config.IntentGeminiModel,
		genai.Text(query),
		generateConfig,
	)

	for chunk := range stream {
		for _, part := range chunk.Candidates[0].Content.Parts {
			if part.Thought && len(part.Text) > 0 {
				out <- types.StreamMessage{
					Type: "thought",
					Content: part.Text,
				}
			}else{
				out <- types.StreamMessage{
					Type: "xml",
					Content: part.Text,
				}
			}
		}
	}
}
