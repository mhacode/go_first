// Go program to show the use of := operator
// to declare local variables
package main

import "fmt"

// using var keyword to declare
// and initialize the variable
// it is package or you can say
// global level scope
var geek1 = 900

func main() {

	// using short variable declaration
	// inside the main function
	// it has local scope i.e. can't
	// accessed outside the main function
	geek2 := 200

	// accessing geek1 inside the function
	fmt.Println(geek1)

	// accessing geek2 inside the function
	fmt.Println(geek2)

}
