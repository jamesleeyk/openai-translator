package main

import (
	"fmt"
	"log"
	"os"

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
	chatClient := CreateChatClient(openAiKey)
	chatClient.setFixedInput()
	scannerInstance := NewScannerHolder("input.txt")
	numLinesToRead := 20
	for {
		rawText, err := getInputFromFile(scannerInstance, numLinesToRead)
		if err != nil {
			log.Fatalf("Scanner error: %v\n", err)
		} else if err == nil && rawText == "" {
			fmt.Println("Reached end of file.")
			break;
		}
		response, err := chatClient.SendMessage(rawText)
		if err != nil {
			log.Fatalf("Message Error: %v\n", err)
		}
		// fmt.Printf("scanned text: %s", rawText)
		writeToFile(response)
	}
	fmt.Print("Done!!!")
}

func loadEnv() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
