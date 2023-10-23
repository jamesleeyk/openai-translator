package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeToFile(message string) {
	// Create or open a text file for writing
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close() // Defer closing the file until the function returns.

	// Create a buffered writer for efficient writing
	writer := bufio.NewWriter(file)

	// Write to the file
	_, err = writer.WriteString(message)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}

	// Flush the writer to ensure data is written to the file
	writer.Flush()
}