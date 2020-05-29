package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	csv_file := "./data.csv"
	score := 0
	dat, err := ioutil.ReadFile(csv_file)
	check(err)
	reader := bufio.NewReader(os.Stdin)
	r := csv.NewReader(strings.NewReader(string(dat)))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("what is:", record[0], " ?")
		answer, err := reader.ReadString('\n')
		check(err)
		answer = strings.Replace(answer, "\n", "", -1)
		if answer != record[1] {
			fmt.Println("\nwrong Answer\nYour score is:", score)
			return
		}
		score++
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
