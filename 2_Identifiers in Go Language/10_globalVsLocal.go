// Go program to show compiler giving preference
// to a local variable over a global variable
package main

import "fmt"

// global variable declaration
var globalVariable int = 100

func main() { // from here local level scope starts

	// local variables inside the main function
	// it is same as global variable
	var localVariable int = 200

	// Display the value
	fmt.Printf("The value of myvariable1 is : %d\n",
		localVariable)

} // here local level scope ends
