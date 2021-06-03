package main

import "fmt"

func main() {
	fmt.Println("Hello")

	rvariable := []string{"GFG", "Geeks", "GeeksForGeeks"}
	for i := 0; i < len(rvariable); i++ {
		fmt.Println(rvariable[i])
	}

	for ii, j := range rvariable {
		fmt.Println(ii, j)
	}

	mmap := map[int]string{
		22 : "geeks",
		33 : "for",
		44 : "GeeksForGeeks",
	}

	for key, val :=  range mmap{
		fmt.Println(key, val)
	}

	for k := range rvariable{
		fmt.Println(k)
	}
}
