package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println("Task1")
	fmt.Println("Task2")
	fmt.Println("Task3")
}
func main() {
	done := make(chan bool)
	t1 := time.NewTicker(2 * time.Second)
	go func() {
		for {
			select {
			case <-t1.C:
				test()
			case <-done:
				return
			}
		}
	}()
	time.Sleep(7 * time.Second)
	t1.Stop()
	done <- true
}
