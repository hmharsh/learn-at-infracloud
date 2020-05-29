package main

import (
	"fmt"
	"os"
	//	"strings"
)

func main() {
	os.Setenv("name", "harshit")
	ee := os.Getenv("name")
	fmt.Println("Hello", ee)
	for _, e := range os.Environ() {

		fmt.Println(e)
		//pair := strings.SplitN(e, "=", 2)
		//fmt.Println(pair[1])
	}
}
