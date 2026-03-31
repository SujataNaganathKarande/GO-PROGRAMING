package main

import (
	"fmt"
)

// function to calculate sum of squares
func sumOfSquares(num int, ch chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	ch <- sum
}

// function to calculate sum of cubes
func sumOfCubes(num int, ch chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	ch <- sum
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	sqChan := make(chan int)
	cubeChan := make(chan int)

	// start goroutines
	go sumOfSquares(number, sqChan)
	go sumOfCubes(number, cubeChan)

	// receive results
	squares := <-sqChan
	cubes := <-cubeChan

	fmt.Println("Sum of squares =", squares)
	fmt.Println("Sum of cubes =", cubes)
	fmt.Println("Final sum of squares and cubes =", squares+cubes)
}
