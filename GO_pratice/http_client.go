package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://gobyexample.com")
	fmt.Println("Hello")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println(res)

	scanner := bufio.NewScanner(res.Body)
	for i := 0; scanner.Scan() && i < 5; i++ { // 1st 5 lines of o/p
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
