package main

import "fmt"

func test(a string) (x string, y int) {
	x = a
	y = 2
	return
}
func main() {
	x, y := test("harshit")
	fmt.Println(x, y)
}
