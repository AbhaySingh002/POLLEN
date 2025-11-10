package ai

import (
	"context"

	"github.com/AbhaySingh002/Pollen/internal/config"

	"google.golang.org/genai"
)

func AiCaller(ctx context.Context, query string, out chan<- string, thoughtChan chan<- string, isIntent bool) {
	defer close(out)
	defer func() {
		if thoughtChan != nil {
			close(thoughtChan)
		}
	}()

	var temp float32
	var thinkingBudget int32
	var sysPrompt string
	var model string
	var includeThoughts bool

	var appConfig config.Config
	appConfig = appConfig.Loader()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  config.Gemini_apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		panic("Error creating genai client:")
	}

	if isIntent {
		temp = float32(0.5)
		thinkingBudget = int32(-1)
		sysPrompt = appConfig.Prompt.IntentPrompt
		model = config.IntentGeminiModel
		includeThoughts = false
	} else {
		temp = float32(0.5)
		thinkingBudget = int32(-1)
		sysPrompt = appConfig.Prompt.SysPrompt
		model = config.CoderGeminiModel
		includeThoughts = true
	}

	generateConfig := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(sysPrompt, genai.RoleUser),
		ThinkingConfig: &genai.ThinkingConfig{
			ThinkingBudget: &thinkingBudget,
			IncludeThoughts: includeThoughts,
		},
		Temperature: &temp,
	}

	stream := client.Models.GenerateContentStream(
		ctx,
		model,
		genai.Text(query),
		generateConfig,
	)

	for chunk, err := range stream {
		if err != nil {
			panic("Error in Generating Stream At Gemini")
		}
		if len(chunk.Candidates) == 0 || len(chunk.Candidates[0].Content.Parts) == 0 {
			continue
		}
		for _, part := range chunk.Candidates[0].Content.Parts {
			if len(part.Text) == 0 {
				continue
			}
			if part.Thought {
				if thoughtChan != nil {
					thoughtChan <- part.Text
				}
			} else {
				out <- part.Text
			}
		}
	}
}