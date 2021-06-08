package main

import "fmt"

// Structure
type details struct {

	// Fields of the
	// details structure
	name    string
	age     int
	gender  string
	psalary int
}

// Nested structure
type employee struct {
	post string
	eid  int
	details
}

func main() {

	values := employee{
		post: "HR",
		eid:  43,
		details: details{
			name:    "sss",
			age:     22,
			gender:  "Maa",
			psalary: 34,
		},
	}

	fmt.Println(values.name)
	fmt.Println(values.age)
	fmt.Println(values.gender)
	fmt.Println(values.psalary)
	fmt.Println(values.post)

	fmt.Println(values.promoted(30))
}

func (d details) promoted(tsalary int) int {
	return d.psalary * tsalary
}
