package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	t1 := time.NewTimer(2 * time.Second)
	x := t1.Stop()
	//	<-t1.C
	fmt.Println(x)
}
