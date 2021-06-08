package main

import "fmt"

type Student struct {
	name    string
	age     int
	address Address
}

type Address struct {
	postAddress  int
	streetNumber string
}

func main() {

	student1 := Student{
		name: "A",
		age:  23,
		address: Address{
			postAddress:  234,
			streetNumber: "SSS",
		}}

	fmt.Println(student1.name)
	fmt.Println(student1.address.postAddress)
	fmt.Println(student1.address.streetNumber)

	address1 := Address{231, "Sababaror"}
	fmt.Println(address1)

}
