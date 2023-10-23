package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInputFromFile(fileName string) string {
	// Open the text file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}
	defer file.Close() // Defer closing the file until the function returns.

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	scannedText := "";
	// Read from the file line by line
	for scanner.Scan() {
		scannedText += scanner.Text()
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}
	return scannedText
}
