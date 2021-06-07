package main

import "fmt"



type author struct {
	name      string
	branch    string
	particles int
	salary    int
}


func (a author) show(){
	fmt.Println("Author's Name = "a.name)
	fmt.Println("Author's Branch = "a.branch)
	fmt.Println("Author's particles = "a.particles)
	fmt.Println("Author's salary = "a.salary)
	
}

func main() {

	res := author{
		name : "Mainul",
		branch : "Khulna",
		particles : "2003"
		salary : "0000",
	}

	res.show()

}
