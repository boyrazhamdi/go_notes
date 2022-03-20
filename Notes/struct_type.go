package main

import "fmt"

func main() {

	// Struct
	// Struct is an underlying type.
	// Pass-by-value
	// Generally we define this type of variables outside of main func.
	var employee1 struct { // Here, employee1 is a defined type or named type.
		id        int
		name      string
		age       int
		isMarried bool
	}

	fmt.Println(employee1)         // Default values will be seen.(0, "", false..)
	fmt.Printf("%#v\n", employee1) // %#v a Go-syntax representation of the value, \n: new line
	fmt.Println(employee1.id)

	employee1.id = 1
	employee1.name = "Jack"
	employee1.age = 30
	employee1.isMarried = true

	fmt.Printf("%#v\n", employee1)
	fmt.Println(employee1.id)
	fmt.Println(employee1.name)
	fmt.Println(employee1.age)
	fmt.Println(employee1.isMarried)

	var employee2 struct {
		id        int
		name      string
		age       int
		isMarried bool
	}

	employee2.id = 2
	employee2.name = "Harper"
	employee2.age = 35
	employee2.isMarried = true

	fmt.Printf("%#v\n", employee2)
	fmt.Println(employee2.id)
	fmt.Println(employee2.name)
	fmt.Println(employee2.age)
	fmt.Println(employee2.isMarried)

	// ... exhausting
	// If you see this type of code, run away.
	// We prefer type instead of struct.
	// What type means?
	// It means, we generate our type in Go => It sounds good?
	// Type in go is similar to class in other languages.

	// Type
	// Pass-by-value
	type employee struct {
		id        int
		name      string
		age       int
		isMarried bool
		hobbies   []string
	}

	// Define type variable
	var e1 employee
	e1.id = 1
	e1.name = "Jack"
	e1.age = 30
	e1.isMarried = true
	e1.hobbies = []string{"Tennis", "Baseball"}

	var e2 employee
	e2.id = 2
	e2.name = "Harper"
	e2.age = 35
	e2.isMarried = true
	e2.hobbies = []string{"Paint", "Cook"}

	fmt.Printf("%#v\n", e1) // No longer these are not struct, these are employee data type that we generated.
	fmt.Printf("%#v\n", e2) // Same as above. Additionally, type is underlying data type based on struct like slice(based on array)

	// Shorthand define type variable
	e3 := employee{
		id:        3,
		name:      "Oblivion",
		age:       40,
		isMarried: false,
		hobbies: []string{
			"Read", "Watch Movie",
		},
	}

	fmt.Println(e3)

	// Define an empty instance for type
	e4 := employee{}
	fmt.Println(e4)
	e4.id = 4
	fmt.Println(e4)

	// Example for pass-by-ref
	e5 := employee{
		id:   5,
		name: "John",
		age:  13,
	}
	fmt.Println(e5)
	e6 := e5
	fmt.Println(e6)
	e6.name = "Bill"
	fmt.Println(e5) // Stay same
	fmt.Println(e6)

	// Inheritance

	// Note:
	// In classic OOP, automobile is automobile, truck is a automobile = is a relation
	// In GO, truck is not a automobile = has a relation

	type automobile struct {
		color  string
		wheels int
		speed  int
	}

	type truck struct {
		automobile
		hasLicensePlate bool
	}

	a1 := automobile{
		color:  "red",
		wheels: 4,
		speed:  180,
	}

	fmt.Println(a1)

	t1 := truck{
		automobile: automobile{
			color:  "gray",
			wheels: 10,
			speed:  120,
		},
		hasLicensePlate: true,
	}

	fmt.Println(t1)

	// Anonymous Struct
	// Struct with initialize
	type car struct {
		make    string
		model   string
		mileage int
	}

	newCar := car{
		make:    "Audi",
		model:   "A6",
		mileage: 100000,
	}

	fmt.Println(newCar)

	anonymousNewCar := struct {
		make    string
		model   string
		mileage int
	}{
		make:    "Audi",
		model:   "A8",
		mileage: 100000,
	}

	fmt.Println(anonymousNewCar)

	// Data type in type
	type mile float64
	type mystr string

	var m1 mile
	m1 = 4.7

	fmt.Printf("%T, %v\n", m1, m1) // type is mile (not float64)
	var f1 float64
	f1 = 5.3
	fmt.Printf("%T, %v\n", f1, f1)
	//fmt.Println(m1 + f1) // type error: mile type + float64 type
	//We can do type conversions
	fmt.Println(m1 + mile(f1))
	fmt.Println(float64(m1) + f1)

	// Shorthand declaration
	m2 := mile(4.6)
	fmt.Println(m2)

	/*Is Go an object-oriented language?
	Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. The concept of “interface” in Go provides a different approach that we believe is easy to use and in some ways more general. There are also ways to embed types in other types to provide something analogous—but not identical—to subclassing. Moreover, methods in Go are more general than in C++ or Java: they can be defined for any sort of data, even built-in types such as plain, “unboxed” integers. They are not restricted to structs (classes).*/

}
