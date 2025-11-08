package config

import (
	_ "embed"
)

//go:embed intentPrompt.txt
var intentPrompt string

//go:embed basePrompt.txt
var sysPrompt string

type Prompt struct {
	IntentPrompt string
	SysPrompt    string
}

type Config struct {
	Prompt Prompt
}

func (c *Config) Loader() Config {
	c.Prompt = Prompt{
		IntentPrompt: intentPrompt,
		SysPrompt:    sysPrompt,
	}
	return *c
}
