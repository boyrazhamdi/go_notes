package main

import "fmt"

func main() {
	x := 5
	y := 5.6

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	fmt.Println(float64(x) + y)
	fmt.Println(x + int(y))
}
