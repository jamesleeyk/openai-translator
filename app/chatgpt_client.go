// Used under MIT license from risafj/chat-stream

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sashabaranov/go-openai"
)

type ChatGPTClient struct {
	Client              *openai.Client
	ctx                 context.Context
	maxTokensPerMessage int
	maxContext          int
	messages            []openai.ChatCompletionMessage
	fixedInput			[]openai.ChatCompletionMessage
}

func CreateChatClient(apiKey string) *ChatGPTClient {
	return &ChatGPTClient{
		Client:              openai.NewClient(apiKey),
		ctx:                 context.Background(),
		maxContext:          32000,
		maxTokensPerMessage: 5000,
	}
}

func (c *ChatGPTClient) SendMessage(msg string) (string, error) {
	st := time.Now()
	fmt.Print("Sending message")
	c.addMessageToMessages(msg, openai.ChatMessageRoleUser)
	fmt.Printf("Time to add message to array: %d", time.Since(st) / 10000)
	queryToSend := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		Messages:  c.messages,
	}
	response, err := c.makeChatGPTRequest(queryToSend)
	fmt.Printf("Time to get response: %d", time.Since(st) / 10000)
	if(err != nil) {
		log.Fatalf("Could not get response from chatGPT api, %v", err)
	}
	c.addMessageToMessages(response, openai.ChatMessageRoleAssistant)
	return response, nil
}

func (c *ChatGPTClient) makeChatGPTRequest(query openai.ChatCompletionRequest) (string, error) {
	res, err := c.Client.CreateChatCompletion(c.ctx, query)
	if err != nil {
		log.Printf("ChatCompletion Error: %v\n", err)
		return "", err
	}
	return res.Choices[0].Message.Content, nil
}

func (c *ChatGPTClient) setFixedInput() {
	initialPromptString := `You are an English novel writer that writes fantasy fiction.
	
	Once you have Korean text, you will translate it into English, preserving as much white space and as many line breaks as possible.`;
	// referenceText := getInputFromFile("sample_translation.txt")
	firstInputPrompt := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: initialPromptString,
	}
	c.fixedInput = append(c.fixedInput, firstInputPrompt)
	// request := openai.ChatCompletionRequest{
	// 	Model:     openai.GPT3Dot5Turbo16K,
	// 	Messages:   []openai.ChatCompletionMessage{firstInputPrompt},
	// };
	// response, err := c.makeChatGPTRequest(request)
	// if (err != nil) {
	// 	log.Fatalf("Could not get response from chatGPT api, %v", err)
	// }
	// c.addMessageToMessages(response, openai.ChatMessageRoleAssistant)
}

func (c *ChatGPTClient) addMessageToMessages(message string, role string) {
	// Add all existing tokens in message content
	// var totalTokens int
	// for _, msg := range c.messages {
	// 	totalTokens += len(msg.Content)
	// }
	// totalTokens += len(message)

	// // if totalTokens is greater than maxContext - maxTokensPerMessage
	// // remove the first message
	// for totalTokens > c.maxContext-c.maxTokensPerMessage {
	// 	removedMessageTokenCount := len(c.messages[0].Content)
	// 	c.messages = c.messages[1:]
	// 	totalTokens -= removedMessageTokenCount
	// }

	c.messages = append(c.messages, openai.ChatCompletionMessage{
		Role:    role,
		Content: message,
	})
}
