package main

import "fmt"

func main() {

	// create buffered channel with capacity 5
	ch := make(chan int, 5)

	// store values in channel
	ch <- 10
	ch <- 20
	ch <- 30

	// check capacity and length
	fmt.Println("Channel Capacity:", cap(ch))
	fmt.Println("Channel Length:", len(ch))

	// read values from channel
	fmt.Println("Reading values from channel:")
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// check modified length after reading
	fmt.Println("Modified Channel Length:", len(ch))
}
