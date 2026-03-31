package main

import "fmt"

// Define the student structure
type student struct {
	rollNo int
	name   string
	grade  string
}

// Method with pointer receiver
func (s *student) show() {
	fmt.Println("Student Details:")
	fmt.Println("Roll No:", s.rollNo)
	fmt.Println("Name   :", s.name)
	fmt.Println("Grade  :", s.grade)
}

func main() {
	// Create a student object
	st := student{
		rollNo: 101,
		name:   "Ram",
		grade:  "A",
	}

	// Call the show method
	st.show()
}
