// Go program to illustrate how to concatenate strings
// Using Sprintf function
package main

import "fmt"

func main() {

	// Creating and initializing strings
	str1 := "Tutorial"
	str2 := "of"
	str3 := "Go"
	str4 := "Language"

	// Concatenating strings using
	// Sprintf() function
	result := fmt.Sprintf("%s%s%s%s", str1,
		str2, str3, str4)

	fmt.Println(result)
}
