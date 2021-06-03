package main

import "fmt"

func main() {
	a := 4
	fmt.Println(a)  // print value
	fmt.Println(&a) // print address
	b := &a         // assign memory address to b
	fmt.Println(*b) //

}
