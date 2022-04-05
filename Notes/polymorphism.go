package main

import (
	"fmt"
)

// What is polymorphism?
// Polymorphism is the ability of an object to take on many forms.
// In this example, we use area() method to calculate the area of different shapes.
// Methods name is same, but the parameters and return values are different.
type shape interface {
	area() float64
}

func printArea(shapes ...shape) {
	for _, shape := range shapes {
		fmt.Println(fmt.Sprintf("%T", shape), "Alan : ", shape.area())
		//fmt.Println(reflect.TypeOf(shape), "Alan : ", shape.area()) // An alternate
	}
}

// shapes ...shape ya da shapes ...interface{} sor.
// interface kısmında değiştirince sorun olmadı, ancak burada shape.area'da hata veriyor.
// interface ile interface{} i de sor.

type triangle struct {
	a float64
	h float64
}

func (t triangle) area() float64 {
	return (t.a * t.h) / 2
}

type square struct {
	a float64
}

func (s square) area() float64 {
	return (s.a * s.a)
}

type rectangle struct {
	a, b float64
}

func (r rectangle) area() float64 {
	return (r.a * r.b)
}

func main() {

	t := triangle{3, 4}
	s := square{4}
	r := rectangle{4, 5}

	printArea(t, s, r)

}
