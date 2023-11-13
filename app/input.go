package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type ScannerHolder struct {
	*bufio.Scanner
}

func NewScannerHolder(fileName string) *ScannerHolder {
	// Open the text file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}
	// defer file.Close() // Defer closing the file until the function returns.

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	return &ScannerHolder{Scanner: scanner}
}

func getInputFromFile(sh *ScannerHolder, numLines int) (string, error) {
	var lines []string
	for i := 0; i < numLines; i++ {
		if !sh.Scanner.Scan() {
			err := sh.Scanner.Err()
			if err != nil {
			  log.Fatal(sh.Scanner.Err())
			} else {
				lines = append(lines, "")
				return strings.Join(lines, "\n"), err
			}
		}
		text := sh.Scanner.Text()
		// if (i == numLines - 1) && (!strings.HasSuffix(text, ".") || !strings.HasSuffix(text, "!") || !strings.HasSuffix(text, "?")) {
		// 	i--
		// }
		lines = append(lines, text)
	}
	return strings.Join(lines, "\n"), nil
}
