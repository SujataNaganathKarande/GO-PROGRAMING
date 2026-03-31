package main

import "fmt"

func main() {
	var r1, c1, r2, c2 int

	fmt.Print("Enter rows and columns of first matrix: ")
	fmt.Scan(&r1, &c1)

	fmt.Print("Enter rows and columns of second matrix: ")
	fmt.Scan(&r2, &c2)

	if c1 != r2 {
		fmt.Println("Matrix multiplication not possible!")
		return
	}

	m1 := make([][]int, r1)
	m2 := make([][]int, r2)
	result := make([][]int, r1)

	for i := range m1 {
		m1[i] = make([]int, c1)
	}
	for i := range m2 {
		m2[i] = make([]int, c2)
	}
	for i := range result {
		result[i] = make([]int, c2)
	}

	fmt.Println("Enter elements of first matrix:")
	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			fmt.Scan(&m1[i][j])
		}
	}

	fmt.Println("Enter elements of second matrix:")
	for i := 0; i < r2; i++ {
		for j := 0; j < c2; j++ {
			fmt.Scan(&m2[i][j])
		}
	}

	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			for k := 0; k < c1; k++ {
				result[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}

	fmt.Println("Multiplication Result:")
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			fmt.Print(result[i][j], " ")
		}
		fmt.Println()
	}
}
