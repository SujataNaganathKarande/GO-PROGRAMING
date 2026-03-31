package main

import "fmt"

type Student struct {
	roll_no   int
	stud_name string
	mark1     float64
	mark2     float64
	mark3     float64
	total     float64
	average   float64
}

func main() {

	var n int
	fmt.Print("Enter number of students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		fmt.Println("\nEnter details for Student", i+1)

		fmt.Print("Roll No: ")
		fmt.Scan(&students[i].roll_no)

		fmt.Print("Name: ")
		fmt.Scan(&students[i].stud_name)

		fmt.Print("Mark1: ")
		fmt.Scan(&students[i].mark1)

		fmt.Print("Mark2: ")
		fmt.Scan(&students[i].mark2)

		fmt.Print("Mark3: ")
		fmt.Scan(&students[i].mark3)

		students[i].total = students[i].mark1 + students[i].mark2 + students[i].mark3
		students[i].average = students[i].total / 3
	}

	fmt.Println("\nStudent Details:")

	for i := 0; i < n; i++ {
		fmt.Println("\nRoll No:", students[i].roll_no)
		fmt.Println("Name:", students[i].stud_name)
		fmt.Println("Total:", students[i].total)
		fmt.Println("Average:", students[i].average)
	}
}
