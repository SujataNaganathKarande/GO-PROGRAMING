package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Define struct matching XML
type Student struct {
	Name  string `xml:"name"`
	Marks int    `xml:"marks"`
}

type Students struct {
	Students []Student `xml:"student"`
}

func main() {
	// Open XML file
	xmlFile, err := os.Open("students.xml")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer xmlFile.Close()

	// Read file content
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Unmarshal XML into struct
	var students Students
	err = xml.Unmarshal(byteValue, &students)
	if err != nil {
		log.Fatal("Error unmarshaling XML:", err)
	}

	// Display structure
	for _, s := range students.Students {
		fmt.Printf("Name: %s, Marks: %d\n", s.Name, s.Marks)
	}
}
