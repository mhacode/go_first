package main

import "fmt"

func main() {

	slc := []string{"c", "python", "c++", "perl"}

	fmt.Println(slc)

	myFunc(slc)

	fmt.Println(slc)
}

func myFunc(e []string) {
	e[2] = "java"
	//fmt.Println(e)
}
