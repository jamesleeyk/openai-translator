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
	initialPromptString := "You are an English novel writer that writes fantasy fiction. Once you have Korean text, you will translate it into English. Do not add any new information when translating.";
	firstInputPrompt := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: initialPromptString,
	}
	c.fixedInput = append(c.fixedInput, firstInputPrompt)
}

func (c *ChatGPTClient) makeChatGPTRequest(query openai.ChatCompletionRequest) (string, error) {
	contextTimeout, cancel := context.WithTimeout(c.ctx, time.Minute*3)
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
		Model:     openai.GPT4,
		Messages:  append(c.fixedInput, c.chatHistory...),
	}
	response, err := c.makeChatGPTRequest(queryToSend)
	if(err != nil ) {
		log.Fatalf("Could not get response from chatGPT api, %v", err)
	}
	c.addNewMessageToChatHistory(response, openai.ChatMessageRoleAssistant)
	return response, nil
}
