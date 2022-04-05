package main

import (
	"fmt"
	"math"
)

// Rectangle Struct and Methods
type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle Struct and Methods
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

// interface is a set of methods.
type shape interface {
	area() float64
	perimeter() float64
}

func interfaceFunc(i shape) {
	fmt.Println(i)
	fmt.Println(i.area())
	fmt.Println(i.perimeter())
	fmt.Printf("%T", i)

}

// ...interface{} is a variadic parameters, which means we can pass any number of arguments to the function. In this example, we pass interface type.
// If we use []interface{} instead of ...interface{}, we can pass only one argument. []interface{} is a interface slice.
func interfaceFunctions(shapes ...interface{}) {
	for _, shape := range shapes {
		switch shape.(type) {
		case rectangle:
			fmt.Println("Rectangle")
			fmt.Println("Area:", shape.(rectangle).area())
			fmt.Println("Perimeter:", shape.(rectangle).perimeter())
		case circle:
			fmt.Println("Circle")
			fmt.Println("Area:", shape.(circle).area())
			fmt.Println("Perimeter:", shape.(circle).perimeter())
			fmt.Println("Diameter:", shape.(circle).diameter())
		}
	}
}

func main() {

	// Rectangle
	rect := rectangle{
		width:  10,
		height: 5,
	}

	// Circle
	circ := circle{
		radius: 5,
	}

	// Interface
	// Thanks to ...interface{} we can pass any number of arguments to the function.
	interfaceFunctions(rect, circ)

	fmt.Println(rectangle{2, 5}.area())

	// Another decleration for instance of rectangle
	rect2 := rectangle{3, 5}
	fmt.Println(rect2.area())
	fmt.Println(rect2.perimeter())
	interfaceFunc(rect2) // We can store what do we want to in interfaceFunc. -> interface's utilities.

	// Another decleration for instance of circle
	circ2 := circle{4}
	fmt.Println(circ2.area())
	fmt.Println(circ2.perimeter())
	fmt.Println(circ2.diameter())
	interfaceFunc(circ2)

}

// https://morioh.com/p/e6edae5f9ddb
