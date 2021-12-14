package main

import "fmt"

func closure() func() {
	text := "Inside closure function."
	internal_function := func() {
		fmt.Println(text)
	}
	return internal_function
}

func main() {
	fmt.Println("Closure!")
	text := "Inside main function."
	fmt.Println(text)

	new_function := closure()
	new_function()
}
