// Used under MIT license from risafj/chat-stream

package main

import (
	"context"
	"log"
	"time"

	"github.com/sashabaranov/go-openai"
)

type ChatGPTClient struct {
	Client              *openai.Client
	ctx                 context.Context
	chatHistory         []openai.ChatCompletionMessage
	fixedInput			[]openai.ChatCompletionMessage
}

func CreateChatClient(apiKey string) *ChatGPTClient {
	return &ChatGPTClient{
		Client:              openai.NewClient(apiKey),
		ctx:                 context.Background(),
	}
}

func (c *ChatGPTClient) setFixedInput() {
	initialPromptString := "Translate the following Korean text into English while preserving its original meaning and style. The translation should be as faithful as possible, avoiding embellishments, omissions, or rephrasings beyond what is necessary for natural readability. Ensure that all narration is consistently rendered in the past tense. Keep character dialogue natural and reflective of the original tone. Do not add or modify content.";
	firstInputPrompt := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: initialPromptString,
	}
	c.fixedInput = append(c.fixedInput, firstInputPrompt)
}

func (c *ChatGPTClient) makeChatGPTRequest(query openai.ChatCompletionRequest) (string, error) {
	contextTimeout, cancel := context.WithTimeout(c.ctx, time.Minute*5)
	defer cancel()
	res, err := c.Client.CreateChatCompletion(contextTimeout, query)
	if err != nil {
		log.Printf("ChatCompletion Error: %v\n", err)
		return "", err
	}
	return res.Choices[0].Message.Content, nil
}

func (c *ChatGPTClient) addNewMessageToChatHistory(message string, role string) {
	c.chatHistory = append(c.chatHistory, openai.ChatCompletionMessage{
		Role:    role,
		Content: message,
	})
}

func (c *ChatGPTClient) SendMessage(msg string) (string, error) {
	c.addNewMessageToChatHistory(msg, openai.ChatMessageRoleUser)
	queryToSend := openai.ChatCompletionRequest{
		Model:     openai.GPT4o,
		Messages:  append(c.fixedInput, c.chatHistory...),
		Temperature: 0.1,
	}
	response, err := c.makeChatGPTRequest(queryToSend)
	if(err != nil ) {
		log.Fatalf("Could not get response from chatGPT api, %v", err)
	}
	c.addNewMessageToChatHistory(response, openai.ChatMessageRoleAssistant)
	// fmt.Print(queryToSend, "response is this!" + response)
	return response, nil
}
