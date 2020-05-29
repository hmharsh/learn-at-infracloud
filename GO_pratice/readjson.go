package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type personal_data struct {
	Name string // should start with capital latter
	City string
}

func main() {
	dat, err := ioutil.ReadFile("./data.json")
	check(err)

	//fmt.Println([]byte(string(dat))) bat is in byte formate and this is how type conversion can be done from byte to string and from string to byte
	errr := ioutil.WriteFile("./replica.json", dat, 0644)
	check(errr)
	var person_dataa []personal_data
	json.Unmarshal(dat, &person_dataa)
	//	fmt.Println(person_data)
	fmt.Printf("1st name from json file : %+v\n", person_dataa[0].Name)
}
