package main

import "fmt"

func worker(id int, job <-chan int, result chan<- int) {
	for r := range job {
		r = r * 2
		result <- r
	}
}
func main() {
	fmt.Println("Hello")
	job := make(chan int, 5)
	result := make(chan int, 5)
	for i := 0; i < 3; i++ {
		go worker(i, job, result)
	}
	for j := 0; j < 2; j++ {
		job <- j
	}
	close(job)
	for r := range result {
		//for j := 0; j < 2; j++ {
		fmt.Println(r)
	}

	close(result)
}
