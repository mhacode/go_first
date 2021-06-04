package main

import "fmt"

func init() {
	fmt.Println("Welcome to init function 1")
}

func init() {
	fmt.Println("Welcome to second init 2")
}

func main() {
	fmt.Println("Welcome to main() func")
}
