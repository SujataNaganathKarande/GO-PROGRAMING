package main

import "fmt"

func main() {

	numbers := []int{
		10,
		20,
		30,
		40,
		50,
	}

	fmt.Println("Slice Elements:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
