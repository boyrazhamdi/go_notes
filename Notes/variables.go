package main

import (
	"fmt"
	"reflect"
)

func main() {
	var var1 = "hello"
	fmt.Println(var1)

	//var1 := "hello world" // We can't do. Because we previously defined var1.
	var1 = "world"
	fmt.Println(var1)

	var1, var2 := "hello", "world" // But we can do this.
	fmt.Println(var1, var2)

	// get type
	fmt.Println(reflect.TypeOf(var1))
}
