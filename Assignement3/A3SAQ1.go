package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Enter elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	largest := arr[0]
	smallest := arr[0]

	for i := 1; i < n; i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}

	fmt.Println("Largest number:", largest)
	fmt.Println("Smallest number:", smallest)
}
