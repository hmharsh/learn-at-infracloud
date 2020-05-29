package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()
	p(time.Now())
	p(time.RFC3339)
	p(t.Format("02/01/2006"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)
	p("test", rand.Intn(50))

	s1 := rand.NewSource(time.Now().Unix())
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), "\n")

}
