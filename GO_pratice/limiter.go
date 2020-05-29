package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(1 * time.Second)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}
