package main
import (
	"fmt"
	"time"
)



func send(a chan<-string){ // to specify direction <-,  simply it can also be, func send(a chan string){
  msg:="message"
  a <- msg
}


func recive(a <-chan string){ // arrow is before chan <-chan string
	msg := <-a
        fmt.Println(msg)
}



func main(){
  b:= make(chan string)
  go  send(b)
  go  recive(b)
  time.Sleep(time.Second)
}
