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

} // here local level scope of main function ends
