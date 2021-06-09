package main

type tank interface {
	Tarea() float64
	Volume() float64
}

func (m myvalue) Tarea() float64 {
	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

func (m myvalue) volume() float64 {
	return 3.1416 * m.radius * m.height
}

type myvalue struct {
	radius float64
	height float64
}

func main() {

}
