package main

import "fmt"

type s struct {
	a int
	b string
	c bool
}

func checks(a *s) bool {
	if a.c == true {
		a.c = false
		return true
	} else {
		return false
	}
}

func create(a s) *s {
	a.b = "hhh"
	return &a
}
func main() {
	a := s{1, "hn", true}
	a.b = "aa"
	z := s{2, "hn", false}
	y := create(z)
	fmt.Println(*y)
	fmt.Println("Hello", checks(&a))
	fmt.Println("Hello", checks(&a))
}
