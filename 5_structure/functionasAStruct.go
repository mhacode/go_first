// Go program to illustrate the function
// as a field in Go structure
package main

import "fmt"

// Finalsalary of function type
type Finalsalary func(int, int) int

type JoiningDate func(string) string

// Creating structure
type Author struct {
	name      string
	language  string
	Marticles int
	Pay       int

	// Function as a field
	salary      Finalsalary
	joiningdate JoiningDate
}

// Main method
func main() {

	// Initializing the fields
	// of the structure
	result := Author{
		name:      "Sonia",
		language:  "Java",
		Marticles: 120,
		Pay:       500,
		salary: func(Ma int, pay int) int {
			return Ma * pay
		},
		joiningdate: func(date string) string {
			return date
		},
	}

	// Display values
	fmt.Println("Author's Name: ", result.name)
	fmt.Println("Language: ", result.language)
	fmt.Println("Total number of articles published in May: ", result.Marticles)
	fmt.Println("Per article pay: ", result.Pay)
	fmt.Println("Total salary: ", result.salary(result.Marticles, result.Pay))
}
