package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 int16 = int16(num1)
	var num3 int = int(num2)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)

	var num4 int64 = 23
	var num5 = int(num4)

	fmt.Println(num4)
	fmt.Println(num5)

}
