// Go program to illustrate the
// concept of the call by reference
package main

import "fmt"

// function which swap values
func swaap(a, b *int) int {
	var o int
	o = *a
	*a = *b
	*b = o

	return o
}

// Main function
func main() {

	var p int = 10
	var q int = 20
	fmt.Printf("p = %d and q = %d", p, q)

	// call by reference
	swaap(&p, &q)
	fmt.Printf("\np = %d and q = %d", p, q)
}
