package main

import "fmt"

func main() {
	num := 10
	myString := "Hello, World!"

	var index = myString[0]
	fmt.Printf("%v, %T\n", index, index)

	for i, v := range myString {
		fmt.Println(i, string(v))
	}

	println("Hello, World! ", num)
}
