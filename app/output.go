package main

import (
	"fmt"
	"os"
)

func writeToFile(message string) {
	// Write the string to a file called output.txt
	// create the file if it does not exist
	// else, append to the file
	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(message + "\n")
	if err2 != nil {
		fmt.Println(err2)
	}
}