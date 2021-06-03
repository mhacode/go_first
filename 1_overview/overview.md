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
1. The complex numbers are divided into two parts are shown in the below table. float32 and float64 are also part of these complex numbers. 
The in-built function creates a complex number from its imaginary and real part and in-built imaginary and real function extract those parts.

```go
// Go program to illustrate
// the use of complex numbers
package main
import "fmt"
 
func main() {
     
   var a complex128 = complex(6, 2)
   var b complex64 = complex(9, 2)
   fmt.Println(a)
   fmt.Println(b)
    
   // Display the type
  fmt.Printf("The type of a is %T and "+
            "the type of b is %T", a, b)
}
```
```
Output: 

(6+2i)
(9+2i)
The type of a is complex128 and the type of b is complex64
```

## Booleans

The boolean data type represents only one bit of information either true or false. The values of type boolean are not converted implicitly or explicitly to any other type.

Example:
```go

Go
// Go program to illustrate
// the use of booleans
package main
import "fmt"

func main() {

    // variables
str1 := "GeeksforGeeks"
str2:= "geeksForgeeks"
str3:= "GeeksforGeeks"
result1:= str1 == str2
result2:= str1 == str3

// Display the result
fmt.Println( result1)
fmt.Println( result2)

// Display the type of
// result1 and result2
fmt.Printf("The type of result1 is %T and "+
"the type of result2 is %T",
result1, result2)

}
```
```
Output:

false
true
The type of result1 is bool and the type of result2 is bool
```
# Strings

The string data type represents a sequence of Unicode code points. 
Or in other words, we can say a string is a sequence of immutable bytes, 
means once a string is created you cannot change that string. 
A string may contain arbitrary data, including bytes with zero value in the human-readable form.

Example:
```go

// Go program to illustrate
// the use of strings
package main
import "fmt"

func main() {

    // str variable which stores strings
str := "GeeksforGeeks"

// Display the length of the string
fmt.Printf("Length of the string is:%d",
len(str))

// Display the string
fmt.Printf("\nString is: %s", str)

// Display the type of str variable
fmt.Printf("\nType of str is: %T", str)
}
```
```
Output:

Length of the string is:13
String is: GeeksforGeeks
Type of str is: string
```

# Go Variables

### Declaring a Variable

var variable_name type = expression

1. no such concept of an uninitialized variable in Go language


```go
package main

// Go program to illustrate
// concept of variable
import "fmt"

func main() {

	// Variable declared and
	// initialized without the
	// explicit type
	var myvariable1 = 20
	var myvariable2 = "GeeksforGeeks"
	var myvariable3 = 34.80

	// Display the value and the
	// type of the variables
	fmt.Printf("The value of myvariable1 is : %d\n",
		myvariable1)

	fmt.Printf("The type of myvariable1 is : %T\n",
		myvariable1)

	fmt.Printf("The value of myvariable2 is : %s\n",
		myvariable2)

	fmt.Printf("The type of myvariable2 is : %T\n",
		myvariable2)

	fmt.Printf("The value of myvariable3 is : %f\n",
		myvariable3)

	fmt.Printf("The type of myvariable3 is : %T\n",
		myvariable3)

}

```
```go
// Go program to illustrate
// concept of variable
package main

import "fmt"

func main() {

    // Variable declared and 
    // initialized without expression
    var myvariable1 int
    var myvariable2 string
    var myvariable3 float64
  
    // Display the zero-value of the variables
    fmt.Printf("The value of myvariable1 is : %d\n",
                                     myvariable1)
  
    fmt.Printf("The value of myvariable2 is : %s\n",
                                     myvariable2)
  
    fmt.Printf("The value of myvariable3 is : %f",
                                     myvariable3)
}
```

```
Output:

The value of myvariable1 is : 0
The value of myvariable2 is :
The value of myvariable3 is : 0.000000
```
2. If you use type, then you are allowed to declare multiple variables of the same type in the single declaration

```go
// Go program to illustrate
// concept of variable
package main
import "fmt"
   
func main() {
   
    // Multiple variables of the same type
    // are declared and initialized
    // in the single line
    var myvariable1, myvariable2, myvariable3 int = 2, 454, 67
   
   // Display the values of the variables
   fmt.Printf("The value of myvariable1 is : %d\n",
                                       myvariable1)
  
   fmt.Printf("The value of myvariable2 is : %d\n",
                                       myvariable2)
  
   fmt.Printf("The value of myvariable3 is : %d",
                                      myvariable3)
}
```
```
Output:

The value of myvariable1 is : 2
The value of myvariable2 is : 454
The value of myvariable3 is : 67
```

3. If you remove type, then you are allowed to declare multiple variables of a different type in the single declaration. The type of variables is determined by the initialized values.

Example:

```.go
// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Multiple variables of different types
// are declared and initialized in the single line
var myvariable1, myvariable2, myvariable3 = 2, "GFG", 67.56

// Display the value and
// type of the variables
fmt.Printf("The value of myvariable1 is : %d\n",
myvariable1)

fmt.Printf("The type of myvariable1 is : %T\n",
myvariable1)

fmt.Printf("\nThe value of myvariable2 is : %s\n",
myvariable2)

fmt.Printf("The type of myvariable2 is : %T\n",
myvariable2)

fmt.Printf("\nThe value of myvariable3 is : %f\n",
myvariable3)

fmt.Printf("The type of myvariable3 is : %T\n",
myvariable3)
}
```

```
Output:

The value of myvariable1 is : 2
The type of myvariable1 is : int

The value of myvariable2 is : GFG
The type of myvariable2 is : string

The value of myvariable3 is : 67.560000
The type of myvariable3 is : float64
```

# Using short variable declaration:

variable_name:= expression

```go
// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Using short variable declaration
myvar1 := 39
myvar2 := "GeeksforGeeks"
myvar3 := 34.67

// Display the value and type of the variables
fmt.Printf("The value of myvar1 is : %d\n", myvar1)
fmt.Printf("The type of myvar1 is : %T\n", myvar1)

fmt.Printf("\nThe value of myvar2 is : %s\n", myvar2)
fmt.Printf("The type of myvar2 is : %T\n", myvar2)

fmt.Printf("\nThe value of myvar3 is : %f\n", myvar3)
fmt.Printf("The type of myvar3 is : %T\n", myvar3)
}
```

```
Output:

The value of myvar1 is : 39
The type of myvar1 is : int

The value of myvar2 is : GeeksforGeeks
The type of myvar2 is : string

The value of myvar3 is : 34.670000
The type of myvar3 is : float64
```
### Key Points

1. Most of the local variables are declared and initialized by using short variable declarations due to their brevity and flexibility.
2. The var declaration of variables are used for those local variables which need an explicit type that differs from the initializer expression, 
   or for those variables whose values are assigned later, and the initialized value is unimportant. 
   
```go
// Go program to illustrate
   // concept of variable
   package main
   import "fmt"

func main() {

// Using short variable declaration
// Multiple variables of same types
// are declared and initialized in
// the single line
myvar1, myvar2, myvar3 := 800, 34, 56

// Display the value and
// type of the variables
fmt.Printf("The value of myvar1 is : %d\n", myvar1)
fmt.Printf("The type of myvar1 is : %T\n", myvar1)

fmt.Printf("\nThe value of myvar2 is : %d\n", myvar2)
fmt.Printf("The type of myvar2 is : %T\n", myvar2)

fmt.Printf("\nThe value of myvar3 is : %d\n", myvar3)
fmt.Printf("The type of myvar3 is : %T\n", myvar3)
}

```

```
Output:

The value of myvar1 is : 800
The type of myvar1 is : int

The value of myvar2 is : 34
The type of myvar2 is : int

The value of myvar3 is : 56
The type of myvar3 is : int
```

In a short variable declaration, you are allowed to initialize a set of variables by the calling function that returns multiple values.
Example:
A short variable declaration acts like an assignment only when for those variables that are already declared in the same lexical block. The variables that are declared in the outer block are ignored. And at least one variable is a new variable out of these two variables as shown in the below example.

Example:
```go


// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Using short variable declaration
// Here, short variable declaration acts
// as an assignment for myvar2 variable
// because same variable present in the same block
// so the value of myvar2 is changed from 45 to 100
myvar1, myvar2 := 39, 45
myvar3, myvar2 := 45, 100

// If you try to run the commented lines,
// then compiler will gives error because
// these variables are already defined
// myvar1, myvar2 := 43, 47
// myvar2:= 200

// Display the values of the variables
fmt.Printf("The value of myvar1 and myvar2 is : %d %d\n",
myvar1, myvar2)

fmt.Printf("The value of myvar3 and myvar2 is : %d %d\n",
myvar3, myvar2)
}
```
```
Output:

The value of myvar1 and myvar2 is : 39 100
The value of myvar3 and myvar2 is : 45 100
```
3. Using short variable declaration you are allowed to declare multiple variables of different types in the single declaration. The type of these variables are determined by the expression.
Example:

```go
// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Using short variable declaration
// Multiple variables of different types
// are declared and initialized in the single line
myvar1, myvar2, myvar3 := 800, "Geeks", 47.56

// Display the value and type of the variables
fmt.Printf("The value of myvar1 is : %d\n", myvar1)
fmt.Printf("The type of myvar1 is : %T\n", myvar1)

fmt.Printf("\nThe value of myvar2 is : %s\n", myvar2)
fmt.Printf("The type of myvar2 is : %T\n", myvar2)

fmt.Printf("\nThe value of myvar3 is : %f\n", myvar3)
fmt.Printf("The type of myvar3 is : %T\n", myvar3)

}
```
```
Output:

The value of myvar1 is : 800
The type of myvar1 is : int

The value of myvar2 is : Geeks
The type of myvar2 is : string

The value of myvar3 is : 47.560000
The type of myvar3 is : float64
```

# Constants- Go Language : Fixed value

```go
package main
 
import "fmt"
 
const PI = 3.14
 
func main()
{
    const GFG = "GeeksforGeeks"
    fmt.Println("Hello", GFG)
 
    fmt.Println("Happy", PI, "Day")
 
    const Correct= true
    fmt.Println("Go rules?", Correct)
}
```
```
Hello GeeksforGeeks
Happy 3.14 Day
Go rules? true
```

## Following are some examples of Integer Constant:

```
85         /* decimal */
0213       /* octal */
0x4b       /* hexadecimal */
30         /* int */
30u        /* unsigned int */
30l        /* long */
30ul       /* unsigned long */
212         /* Legal */
215u        /* Legal */
0xFeeL      /* Legal */
078         /* Illegal: 8 is not an octal digit */
032UU       /* Illegal: cannot repeat a suffix */
```
```go

package main

import "fmt"

func main()
{
const A = "GFG"
var B = "GeeksforGeeks"

    // Concat strings.
    var helloWorld = A+ " " + B
    helloWorld += "!"
    fmt.Println(helloWorld)
     
    // Compare strings.
    fmt.Println(A == "GFG")  
    fmt.Println(B < A)
}
```

```
Output:

GFG GeeksforGeeks!
true
false
```

## Go Operators
1. Arithmetic Operators
2. Relational Operators
3. Logical Operators
4. Bitwise Operators
5. Assignment Operators
6. Misc Operators
```go
// Go program to illustrate the
// use of arithmatic operators
package main
   
import "fmt"
   
func main() {
   p:= 34
   q:= 20
      
   // Addition
   result1:= p + q
   fmt.Printf("Result of p + q = %d", result1)
      
   // Subtraction
   result2:= p - q
   fmt.Printf("\nResult of p - q = %d", result2)
      
   // Multiplication
   result3:= p * q
   fmt.Printf("\nResult of p * q = %d", result3)
      
   // Division
   result4:= p / q
   fmt.Printf("\nResult of p / q = %d", result4)
      
   // Modulus
   result5:= p % q
   fmt.Printf("\nResult of p %% q = %d", result5)
}
```
```
Result of p + q = 54
Result of p - q = 14
Result of p * q = 680
Result of p / q = 1
Result of p % q = 14
```
## Logical Operators
They are used to combine two or more conditions/constraints or to complement the evaluation of the original condition in consideration.


Logical AND: The ‘&&’ operator returns true when both the conditions in consideration are satisfied. Otherwise it returns false. For example, a && b returns true when both a and b are true (i.e. non-zero).

Logical OR: The ‘||’ operator returns true when one (or both) of the conditions in consideration is satisfied. Otherwise it returns false. For example, a || b returns true if one of a or b is true (i.e. non-zero). Of course, it returns true when both a and b are true.

Logical NOT:The ‘!’ operator returns true the condition in consideration is not satisfied. Otherwise it returns false. For example, !a returns true if a is false, i.e. when a=0.
Example:
```go

// Go program to illustrate the
// use of logical operators
package main
import "fmt"
func main() {
var p int = 23
var q int = 60

    if(p!=q && p<=q){ 
        fmt.Println("True")
    }
        
    if(p!=q || p<=q){ 
        fmt.Println("True")
    }
        
    if(!(p==q)){ 
        fmt.Println("True")
    }

}
```

```
Output:

True
True
True
```
## Misc Operators

&: This operator returns the address of the variable.
*: This operator provides pointer to a variable.
<-:The name of this operator is received. It is used to receive a value from the channel.

```go
package main

import "fmt"

func  main()  {
	a := 4
	fmt.Println(a) // print value
	fmt.Println(&a) // print address
	b := &a // assign memory address to b
	fmt.Println(*b) //

}
```

## Rune in Golang
```go
// Simple Go program to illustrate
// how to create a rune
package main
  
import (
    "fmt"
    "reflect"
)
  
func main() {
  
    // Creating a rune
    rune1 := 'B'
    rune2 := 'g'
    rune3 := '\a'
  
    // Displaying rune and its type
    fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s", rune1,
                             rune1, reflect.TypeOf(rune1))
      
    fmt.Printf("\nRune 2: %c; Unicode: %U; Type: %s", rune2,
                               rune2, reflect.TypeOf(rune2))
      
    fmt.Printf("\nRune 3: Unicode: %U; Type: %s", rune3, 
                                 reflect.TypeOf(rune3))
  
}
```

```
Rune 1: B; Unicode: U+0042; Type: int32
Rune 2: g; Unicode: U+0067; Type: int32
Rune 3: Unicode: U+0007; Type: int32
```
## Scope of Variables in Go
1. Local Variables(Declared Inside a block or a function)
2. Global Variables(Declared outside a block or a function)

### Local Variable:

1. Variables that are declared inside a function, or a block are termed as Local variables. These are not accessible outside the function or block.
2. These variables can also be declared inside the for, while statement etc. inside a function.
3. However, these variables can be accessed by the nested code blocks inside a function.
4. These variables are also termed as the block variables.
5. There will be a compile-time error if these variables are declared twice with the same name in the same scope.
6. These variables don’t exist after the function’s execution is over.
7. The variable which is declared outside the loop is also accessible within the nested loops. It means a global variable will be accessible to the methods and all loops. The local variable will be accessible to loop and function inside that function.
8. A variable which is declared inside a loop body will not be visible to the outside of loop body.

```go
// Go program to illustrate the
// local variables
package main

import "fmt"

// main function
func main() { // from here local level scope of main function starts 

	// local variables inside the main function
	var myvariable1, myvariable2 int = 89, 45

	// Display the values of the variables
	fmt.Printf("The value of myvariable1 is : %d\n",
		myvariable1)

	fmt.Printf("The value of myvariable2 is : %d\n",
		myvariable2)

} // here local
```
### Global Variables

1. The variables which are defined outside a function, or a block is termed as Global variables.
2. These are available throughout the lifetime of a program.
3. These are declared at the top of the program outside all the functions or blocks.
4. These can be accessed from any portion of the program.
