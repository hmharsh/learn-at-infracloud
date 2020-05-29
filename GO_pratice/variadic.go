package main

import "fmt"

func total(a ...int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}
func main() {
	fmt.Println("Hello")
	sli := []int{1, 2, 3}

	fmt.Println(sli)
	fmt.Println(total(sli...))
}
