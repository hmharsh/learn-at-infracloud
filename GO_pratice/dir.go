package main

import "os"

func main() {
	os.Mkdir("./Harshit", 0644)     // create dir
	defer os.RemoveAll("./Harshit") // remove dir
}
