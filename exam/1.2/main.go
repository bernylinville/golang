package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)

	i := 0

	// consume
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			time.Sleep(5 * time.Second)
			fmt.Printf("Send message: %d\n", <-messages)
		}
	}()

	// produce
	for i < 10 {
		messages <- i
		i++
	}

	time.Sleep(10 * time.Second)
}
