package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, checkpoint chan bool) {
	defer wg.Done()

	fmt.Printf("Worker %d is assembling its part...\n", id)

	// simulate work
	time.Sleep(time.Duration(id) * time.Second)

	fmt.Printf("Worker %d finished work and reached checkpoint\n", id)

	// signal that worker reached checkpoint
	checkpoint <- true

	// wait for checkpoint release
	<-checkpoint

	fmt.Printf("Worker %d continues to next task\n", id)
}

func main() {

	numWorkers := 4
	var wg sync.WaitGroup
	checkpoint := make(chan bool, numWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, checkpoint)
	}

	// wait until all workers reach checkpoint
	for i := 0; i < numWorkers; i++ {
		<-checkpoint
	}

	fmt.Println("\nAll workers reached checkpoint. Assembling parts together...\n")

	// release workers
	for i := 0; i < numWorkers; i++ {
		checkpoint <- true
	}

	wg.Wait()
	fmt.Println("\nAll workers completed the next task.")
}
