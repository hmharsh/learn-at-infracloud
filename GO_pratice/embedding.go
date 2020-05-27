package main
import "fmt"
type car struct {
	name string
	price int
}
type mycar struct{
	car
	parking  string 
}
func (c car) showcar(){
	fmt.Println("here is the car with name:",c.name," and price:",c.price)
}
func main(){
  c:=car{"alto",4}
  m:=mycar{car{"i10",5},"park"}
  fmt.Println(c)
  fmt.Println(m)
  m.car.showcar()
  m.showcar()
//  c.showcar()
}
