package main
import "fmt"
func main(){
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "ab"
	s[2] = "abcd"
	s = append(s,"x")
	//	s[3] = "abc"
	fmt.Println("Hello",s)
}
