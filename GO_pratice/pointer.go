package main
import "fmt"
func main(){
  var i int
   i=10
   p:= &i

  fmt.Println("Hello",i,p,*p,2**p)
}
