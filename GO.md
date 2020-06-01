# Installing and setting up GO
CP: (https://www.codementor.io/@tucnak/using-golang-for-competitive-programming-h8lhvxzt3)
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

Whatever the way interface is implemented, Interface values can be thought of as a tuple of a value and a concrete type: 
` as (value, type) eg: (3.141592653589793, main.F) `


# Type switches
```
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %T %v is %v\n", v,v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```


# Stringers
```
func (p person) String() string{
   return (fmt.Sprintf("name: %v and age %v",p.name,p.age))  //Sprintf for string print
}
```

# Error Handeling 

==> For any kind of datatype (primary or secondary) if method named "Error() string" is created (using method feature), now if that object is printed using fmt.Println(), this Error method will be called and return value (string) of this Error() method wll be printed out
error handeling can be done is this Error function()
Example:
```
package main
import (
        "fmt"
//        "errors"
       )

type A int
func (a A) Error() string{
return "Eerroorr"
}

func main(){
  var a A
  a=1
  fmt.Println(a)
//  e:=errors.New("test error")
//  fmt.Println(e)
}
```


using something like "err != nil" 
```
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

## Example (also refer: https://medium.com/rungo/error-handling-in-go-f0125de052f0)
```
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64


func (e ErrNegativeSqrt) Error() string{
    v:=fmt.Sprint(float64(e))
	return (fmt.Sprintf("cannot Sqrt negative number: %v", v))
}


func Sqrt(x float64) (float64, error) {
    if(x<0){
	 y:=ErrNegativeSqrt(x)
	 return 0,y 
	}
    var z float64 = 1.0
    for i:= 0; i<10 ; i++ {
    z -= (z*z - x) / (2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	/*i, err := Sqrt(-2)
	if(err !=nil) {
	  fmt.Printf(err)
	  return
	}	
	fmt.Println(i)*/
}

```

# Reader
```
    r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	fmt.Printf("%v\n",b)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
```	

# goroutine 
## To start a new thread, use go prefix with specific line of function call 
    go f(x, y, z)

# chennels
 	ch := make(chan int)
	 ch <- v  
	 v := <-ch  
## Example
```
package main
import ("fmt")
func put(i int, ch chan int){
	ch <- i
}
func get(ch chan int) int{
	v := <-ch 
	return v
}
func main(){
	ch := make(chan int)
	go put(25,ch)
	v:=get(ch)
	fmt.Println(v)
}
```
## channel example 2
```
package main
import "fmt"
func main(){
        message := make(chan string)
        go func(a string){
        message <- a
        message <- "ab"
        }("hii i am harshit")
        msg1 := <-message
        fmt.Println(msg1)
        msg2 := <-message
        fmt.Println(msg2)
}
```

## using channel for synchronization 
 https://gobyexample.com/channel-synchronization


## Chennel with buffer
```
func main(){
	ch := make(chan int, 2)	
	ch <- 10
	ch <-  20
	fmt.Println( <-ch)
	fmt.Println( <-ch)	
}
```

## closing chennal is imp
  close(ch) from senders side to let reciver (specially range loop) know
  https://tour.golang.org/concurrency/4 
  `v, ok := <-ch`
  Ok will be false if chennal is closed

```

## ranging on channel
range will return only value, unlike array (index,value)
```
// queue is name of channel
for elem := range queue {
		fmt.Println(elem)
	}
```
func main(){
	ch := make(chan int, 2)	
	ch <- 10
	close(ch)
	//ch <-  20
	fmt.Println( <-ch)
	x,OK := <-ch
	fmt.Println(x,OK)
}
```


# direction with channel
```
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
```
# Select
default section can also be used with select: https://tour.golang.org/concurrency/6
```
package main
import ("fmt")

func main(){
n:=5
ch := make(chan int)	
go say(n,"h", ch)
for i:=0; i<=n;i++{
select {
case msg1:= <- ch:
	fmt.Println(msg1)
   }// select
 } // for loop
} // main function
func say(n int,s string, ch chan int){
 for i:=0;i<=n;i++{
    ch <- i
 }
}
```

## Make select to timeout 
after specific timeout `Add extra case`
either use timeout or use default case for Non-Blocking Channel Operations (to proceed further even if no data is recived)


```
select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }
```
# Mutex
lock or unlock any specific data structure to be updated 
simple way to implement Mutex
```
package main

import (
        "fmt"
        "sync"
        "time"
)

type age struct {
        years int
        mux   sync.Mutex
}

func (a *age) inc() {
        a.mux.Lock()
        a.years = a.years + 1
        a.mux.Unlock()
}
func main() {
        fmt.Println("Hello")
        me := age{years: 1}

        fmt.Println(me)
        for i := 0; i < 20; i++ {
                go me.inc()
        }
        time.Sleep(time.Second)
        fmt.Println(me)

}
```
# Input

```
package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")
	name,_  := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	fmt.Printf("\n your name is %v, hii \n", name)
}
```

# Write to a file
Different Approaches like: ioutil, os.openfile and using writefile function
ioutil one is:
```
package main
import (
	"io/ioutil"
	"log"
)
func main(){
	var data_to_write string
	data_to_write = "hello this is harshit mahajan"
	err := ioutil.WriteFile("/etc/hhhhh", []byte(data_to_write), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
```

# Custom exit code
  os.Exit(0)

# CLI options 
accepting command line arguments

```
package main
import (
	"fmt"
	"flag"
	"os"
)
func main(){
	// declaring variables
	var name string
	// defining flags
	flag.StringVar(&name,"my_name","harshit","name to be used")   // now call this go code with --my_name=... ,similarly other data type can be used like boolVar
	//read from cli and copy into variable whose address is &address
	flag.Parse()
	// if wrong input is provided, in condition check for valid argument
	if(false){
	flag.Usage()
	os.Exit(1)
	}
	// conditionally read from stdin: read input conditinally
	// print out the input
	fmt.Println(name)
}
```



# Install external packages
go get github.com/spf13/cobra
go get -u github.com/spf13/cobra/cobra (also intstall binary if present)
-u : also installed dependencies of external library

go get: to install somebody else code
go install: package all code in a binary under bin directory(for own code)


# build for windows
GOOS=windows go build // generate exe file in current dir, GOOS and GOARCH can be defined
(some values: https://golang.org/cmd/go/#hdr-Environment_variables)
(all values: https://golang.org/doc/install/source#environment)
darwin is for mac os
# cobra 
is a framework to scaffold package structure 
cobra init --pkg-name github.com/hmarsh/myapp myapp
github.com/hmarsh/myapp: path after src
myapp: relative path of where to generate app directory 

# embedding types: 
  concept => https://travix.io/type-embedding-in-go-ba40dd4264df


# Timer
```
func main() {
        fmt.Println(time.Now())
        t1 := time.NewTimer(2 * time.Second)
        x := t1.Stop()
        //      <-t1.C
        fmt.Println(x)
}
```

# Ticker
```
func test() {
        fmt.Println("Task1")
        fmt.Println("Task2")
        fmt.Println("Task3")
}
func main() {
        done := make(chan bool)
        t1 := time.NewTicker(2 * time.Second)
        go func() {
                for {
                        select {
                        case <-t1.C:
                                test()
                        case <-done:
                                return
                        }
                }
        }()
        time.Sleep(7 * time.Second)
        t1.Stop()
        done <- true
}
```


# Worker
```
func worker(id int, job <-chan int, result chan<- int) {
        for r := range job {
                r = r * 2
                result <- r
        }
}
func main() {
        fmt.Println("Hello")
        job := make(chan int, 5)
        result := make(chan int, 5)
        for i := 0; i < 3; i++ {
                go worker(i, job, result)
        }
        for j := 0; j < 2; j++ {
                job <- j
        }
        close(job)
        for r := range result {
                //for j := 0; j < 2; j++ {
                fmt.Println(r)
        }

        close(result)
}
```

# Limiter
```
func main() {

        requests := make(chan int, 5)
        for i := 1; i <= 5; i++ {
                requests <- i
        }
        close(requests)

        limiter := time.Tick(1 * time.Second)

        for req := range requests {
                <-limiter
                fmt.Println("request", req, time.Now())
        }
}
```


# Wait group and Atomic counter

```
package main
import (
        "fmt"
        "sync"
        "sync/atomic"
)
func main() {
        var wg sync.WaitGroup
        var count uint64
        for i := 1; i <= 100; i++ {
                wg.Add(1)
                go func() {
                        for j := 1; j <= 10; j++ {
                                atomic.AddUint64(&count, 1)
                        }

                        wg.Done()
                }()
        }
        wg.Wait()
        fmt.Println(count)
}
```

# Stateful goroutines
to synchronize access to shared state across multiple goroutines. its Another option than using mutex
-  This channel-based approach aligns with Go’s ideas of sharing memory 
   by communicating and having each piece of data owned by exactly 1 goroutine.
   (use by implementing channel in struct)


# sorting
```
package main

import (
        "fmt"
        "sort"
)

func main() {
        str := []string{"ab", "aa", "d", "z"}
        sort.Strings(str)
        fmt.Println(str)

        list := []int{1, 2, 3}
        sort.Ints(list)
        fmt.Println(list)
}
```
For custom sort logic (like sort by lengh of string etc) implement the Len, Swap and Less method with that customised data type, 
the actual comparision logic will reside under Less (https://gobyexample.com/sorting-by-functions)



# Panic
generate panic by
`panic(err)`

# collection functions:
https://gobyexample.com/collection-functions

# string functions:
https://gobyexample.com/string-functions

# Fprintf()
 You can format+print to io.Writers other than os.Stdout using Fprintf.
 e.g.: fmt.Fprintf(os.Stderr, "an %s\n", "error")

# regexp library
 match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
 https://gobyexample.com/regular-expressions

 # JSON
 import ("encoding/json")

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
# Time
https://gobyexample.com/time
```
package main

import (
        "fmt"
        "time"
)
func main() {
        p := fmt.Println
        p("Hello")
        //      now := time.Now()

        // 2020-05-28 17:53:09.8811694 +0530 IST m=+0.001759601
        //want (int, time.Month, int, int, int, int, int, *time.Location)
        // 2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

        /*
           p(then.Year())
           p(then.Month())
           p(then.Day())
           p(then.Hour())
           p(then.Minute())
           p(then.Second())
           p(then.Nanosecond())
           p(then.Location())

        */
		/*
		fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())
		*/
        than1 := time.Date(2020, 05, 28, 14, 53, 9, 8811694, time.UTC)
        than2 := time.Date(2020, 05, 28, 15, 53, 9, 8811694, time.UTC)
        p(than2)
        p(than1)
        diff := than1.Sub(than2)
        fmt.Println(diff)
}
```

time -> epoch 
```
  now := time.Now()
  secs := now.Unix()  // time.now().Unix()
  nanos := now.UnixNano()
```
and 
epoch to time
```
    time.Unix(secs, 0)
    time.Unix(0, nanos)
```
https://gobyexample.com/epoch

# formatting and parsing time
https://gobyexample.com/time-formatting-parsing
```
package main
import (
	"fmt"
	"time"
)
func main() {
	p := fmt.Println
	t := time.Now()
	p(time.Now())
	p(time.RFC3339)
	p(t.Format("02/01/2006"))  // formate value of t according to supplied formate
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")  // formate, time
	p(t2)
}
```

# Random
math/rand 
`mt.Print(rand.Intn(100))`
generate int float random number
apply seed to generate eandom numbers
https://gobyexample.com/random-numbers

- library strcnv convert string to numbers (int, float etc) `strconv.ParseFloat("1.234", 64)`, with base 10 `strconv.Atoi("135")` 
- net, net/url, `u, err := url.Parse("url")`  https://gobyexample.com/url-parsing
- crypto/sha1: convert any string or data into sha1 hash  https://gobyexample.com/sha1-hashes
- base64 encoding/decoding



# File handeling
```
import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)
```
3 ways to read a file
- 1. dat, err := ioutil.ReadFile("/tmp/dat") // string(dat) contains file content 
- 2. f, err := os.Open("/tmp/dat") 
     b1 := make([]byte, 5)
     n1, err := f.Read(b1)
     fmt.Printf("%v\n", string(b2[:n2]))
     // o3, err := f.Seek(6, 0), 
     // seek function takes two arguments: 
     // 1st relative location from where to read which may be negative
     // 2nd [0,1,2] 0 means start from file start, 1 means start from current pointer location, 2 means pointer at the end of the file
     // seek(-2,2) last two bytes of file
- 3. r4 := bufio.NewReader(f)
     b4, err := r4.Peek(5)
     fmt.Printf("5 bytes: %s\n", string(b4))



3 way to write to a file (https://gobyexample.com/writing-files)
- 1. d1 := []byte("hello\ngo\n")
     err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
- 2. f, err := os.Create("/tmp/dat2")
     defer f.Close()
     d2 := []byte{115, 111, 109, 101, 10}
     n2, err := f.Write(d2)
     f.Sync() // Issue a Sync to flush writes to stable storage.
- 3. w := bufio.NewWriter(f)
     n4, err := w.WriteString("buffered\n")
     w.Flush() // Use Flush to ensure all buffered operations have been applied to the underlying writer.


# path/filepath
(https://gobyexample.com/file-paths)
For generic path related operation like concatinating path, 
    finding and checking absolue or relative path,
    finding extension from filepath 
    saparating filename and dir path from a path
    generating relative path (by using absolute path and current path)

# directory
creation ->  err := os.Mkdir("subdir", 0755)
deletion ->  os.RemoveAll("subdir")
directory hierarchy creation -> err = os.MkdirAll("subdir/parent/child", 0755)
change directory -> err = os.Chdir("subdir/parent/child")
read directory -> c, err := ioutil.ReadDir(".") // . is relative path to path set by change directory
now traverse over c by using 

```
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }
```

# tmp
To create/delete file/dir to default tmp location with random temp name (https://gobyexample.com/temporary-files-and-directories)

# testing library
just import the "testing", now when the name of function start with Test (like TestIntMinBasic), this sunction called automatically tested
also t.Run(testname, func(){}) can be used to run test
(https://github.com/infracloudio/citadel/tree/master/go#concurrency-controls)

# command line: https://gobyexample.com/command-line-flags
   args: os.Args
   flags: import ("flag")
          wordPtr := flag.String("word", "foo", "a string")
          flag.Parse()
          fmt.Println("word:", *wordPtr)
   subcommands: https://gobyexample.com/command-line-subcommands

# Env var
os.Setenv("FOO", "1")   
fmt.Println("FOO:", os.Getenv("FOO"))

// get all enc var
```
for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
}
```


# http clients and server 
https://gobyexample.com/http-clients
https://gobyexample.com/http-servers
context: helpful in error handeling (handel if something goes wrong): https://gobyexample.com/context
(can be used anywhere not only with http)


# Exit
// defers will not be run when using os.Exit
os.Exit(3) // 3 is the exit code

# process
https://gobyexample.com/spawning-processes


# Go modules
go mod init
go mod tidy 
go list 
go list -m all

https://blog.golang.org/using-go-modules
