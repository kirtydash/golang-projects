package main

import (
	"fmt"
	"sync"
)

func run(ch chan int, wg *sync.WaitGroup) {
   	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes

	for value := range ch {
		fmt.Println("message received", value)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	ch := make(chan int, 5)

	var wg sync.WaitGroup

	for i := 0; i < len(a); i++ {
		ch <- a[i]
	}

	wg.Add(1)

	go run(ch,&wg)

	close(ch)
	

	wg.Wait()

	fmt.Println("Waiting for goroutine to finish...")
}

