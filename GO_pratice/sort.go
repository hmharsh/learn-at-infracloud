package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"ab", "aa", "d", "z"}
	sort.Strings(str)
	fmt.Println(str)

	list := []int{1, 2, 3}
	sort.Ints(list)
	fmt.Println(list)
}
