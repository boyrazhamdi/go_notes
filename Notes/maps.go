package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	//https://medium.com/@kdnotes/how-to-sort-golang-maps-by-value-and-key-eedc1199d944

	// dict := map[keyType]valueType{}
	// maps -> pass by referance

	passStudents := map[string]bool{
		"Ali":  true,
		"Ayşe": false,
	}

	fmt.Println(passStudents)

	customers := map[string]int{
		"Ahmet":   40,
		"Ayşe":    17,
		"Selim":   14,
		"Mustafa": 70,
	}

	fmt.Println(customers)
	fmt.Println(reflect.ValueOf(customers).MapKeys()) // Gives all keys

	fmt.Println(customers["Ahmet"])

	// Checking Key if it's exist in map.
	value, isExist := customers["Aslı"]
	fmt.Println(value, isExist)

	value, isExist = customers["Ahmet"]
	fmt.Println(value, isExist)

	// Add element
	customers["Aslı"] = 25 // Adds not in order!
	fmt.Println(customers)

	// Delete element
	delete(customers, "Aslı")
	fmt.Println(customers)

	// Get number of elements in map
	fmt.Println(len(customers))

	// Get type
	fmt.Println(reflect.TypeOf(customers))        // We get types of key and value
	fmt.Println(reflect.TypeOf(customers).Key())  // Type of Key
	fmt.Println(reflect.TypeOf(customers).Elem()) // Type of Value
	fmt.Println(reflect.TypeOf(customers).Kind())
	fmt.Println(reflect.ValueOf(customers).Kind())

	// Create emtpy map
	emptyMap := make(map[string]int)
	fmt.Println(emptyMap)

	// Pass by reference
	copyCustomers := customers
	fmt.Println(copyCustomers)
	delete(copyCustomers, "Ahmet")
	fmt.Println(copyCustomers)
	fmt.Println(customers)

	// Loop in map

	for k, v := range customers {
		fmt.Println(k, v)
	}

	// Just keys
	for k := range customers {
		fmt.Println(k)
	}

	// Just values
	for _, v := range customers {
		fmt.Println(v)
	}

	// Sort
	//sortedCustomers := make([]string, 0, len(customers)) // make empty slice
	var sortedCustomers []string
	for k := range customers {
		sortedCustomers = append(sortedCustomers, k)
	}

	sort.Strings(sortedCustomers) // sort slice

	for _, k := range sortedCustomers {
		fmt.Println(k, customers[k])
	}

}
