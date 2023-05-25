package samples

import (
	"fmt"
	"sync"
)

func maingroutine() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	ch := make(chan int)

	wg.Add(2)

	// Goroutine 1: Increment counter
	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			mutex.Lock()
			ch <- i // Send value to channel
			mutex.Unlock()
		}
	}()

	// Goroutine 2: Print counter values
	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			mutex.Lock()
			value := <-ch // Receive value from channel
			fmt.Println("Counter:", value)
			mutex.Unlock()
		}
	}()

	wg.Wait()
	close(ch)
}
