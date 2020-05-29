package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	p("Hello")
	//	now := time.Now()

	// 2020-05-28 17:53:09.8811694 +0530 IST m=+0.001759601
	//want (int, time.Month, int, int, int, int, int, *time.Location)
	// 2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

	/*
	   p(then.Year())
	   p(then.Month())
	   p(then.Day())
	   p(then.Hour())
	   p(then.Minute())
	   p(then.Second())
	   p(then.Nanosecond())
	   p(then.Location())

	*/
	than1 := time.Date(2020, 05, 28, 14, 53, 9, 8811694, time.UTC)
	than2 := time.Date(2020, 05, 28, 15, 53, 9, 8811694, time.UTC)
	p(than2)
	p(than1)
	diff := than1.Sub(than2)
	fmt.Println(diff)
}
