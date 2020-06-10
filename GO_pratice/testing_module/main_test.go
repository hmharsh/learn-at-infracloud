package main

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func ExampleHello() {
	fmt.Println("hello harshit")
	fmt.Println("hello harshit")
	// Output:
	// hello harshit
	// hello harshit
}

func ExamplePerm() {
	for _, value := range rand.Perm(5) {
		fmt.Println(value)
	}
	// Unordered output:
	// 4
	// 2
	// 1
	// 3
	// 0
}

func TestMain(m *testing.M) {
	// just like main function, but for testing
	// call flag.Parse() here if TestMain uses flags

	fmt.Println("run here if something is required to run before the test")
	os.Exit(m.Run()) // this will call the testing
}

func TestCal(t *testing.T) {

	if testing.Short() { // true only if go test -short flag is supplied
		t.Skip("skipping test in short mode.") // skip function can be used based on condition
	}

	if Cal(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
	if Cal(3) != 5 {
		t.Error("Expected 2 + 3 to equal 5")
	}
}

func TestMultidata(t *testing.T) {
	type testData struct {
		input    int
		expacted int
	}
	x := []testData{
		{-10, -8},
		{-1, 1},
		{0, 2},
		{4, 6},
	}
	for _, data := range x {
		if Cal(data.input) != data.expacted {
			t.Error("Expected 2 + ", data.input, " to equal", data.expacted)
		}
	}

}

func TestCalNew(t *testing.T) {
	// like this from just 1 function we can call many test function,call Run("test name",function(t *testing T))
	//t.Run("test1", func(t *testing.T) { t.Error("Test") })
}

// Simple benchmark function
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
