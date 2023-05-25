package samples

import (
	"fmt"
	"sync"
)

func mainOrderPrint() {
	const N = 10

	var wg sync.WaitGroup
	wg.Add(3)

	// Channel for sending numbers
	ch := make(chan int)

	// Goroutine to send numbers
	go func() {
		defer wg.Done()
		for i := 1; i <= N; i++ {
			ch <- i
			fmt.Println("Sender 1 sent:", i)
			// Wait for receiver to print the number
			<-ch
		}
	}()

	// Goroutine to send numbers
	go func() {
		defer wg.Done()
		for i := 1; i <= N; i++ {
			// Wait for receiver to print the number
			<-ch
			ch <- i
			fmt.Println("Sender 2 sent:", i)
		}
	}()

	// Goroutine to receive and print numbers
	go func() {
		defer wg.Done()
		for i := 1; i <= N; i++ {
			// Wait for a number to be sent
			num := <-ch
			fmt.Println("Receiver received:", num)
			// Allow the sender to send the next number
			ch <- num
		}
	}()

	wg.Wait()
}
