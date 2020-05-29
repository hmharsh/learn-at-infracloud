package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
	wordPtr := flag.String("word", "foo", "a string")
	flag.Parse()
	fmt.Println("word:", *wordPtr)
}
