package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := "ls -a -l -h"
	op, _ := exec.Command("bash", "-c", cmd).Output()
	fmt.Println(string(op))

}
