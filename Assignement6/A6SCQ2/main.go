package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// File to inspect
	filename := "example.txt"

	// Open the file
	fileInfo, err := os.Stat(filename)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Print file information
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("Size (bytes):", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Println("Last Modified:", fileInfo.ModTime())
}
