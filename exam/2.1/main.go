package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("Send message: %d\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
