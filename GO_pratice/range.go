package main

import "fmt"

func main() {
	//  arr:=[10]int {1,2,3,4,5,6,7,8,9,0}
	//  for i,v := range arr{

	// fmt.Println("index", i,"  value:", v)
	//  }
	m := make(map[string]int)
	m["harshit"] = 2
	m["mahajan"] = 3
	m["hm"] = 4
	for i, v := range m {
		fmt.Println(i, v)
	}
	fmt.Println("Hello", m)

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
