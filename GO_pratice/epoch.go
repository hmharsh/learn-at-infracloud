package main

import (
	"fmt"
	"time"
)

func main() {
	sec := time.Now().Unix()
	fmt.Println(time.Unix(sec, 0))
}
