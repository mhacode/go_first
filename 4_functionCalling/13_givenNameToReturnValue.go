package main

import "fmt"

func myfunction(p, q int) (rectangle int, square int) {
	rectangle = p * q
	square = p * p
	return
}

func main() {
	var area1, area2 = myfunction(2, 4)
	fmt.Println(area1)
	fmt.Println(area2)
}
