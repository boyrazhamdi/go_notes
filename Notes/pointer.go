package main

import "fmt"

func main() {

	x := 3

	fmt.Println(x)
	fmt.Println(&x) // & --> address operator

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", &x, &x) // *int and int, these are different types.

	name := "hello"
	y := &name
	fmt.Printf("%T, %v\n", y, y)

	z := 5
	fmt.Println(z)        // 5
	fmt.Println(&z)       // address of 5: 0xc0000aa0b0 => type: *int
	fmt.Println(*(&z))    // dereferencing
	fmt.Println(&(*(&z))) // same address of 5: 0xc0000aa0b0

	x1 := 10
	x2 := x1
	fmt.Println(x1, x2) // 10 10
	x1 = 5
	fmt.Println(x1, x2) // 5 10 -> pass by value

	x3 := 10
	x4 := &x3            // We define x4, address of x3
	fmt.Println(x3, x4)  // 10 0xc000012120
	fmt.Println(x3, *x4) // 10 10

	*x4 = 15
	fmt.Println(x3, *x4) // 15 15 -> (not 10 15 !!) pass by reference

	//
	fmt.Println("---")

	a := 5
	fmt.Println(a) // 5
	double(a)      // 10
	fmt.Println(a) // 5

	b := 5
	fmt.Println(b)        // 5
	doubleWithPointer(&b) // 10
	fmt.Println(b)        // 10

}

func double(num int) {
	num *= 2
	fmt.Println(num)
}

func doubleWithPointer(num *int) { // double --> pointer method
	*num *= 2
	fmt.Println(*num)
}

// x = 5     (value)   => 0x000 (Address of x) => &x => type *int
// y = 0x000 (address) =>    5  (Value of y)   => *x => type  int
