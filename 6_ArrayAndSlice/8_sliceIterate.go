package main

import "fmt"

func main() {
	mySlice := []string{"This", "is", "Go"}

	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}

	for index := range mySlice {
		fmt.Println(mySlice[index])

	}

	for _, ele := range mySlice {
		fmt.Println(ele)
	}
}
