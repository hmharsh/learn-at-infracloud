// Go's _structs_ are typed collections of fields.
// They're useful for grouping data together to form
// records.

package main

import "fmt"

// This `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}
type all struct {
   persons person 
   count   int
}

// `newPerson` constructs a new person struct with the given name.
func newPerson(name string) *person {
	// You can safely return a pointer to local variable
	// as a local variable will survive the scope of the function.
	p := person{name: name}
	p.age = 42
	return &p
}

func newPersons(name string) *all{
     s := all{ person{"Bob", 20}, 12 }
     return &s
}



func array() *[]person{
     s :=make([]person, 1)
     s[0] = person{"Bob", 20}
     return  &s
}

func main() {

	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	

	fmt.Println(newPersons("abc"))
	fmt.Println(array())
	
	

  
}
