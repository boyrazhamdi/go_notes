package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	// int to int
	var index int8 = 15
	var bigIndex int32
	bigIndex = int32(index)
	fmt.Println(bigIndex)
	fmt.Printf("index data type:    %T\n", index)
	fmt.Printf("bigIndex data type: %T\n", bigIndex)

	// Here, reverse is possible but probably data loss may happen.

	//int to float and float to int
	x := 5
	y := 5.6

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	fmt.Println(float64(x) + y)
	fmt.Println(x + int(y))

	// int to string
	a := strconv.Itoa(12)
	fmt.Printf("%q %T\n", a, a)

	// string to int
	lines_yesterday := "50"
	lines_today := "108"

	yesterday, err := strconv.Atoi(lines_yesterday)
	if err != nil {
		log.Fatal(err)
	}

	today, err := strconv.Atoi(lines_today)
	if err != nil {
		log.Fatal(err)
	}
	lines_more := today - yesterday

	fmt.Println(lines_more)

	b := "not a number"
	c, err := strconv.Atoi(b)
	fmt.Println(c)
	fmt.Println(err)

	// string to float
	f := 5524.53
	fmt.Println("Sammy has " + fmt.Sprint(f) + " points.")

	// str to byte and byte to str
	h := "my string"

	i := []byte(h)

	j := string(i)

	fmt.Println(h)

	fmt.Println(i)

	fmt.Println(j)

}
