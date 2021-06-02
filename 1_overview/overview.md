# Introduction
1. procedural programming language
2. developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google
3. launched in 2009 as an open-source programming language.

## Hello world Program
```go
package main 
  
import "fmt"
 
func main() {
 
     // prints geeksforgeeks
     fmt.Println("Hello, geeksforgeeks")
}
```
output : Hello, geeksforgeeks
# Explanation of the syntax of Go program:

## package main : (compolsory)
package main of the program. intital point to run the program

## import "fmt" :
it tells the compiler to include this files to the underlying package.

## func main() function :
It is beginning of execution of program.

## fmt.Println("hello"): 
Standard library function to print something as output.

# Comment : like as java

# Following is another example:
```go
package main

import "fmt"

func main()  {
	fmt.Println("1 + 1 = ", 1 +1);
}
```
// 1 + 1  = 2

# Why this “Go language”?

1. Go language is an effort to combine the ease of programming of an interpreted, dynamically typed language with the efficiency and safety of a statically typed, compiled language.
2. It also aims to be modern, with support for networked and multicore computing.

# What excluding in Go which is present in other languages?

1. Go attempts to reduce the amount of typing in both senses of the word. Throughout its design, developers tried to reduce clutter and complexity.
2. There are no forward declarations and no header files; everything is declared exactly once.
3. Stuttering is reduced by simple type derivation using the := declare-and-initialize construct.
4. There is no type hierarchy: types just are, they don’t have to announce their relationships.
5. Like threading consumes 1MB whereas Goroutine consumes 2KB of memory, hence at the same time, we can have millions of goroutine triggered. 
6. Golang a strong language that handles concurrency like C++ and Java. 
   
# Advantages and Disadvantages of Go Language

## Advantages:

1. Flexible- It is concise, simple and easy to read.
2. Concurrency- It allows multiple process running simultaneously and effectively.
3. Quick Outcome- Its compilation time is very fast.
4. Library- It provides a rich standard library.
5. Garbage collection- It is a key feature of go. Go excels in giving a lot of control over memory allocation and has dramatically reduced latency in the most recent versions of the garbage collector.
6. It validates for the interface and type embedding.

## Disadvantages:

1. It has no support for generics, even if there are many discussions about it.
2. The packages distributed with this programming language is quite useful but Go is not so object-oriented in the conventional sense.
3. There is absence of some libraries especially a UI tool kit.

## Some popular Applications developed in Go Language
1. Docker: a set of tools for deploying linux containers
2. Openshift: a cloud computing platform as a service by Red Hat.
3. Kubernetes: The future of seamlessly automated deployment processes
4. Dropbox: migrated some of their critical components from Python to Go.
5. Netflix: for two part of their server architecture.
6. InfluxDB: is an open-source time series database developed by InfluxData.
7. Golang: The language itself was written in Go.

# Features of go language 

1. Language Design: 
   1. Simple and easy to understand. 
   2. Object-Oriented support in the language.
   3. It prefers Composition over Inheritance (Composition).
   4. Do More with Less” is the mantra.
    
2. Package Management:
    1. Support is provided directly in the tooling to get external packages and publish your own packages in a set of easy commands.

3. Powerful standard library:
    1. Go has powerful standard library, which is distributed as packages.
   
4. Static Typing :
    1. Go is static typed language
    2. In this compiler not just work on compiling the code successfully but also ensures on type conversions and compatibility. 
       
5. Testing Support: 
   1. Go provides us the unit testing features by itself
   2. a simple mechanism to write your unit test parallel with your code because of this you can understand you code coverage by your own tests.

6. Platform Independent: 
   1. Go language is just like Java language as it support platform independence.
   2. Due to its modular design and modularity i.e., the code is compiled and is converted into binary form which is as small as possible and hence, it requires no dependency.
   3. Its code can be compiled in any platform or any server and application you work on.
   
# Go not supported ?:
   1. inheritance
   2. no while
   3. no volatile keyword
   4. no method overloading
   5. doesn't have constructor and destructor
   6. Doesn’t support private and public keywords,

Public(Upper case) or private(lower case).

#  Chapter 2 : Identifiers in Go Language

An identifier can be - used for identification purposes
1. variable name
2. function name
3. constant
4. statement labels package name
5. types

Identifier length should be  (4 – 15 letters only).(Recommended)

the underscore character(_) is known as a blank identifier.

The identifier which is allowed to access it from another package is known as the exported identifier.
1. Identifier’s name should be in the Unicode upper case letter.
2.  should be declared in the -
    - package block
    - it is a variable name
    - it is a method name.

```go
package main

import "fmt"

func main()  {
	// String type
	var name = "how are you"
	fmt.Print(name)
}
```
### There is total of three identifiers available in the above example:

main: Name of the package<br/>
main: Name of the function<br/>
name: Name of the variable<br/>

## // Valid identifiers:
_geeks23<br/>
geeks<br/>
gek23sd<br/>
Geeks<br/>
geeKs<br/>
geeks_geeks<br/>

## // Invalid identifiers:
212geeks<br/>
if<br/>
default<br/>

## Following is the list of predeclared identifiers:

### For Constants:
```
true, false, iota, nil
```
### For Types:
```
int, 
int8, 
int16, 
int32, 
int64, 

uint,
uint8, 
uint16, 
uint32, 
uint64, 
uintptr,

float32, 
float64, 

complex128, 
complex64,

bool, 
byte, 
rune, 
string, 
error
```
### For Functions:
```
make, 
len, 
cap, 
new, 
append, 
copy, 
close,
delete, 
complex, 
real, 
imag, 
panic, 
recover
```
# Go Keywords
```go
// Go program to illustrate the
// use of keywords
package main
import "fmt"
  
// Here, package, import, func,
// var are keywords
func main() {
  
// Here, a is a valid identifier
var a = "GeeksforGeeks" 
  
fmt.Println(a)
  
// Here, the default is an
// illegal identifier and
// compiler will throw an error
// var default = "GFG"
}
```
### total 25 keywords present
```go

// Go program to illustrate 
// the use of keywords
  
// Here package keyword is used to 
// include main package in the program
package main
  
// import keyword is used to 
// import "fmt" in your package
import "fmt"
  
// func is used to
// create function
func main() {
  
    // Here, var keyword is used 
    // to create variables
    // Pname, Lname, and Cname 
    // are the valid identifiers
    var Pname = "GeeksforGeeks" 
    var Lname = "Go Language" 
    var Cname = "Keywords"
      
    fmt.Printf("Portal name: %s", Pname)
    fmt.Printf("\nLanguage name: %s", Lname)
    fmt.Printf("\nChapter name: %s", Cname)
  
}
```
## Data Types in Go

1. Basic type: Numbers, strings, and booleans come under this category.
2. Aggregate type: Array and structs come under this category.
3. Reference type: Pointers, slices, maps, functions, and channels come under this category.
4. Interface type

```
Result is: 14.440000
The type of c is : float64
```
## Complex Numbers



