package main

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (r rectangle) perim() float64 {
	return 2 * (r.length + r.width)
}

func measurement(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	rectangle_1 := rectangle{
		length: 10,
		width:  20,
	}
	measurement(rectangle_1)
}