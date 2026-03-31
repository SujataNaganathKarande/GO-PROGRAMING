package main

import "fmt"

type Employee struct {
	eno    int
	ename  string
	salary float64
}

func main() {
	var n int
	fmt.Print("Enter number of employees: ")
	fmt.Scan(&n)

	employees := make([]Employee, n)

	for i := 0; i < n; i++ {
		fmt.Println("\nEnter details of Employee", i+1)

		fmt.Print("Employee No: ")
		fmt.Scan(&employees[i].eno)

		fmt.Print("Employee Name: ")
		fmt.Scan(&employees[i].ename)

		fmt.Print("Salary: ")
		fmt.Scan(&employees[i].salary)
	}

	maxSalary := employees[0].salary

	for i := 1; i < n; i++ {
		if employees[i].salary > maxSalary {
			maxSalary = employees[i].salary
		}
	}

	fmt.Println("\nEmployee(s) with Maximum Salary:")
	for i := 0; i < n; i++ {
		if employees[i].salary == maxSalary {
			fmt.Println("Employee No:", employees[i].eno)
			fmt.Println("Name:", employees[i].ename)
			fmt.Println("Salary:", employees[i].salary)
			fmt.Println()
		}
	}
}
