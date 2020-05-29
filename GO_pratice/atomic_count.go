package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var count uint64
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			for j := 1; j <= 10; j++ {
				atomic.AddUint64(&count, 1)
			}

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
