// Go Program to illustrate the concept of call by value
package main

import "fmt"

// function which swap values
func swap(a, b int) int {
	var temp int
	temp = a
	a = b
	b = temp
	return temp
}

// Main function
func main() {
	var p int = 10
	var q int = 20
	fmt.Printf("p = %d and q = %d", p, q)

	swap(p, q)
	fmt.Printf("\np = %d and q = %d", p, q)
}
