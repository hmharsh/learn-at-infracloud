package main

import (
	"fmt"
	//	"errors"
)

type A int

func (a A) Error() string {
	//return  "hh"
	return (fmt.Sprintf("error %d", a))
}

func main() {
	var a A
	a = 1
	fmt.Println(a)
	//  e:=errors.New("test error")
	//  fmt.Println(e)
}
