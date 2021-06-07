package main

import "fmt"

func main() {

	mul, div := mul_div(105, 7)

	fmt.Println(mul)
	// fmt.Println(div)
}

func mul_div(n1 int, n2 int) (int, int) {
	return n1 * n2, n1 / n2
}
