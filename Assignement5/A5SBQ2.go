package main

import "fmt"

func main() {

	ch := make(chan int)
	n := 10 // number of Fibonacci terms

	// goroutine to write Fibonacci series to channel
	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			ch <- a
			a, b = b, a+b
		}
		close(ch)
	}()

	// main goroutine reads from channel
	fmt.Println("Fibonacci Series:")
	for num := range ch {
		fmt.Println(num)
	}
}
