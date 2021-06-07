package main

import "fmt"

type author struct {
	name   string
	branch string
}

func (a *author) show_1(abranch string) {
	(*a).branch = abranch
}

func (a author) show_2() {
	a.name = "GGG"
	fmt.Println("Author's name(Before) : ", a.name)
}

func main() {
	res := author{
		name:   "sona",
		branch: "CSE",
	}

	fmt.Println("Branch Name(Before): ", res.branch)

	res.show_1("ECE")

	fmt.Println("Branch Name(After) :", res.branch)

	(&res).show_2()
	fmt.Println("Author's Name(After) :", res.name)

}
