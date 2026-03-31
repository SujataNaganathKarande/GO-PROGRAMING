package main

import (
	"fmt"
	"sort"
)

// Define the student structure
type Student struct {
	rollNo     int
	name       string
	percentage float64
}

// Method to display student information
func (s Student) show() {
	fmt.Printf("Roll No: %d, Name: %s, Percentage: %.2f\n", s.rollNo, s.name, s.percentage)
}

// Method to display all students in descending order of percentage
func showDescending(students []Student) {
	// Sort slice in descending order of percentage
	sort.Slice(students, func(i, j int) bool {
		return students[i].percentage > students[j].percentage
	})

	fmt.Println("\nStudents in Descending Order of Percentage:")
	for _, s := range students {
		s.show()
	}
}

func main() {
	var n int
	fmt.Print("Enter number of students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	// Input student details
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for student %d:\n", i+1)
		fmt.Print("Roll No: ")
		fmt.Scan(&students[i].rollNo)
		fmt.Print("Name: ")
		fmt.Scan(&students[i].name)
		fmt.Print("Percentage: ")
		fmt.Scan(&students[i].percentage)
	}

	// Display students sorted by percentage
	showDescending(students)
}
