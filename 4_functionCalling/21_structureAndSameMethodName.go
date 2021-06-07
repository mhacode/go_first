package main

import "fmt"

type value_1 string
type value_2 int

func (a value_1) display() value_1 {
	return a + "forGeeks"
}

func (p value_2) display() value_2 {
	return p + 298
}

func main() {
	res1 := value_1("Geeks")
	res2 := value_2(234)

	fmt.Println("Result 1 :", res1.display())
	fmt.Println("Result 2 :", res2.display())
}
