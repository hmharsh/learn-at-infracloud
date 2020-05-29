package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csv_file := "./data.csv"
	score := 0
	time_limit_ptr := flag.Int("a", 30, "time limit in seconds")
	flag.Parse()
	dat, err := ioutil.ReadFile(csv_file)
	check(err)
	reader := bufio.NewReader(os.Stdin)
	r := csv.NewReader(strings.NewReader(string(dat)))
	timerExceed := time.After(time.Duration(*time_limit_ptr) * time.Second)
	//<-timerExceed

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("what is:", record[0], " ?")
		answerChan := make(chan string)
		go func() {
			answer, err := reader.ReadString('\n')
			check(err)
			answer = strings.Replace(answer, "\n", "", -1)
			answerChan <- answer
		}()

		select {
		case <-timerExceed:
			{

				fmt.Println("\nTime limit exceeded\nYour score is:", score)
				return
			}
		case answer := <-answerChan:
			{

				if answer != record[1] {
					fmt.Println("\nwrong Answer\nYour score is:", score)
					return
				}
				score++

			}
		}

	}

	fmt.Println("\nYou won\nYour score is:", score)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
