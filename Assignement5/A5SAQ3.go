package main

import (
	"fmt"
)

// goroutine to receive and print even numbers
func evenNumbers(ch chan int) {
	for num := range ch {
		fmt.Println("Even number received:", num)
	}
}

// goroutine to receive and print odd numbers
func oddNumbers(ch chan int) {
	for num := range ch {
		fmt.Println("Odd number received:", num)
	}
}

func main() {
	// slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenChan := make(chan int)
	oddChan := make(chan int)

	// start goroutines
	go evenNumbers(evenChan)
	go oddNumbers(oddChan)

	// check numbers and send to respective channels
	for _, num := range numbers {
		if num%2 == 0 {
			evenChan <- num
		} else {
			oddChan <- num
		}
	}

	// close channels
	close(evenChan)
	close(oddChan)

	// small delay to allow goroutines to finish printing
	fmt.Scanln()
}
