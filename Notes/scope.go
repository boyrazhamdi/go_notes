package main

import "fmt"

// These variables use memory until the project finishes.
var variable1 = "Var1: Package Scope"
var variable5 = "Var5: No one can change me"

//variable2 := "Var2" // we can't define a shorthand variable in package like this. Because, go wants keyword(package/import/var/func).

func main() {
	// But these variable use memory until func finished. That's why we use scope.
	variable3 := "Var3: Func Scope"

	fmt.Println(variable5)

	variable5 = "Var5: Get over here"

	fmt.Println(variable1)
	fmt.Println(variable3)
	fmt.Println(variable5)

	printsVar()

	//fmt.Println(variable4) // Same Ex1
}

func printsVar() {

	fmt.Println(variable1)
	//fmt.Println(variable3) // Ex1: This function can not access the other func's variable(in this context: variable3)

	variable4 := "Var4: Func Scope"

	fmt.Println(variable4)
}
