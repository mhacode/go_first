package main

import "fmt"

func main() {
	fmt.Println("While loop In Go")
	// while loop
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
	}
	//
	fmt.Println("For Loop in Go")
	for i := 0; i <10; i++ {
		fmt.Println(i)
	}
}
