// Go program to demonstrate the multiple
// variable declarations using var keyword
package main

import "fmt"

func main() {

	// Multiple variables of the same type
	// are declared and initialized
	// in the single line along with type
	var geek1, geek2, geek3 int = 232, 784, 854

	// Multiple variables of different type
	// are declared and initialized
	// in the single line without specifying
	// any type
	var geek4, geek5, geek6 = 100, "GFG", 7896.46

	// Display the values of the variables
	fmt.Printf("The value of geek1 is : %d\n", geek1)

	fmt.Printf("The value of geek2 is : %d\n", geek2)

	fmt.Printf("The value of geek3 is : %d\n", geek3)

	fmt.Printf("The value of geek4 is : %d\n", geek4)

	fmt.Printf("The value of geek5 is : %s\n", geek5)

	fmt.Printf("The value of geek6 is : %f", geek6)

}
