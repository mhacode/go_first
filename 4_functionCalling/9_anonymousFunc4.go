package main

import "fmt"

func GFg() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "GeeksforGeeks"
	}
	return myf
}

func main() {
	val := GFg()
	fmt.Println(val("Welcome ", "to "))
}
