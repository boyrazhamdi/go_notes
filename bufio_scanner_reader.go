package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// bufio.NewScanner
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type your name:")
	scanner.Scan()
	input := scanner.Text() // Scan yaptıktan sonra bool döndürüyor, texti nasıl yakalıyoruz ki?
	fmt.Printf("You typed: %v", input)
	fmt.Println("")
	fmt.Println("Type your age:")
	scanner.Scan()
	//input1, _ := .ParseInt(scanner.Text(), 10, 64) // text, base, byte // we can use atoi instead.
	input1, _ := strconv.Atoi(scanner.Text())
	fmt.Println(input1)

	// scanner1 := bufio.NewScanner(os.Stdin).Scan() // true döndürüyor
	// scanner1.Text() // hata veriyor, aynı işlemler olmuş olmuyor mu?
	// fmt.Println(scanner1)

	// bufio.NewReader
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n') //  \n means, when you entered, func catch the text.
	// fmt.Println(input)

	// deneme
	fmt.Println(input)

}
