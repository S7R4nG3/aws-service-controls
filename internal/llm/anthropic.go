package llm

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/S7R4nG3/aws-service-controls/internal/utils"
	"github.com/anthropics/anthropic-sdk-go"
)

type LlmPrompt struct {
	Text      string
	Response  string
	Outfile   string
	Prefill   string
	Model     anthropic.Model
	MaxTokens int
}

func NewLlmPrompt(opts ...func(*LlmPrompt)) *LlmPrompt {
	np := &LlmPrompt{
		Model:     anthropic.ModelClaudeSonnet4_6,
		MaxTokens: 20000,
	}
	for _, o := range opts {
		o(np)
	}
	return np
}

func WithText(text string) func(*LlmPrompt) {
	return func(p *LlmPrompt) {
		p.Text = text
	}
}

func WithOutfile(filename string) func(*LlmPrompt) {
	return func(p *LlmPrompt) {
		p.Outfile = filename
	}
}

func WithPrefill(text string) func(*LlmPrompt) {
	return func(p *LlmPrompt) {
		p.Prefill = text
	}
}

func (p LlmPrompt) Run() string {
	client := anthropic.NewClient()
	file, _ := os.Create(p.Outfile)
	defer file.Close()
	writer := bufio.NewWriter(file)
	text := p.Text
	if p.Prefill != "" {
		text = fmt.Sprintf("%s\n\nYour response must begin with exactly: %s", p.Text, p.Prefill)
	}
	messages := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock(text)),
	}
	stream := client.Messages.NewStreaming(context.TODO(), anthropic.MessageNewParams{
		Model:     p.Model,
		MaxTokens: int64(p.MaxTokens),
		Messages:  messages,
	})

	message := anthropic.Message{}
	for stream.Next() {
		event := stream.Current()
		err := message.Accumulate(event)
		utils.Check(err, "Error accumulating streaming events...")

		switch eventVariant := event.AsAny().(type) {
		case anthropic.ContentBlockDeltaEvent:
			switch deltaVariant := eventVariant.Delta.AsAny().(type) {
			case anthropic.TextDelta:
				fmt.Print(".")
				writer.WriteString(deltaVariant.Text)
			}
		}
	}
	utils.Check(stream.Err(), "Error streaming response from Anthropic API...")
	writer.Flush()
	contents, _ := io.ReadAll(file)
	p.Response = string(contents)
	return p.Response
}
