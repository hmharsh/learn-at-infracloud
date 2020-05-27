package main
import "fmt"
func fact(a int) int{
	
	if  a == 1 {
	 return a
	}else {

		 a := a*fact(a-1)
		 return a
	}
}
func main(){
  fmt.Println("Hello")
  b:= fact(5)
  fmt.Println(b)
}
