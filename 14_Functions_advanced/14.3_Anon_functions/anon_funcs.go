package main

import "fmt"

func main() {
	func() {
		fmt.Println("I am a anonymous function!")
	}()

	anon_func_variable := func() string {
		return "I am a anonymous function in a variable!"
	}
	result := anon_func_variable()
	fmt.Println(result)

}
