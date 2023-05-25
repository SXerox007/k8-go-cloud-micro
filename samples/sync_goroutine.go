package samples

import (
	"fmt"
	"sync"
)

func sendNumbers1(numCh chan<- int, doneCh1 <-chan bool, doneCh2 chan<- bool, N int) {
	for i := 1; i <= N; i += 2 {
		<-doneCh1       // Wait for the signal from the other goroutine
		numCh <- i      // 1,3,5,7,9
		doneCh2 <- true // Signal completion to the other goroutine
	}

	close(doneCh2) // Close doneCh2 to signal completion to the other goroutine
}

func sendNumbers2(numCh chan<- int, doneCh1 chan<- bool, doneCh2 <-chan bool, N int) {
	for i := 2; i <= N; i += 2 {
		<-doneCh2       // Wait for the signal from the other goroutine
		numCh <- i      // 2,4,6,8,10
		doneCh1 <- true // Signal completion to the other goroutine
	}

	close(doneCh1) // Close doneCh1 to signal completion to the other goroutine
}

func displayNumbers(numCh <-chan int, wg *sync.WaitGroup, N int) {
	defer wg.Done()
	for i := 1; i <= N; i++ {
		num := <-numCh
		fmt.Println(num)
	}
}

func mainSync() {
	N := 10

	numCh := make(chan int)
	doneCh1 := make(chan bool)
	doneCh2 := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)

	go sendNumbers1(numCh, doneCh1, doneCh2, N)
	go sendNumbers2(numCh, doneCh1, doneCh2, N)
	go displayNumbers(numCh, &wg, N)

	doneCh1 <- true // Start the synchronization

	wg.Wait() // Wait for all goroutines to finish
	close(numCh)
	fmt.Println("All numbers displayed")
}
