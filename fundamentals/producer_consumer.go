package main

import (
	"fmt"
	"sync"
)

func main() {

	producerConsumerProgramPattern()
}

func producerConsumerProgramPattern() {
	// Create a channel to communicate between producer and consumer
	ch := make(chan int)

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start the producer goroutine
	wg.Add(1)
	go producer(ch, &wg)

	// Start the consumer goroutine
	wg.Add(1)
	go consumer(ch, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
}

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		// Send the value to the channel
		ch <- i
	}

	// Close the channel to signal the consumer that no more values will be sent
	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		// Process the value received from the channel
		fmt.Println("Received:", num)
	}
}
