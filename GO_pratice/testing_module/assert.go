// https://gist.github.com/samalba/6059502

package main

import (
        "fmt"
        "testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
        if a == b {
                return
        }
        if len(message) == 0 {
                message = fmt.Sprintf("%v != %v", a, b)
        }
        t.Fatal(message)
}

func TestSimple(t *testing.T) {
        a := 42
        assertEqual(t, a, 42, "")                                                                                                 assertEqual(t, a, 42, "This message is displayed in place of the default one")
}
