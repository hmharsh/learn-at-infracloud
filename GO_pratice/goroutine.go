package main
import (
	"fmt"
	"time"
       )

func say(s string){
for i:=0;i<10;i++{
	fmt.Println(i,":-", s)
}
}
func main(){
  go say("harshit")
  go func(){
	  fmt.Println("hii")
  }()
  time.Sleep(2*time.Second)
}
