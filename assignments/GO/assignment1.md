##  Findout type of error in Go and how it is defined?


- `For any kind of datatype (primary or secondary) if method named "Error() string" is created (using method feature), now if that object is printed using fmt.Println(), this Error method will be called and return value (string) of this Error() method wll be printed out
error handeling can be done is this Error function()`

- errors is external library used to define simple errors

- For functions as a standard err should be returned as last return value, simply as
`func f2(arg int) (int, error) {`


Example1: 
```
package main
import (
        "fmt"
//        "errors"
       )

type A int
func (a A) Error() string{
 return "Error!"
 //return (fmt.Sprintf("error %d",a))
}

func main(){
  var a A
  a=1
  fmt.Println(a)
//  e:=errors.New("test error")
//  fmt.Println(e)
}
```
