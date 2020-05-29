package main

import "fmt"

type I interface {
	abc()
}
type myInt int

func (mi myInt) abc() {
	fmt.Println("Calling abc()", mi)
}

func cal(i I) {
	i.abc()
}

func main() {
	//var i I

	var x myInt
	x = 1
	//i=x
	//i.abc()
	cal(x)
	fmt.Println("Hello", x)
}
