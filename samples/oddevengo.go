package samples

import (
	"fmt"
	"sync"
)

func printOddNumbers(wg *sync.WaitGroup, oddCh chan<- int) {
	defer wg.Done()
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}
}

func printEvenNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 100; i += 2 {
		fmt.Println(i)
	}
}

func mainOddEven() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printOddNumbers(&wg, make(chan int))
	go printEvenNumbers(&wg)

	wg.Wait()
}
