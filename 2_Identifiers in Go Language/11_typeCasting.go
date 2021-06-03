package main

import "fmt"

func main() {
	var total int = 846
	var number int = 19
	var avg float32

	avg = float32(total) / float32(number)
	fmt.Println(avg)
	fmt.Printf("%v\n", avg)
	fmt.Printf("%f\n", avg)
}
