package main

import "fmt"

func main() {
	fmt.Println("Pointers!!")

	var simple_variable int = 10
	var simple_variable_copy int = simple_variable

	fmt.Println(simple_variable, simple_variable_copy)

	simple_variable++

	fmt.Println(simple_variable, simple_variable_copy)

	// A pointer is a memory reference
	var simple_variable2 int = 100
	var simple_pointer *int = &simple_variable2
	fmt.Println(simple_variable2, simple_pointer) // dereferencing
	simple_variable2++
	fmt.Println(simple_variable2, simple_pointer)

}
