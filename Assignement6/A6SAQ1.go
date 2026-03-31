package main

import (
	"fmt"
	"sort"
)

// Define Student struct
type Student struct {
	Name  string
	Marks int
}

func main() {

	// Create slice of students
	students := []Student{
		{"Amit", 78},
		{"Riya", 92},
		{"Rahul", 65},
		{"Sneha", 88},
	}

	// Sort students based on Marks
	sort.Slice(students, func(i, j int) bool {
		return students[i].Marks < students[j].Marks
	})

	fmt.Println("Students sorted by marks:")

	for _, s := range students {
		fmt.Println(s.Name, s.Marks)
	}
}
