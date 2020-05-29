package main

import "fmt"

func main() {
	message := make(chan string)
	go func(a string) {
		message <- a
		message <- "ab"
	}("hii i am harshit")
	msg1 := <-message
	fmt.Println(msg1)
	msg2 := <-message
	fmt.Println(msg2)
}
