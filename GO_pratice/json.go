package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := map[string]int{"h": 1, "m": 2}
	j, _ := json.Marshal(x)
	fmt.Println(string(j))
}
