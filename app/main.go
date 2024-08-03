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
	glossaryScanner := NewScannerHolder("glossary.txt")
	scannerInstance := NewScannerHolder("input.txt")
	numLinesToReadGlossary := 30
	numLinesToRead := 50
	for {
		rawText, err := getInputFromFile(glossaryScanner, numLinesToReadGlossary)
		if err != nil {
			log.Fatalf("Scanner error: %v\n", err)
		} else if err == nil && rawText == "" {
			fmt.Println("Scanned glossary.")
			break;
		}
		_, err = chatClient.SendMessage(rawText)
		if err != nil {
			log.Fatalf("Message Error: %v\n", err)
		}
	}
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
