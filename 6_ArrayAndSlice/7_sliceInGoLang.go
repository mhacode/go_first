package main

import "fmt"

func main() {

	// arr := [4]string{"this", "is", "go", "lang"}
	var arr = [4]string{"this", "is", "go", "lang"}

	fmt.Println(arr)

	mySlice := arr[1:3]

	fmt.Println(mySlice)

}
