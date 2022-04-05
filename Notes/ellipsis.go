package main

import "fmt"

var strSlice = []string{"India", "Canada", "Japan"}

// This is a variadic function.
// A variadic function is a function that can accept an arbitrary number of arguments.
func dnm(a ...string) {
	fmt.Println(a)
}

func main() {
	//dnm(strSlice) // will fail.
	dnm(strSlice...) // we can use ... to pass a slice as an argument.

}

// https://yourbasic.org/golang/three-dots-ellipsis/
