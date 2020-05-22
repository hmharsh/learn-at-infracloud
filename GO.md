# Installing and setting up GO
```
sudo su
wget https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz
tar -xzf go1.14.3.linux-amd64.tar.gz
cp -r ./go /usr/local/
mkdir -p $HOME/go/{bin,src} // working directory bin for binary and src for code
```

Append the line below in ~/.barshrc (orbashrc of $HOME)
```
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

Run
```
source ~/.barshrc
go version
```

Install git
Extend VIM (optional)
use: https://github.com/junegunn/vim-plug (just curl)
and in vimrc

```
call plug#begin('~/.vim/plugged')

Plug 'fatih/vim-go'

call plug#end()

filetype off
filetype plugin indent on

set number
set noswapfile
set noshowmode
set ts=2 sw=2 sts=2 et
set backspace=indent,eol,start

" Map <leader> to comma
let mapleader=","

if has("autocmd")
  autocmd FileType go set ts=2 sw=2 sts=2 noet nolist autowrite
endif
```

```
vim +PlugInstall +qa
vim +GoInstallBinaries +qa
```

Go  to $HOME/go/src
mkdir helloapp
vim main.go
```
package main
import "fmt"
func main() {
  fmt.Println("Hello harshit")
  fmt.Println("The time is", time.Now())
}

```
go run main.go
go build  
=> https://tour.golang.org/ // al syntax


# comments
/* multi line */
// single line


# complex no
import "math/cmplx"
z      complex128 = cmplx.Sqrt(-5 + 12i)

# print string
fmt.Printf("%q\n",name)

# Type inference
fmt.Printf("v is of type %T\n", v)

'''
Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

The init and post statements are optional in for loop, semicolons can also be dropped to use for as while loop

infinite loop: for {}
'''

IF initialize;condition {
	initialized variable is accessable only from this block and else block
}
```
if v := math.Pow(x, n); v < lim {
		return v
	}
```

float64
```
%g  whichever of %e or %f produces more compact output
%e  scientific notation, e.g. -1234.456e+78
%f  decimal point but no exponent, e.g. 123.456
%T type

all at: https://golang.org/pkg/fmt/
```

# defer
```
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
```

Pointers
&: address of
*: value at 
var p *int // declate pointer type of variable to store address

```
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

```

pointers to struct:  (*p).X = 1e9  or p.X = 1e9  -->(ie9 means 10 to the power 9))


sequence (for both array and slice)
size -> data type -> data
days:= []string {"s","m","t"}
var s []string = days[0:1]
a := days[0:2]

make function, to define dynamic sized slice => a := make([]int, 5)
append function: s = append(s, 0)

# Two values returned by range
    for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	=> i,v or _,v or i,_
# Map
	var m map [string]int
	m=make(map[string]int)
	m["harshit"] = 1
	delete(m, "harshit")
	fmt.Println(m["harshit"])

# word count function
func WordCount(s string) map[string]int {
	var wc map [string]int
	wc=make(map[string]int)
	for _,v := range strings.Fields(s) {
		if _,b := wc[v];b {
			wc[v] = wc[v]+1
		}else{
			wc[v] = 1
		}
	} 
     return wc
}


# Function as Variable 
```
func mult (x float64, y float64) float64 {
	return x*y
}
func cal (a func(float64,float64)float64) float64{
	return (a(3,4))
}
func main(){
    fmt.Println(cal(mult))
}
```



# closure 
- function that return a function  (the inner returning function is called closure)
- function tide up a value
- function hold a annymous function 

```
func outer() func () int {
	a := 0
	return func() int{
		a+=1
		return a
	}
}
```

```
 A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
```

# Method
```
type Vertex struct {
	X float64
	Y float64
}
func (V Vertex) ab() float64{
    return  (V.X * V.Y)
}
func main(){
  v:= Vertex{3,4}
  fmt.Println(v.ab()) 
}
```
by using `type MyFloat float64` methods can be applied to non struct objects too

# Interface
```
An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.
```
``` 
Interface
```
type I interface {
	M()
} 
```
Implimentation: 
```
    // Way: 1
 	var a I
	var h = T{"hello"}
	a=h
	a.M()
	// Way: 2, Implisit
	var i I = T{"hello"}
	i.M()
```


