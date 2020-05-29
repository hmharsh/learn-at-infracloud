package main

import (
	"fmt"
	"sync"
	"time"
)

type age struct {
	years int
	mux   sync.Mutex
}

func (a *age) inc() {
	a.mux.Lock()
	a.years = a.years + 1
	a.mux.Unlock()
}
func main() {
	fmt.Println("Hello")
	me := age{years: 1}

	fmt.Println(me)
	for i := 0; i < 20; i++ {
		go me.inc()
	}
	time.Sleep(time.Second)
	fmt.Println(me)

}
