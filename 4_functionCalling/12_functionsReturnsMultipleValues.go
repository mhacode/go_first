package main

import "fmt"

func myfunc(p, q int) (int, int, int) {
	return p - q, p * q, p + q
}

func main() {
	var myvar1, myvar2, myvar3 = myfunc(4, 2)
	fmt.Println(myvar1)
	fmt.Println(myvar2)
	fmt.Println(myvar3)

}
