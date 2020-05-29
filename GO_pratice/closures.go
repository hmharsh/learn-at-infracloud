package main

import "fmt"

func outer() func() int { // will return a func() int
	var a int
	a = 0
	return func() int {
		a += 1
		return a
	}
}

func main() {
	count := outer() // will initialise a
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
}
