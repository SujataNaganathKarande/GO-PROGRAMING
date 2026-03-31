package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// File to append content
	filename := "example.txt"

	// Open file in append mode, create if it doesn't exist
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Content to append
	content := "This is a new line added to the file.\n"

	// Write content to file
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	fmt.Println("Content appended successfully.")
}
