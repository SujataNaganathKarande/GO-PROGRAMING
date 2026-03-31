package main

import "fmt"

// function to send values to channel
func sendData(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // close the channel
}

func main() {

	ch := make(chan int)

	// start goroutine to send data
	go sendData(ch)

	// receive values using for range loop
	fmt.Println("Receiving values from channel:")
	for value := range ch {
		fmt.Println(value)
	}

	fmt.Println("Channel closed")
}
