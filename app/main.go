package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Env vars
var openAiApiKey = "OPENAI_API_KEY"

func main() {
	loadEnv()
	
	
	openAiKey := os.Getenv(openAiApiKey)
	if openAiKey == "" {
		log.Fatalf("%s not set", openAiApiKey)
	}
	st := time.Now()
	chatClient := CreateChatClient(openAiKey)
	fmt.Print("Reading input from file")
	allRawText := getInputFromFile("input.txt")
	fmt.Printf("Time to read input: %d", time.Since(st) / 10000)
	chatClient.setFixedInput()
	fmt.Printf("Time to set fixed input: %d", time.Since(st) / 10000)
	for {
		// parse raw text 1000 char
		// send into chat gpt
		response, err := chatClient.SendMessage(allRawText)
		if err != nil {
			fmt.Printf("ChatCompletionStream Error: %v\n", err)
			break
		}
		writeToFile(response)
		break;
	}
}

func loadEnv() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
