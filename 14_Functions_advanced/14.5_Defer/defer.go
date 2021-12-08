package main

import "fmt"

func function_one() {
	fmt.Println("Executing function one")
}
func function_two() {
	fmt.Println("Executing function two")
}

func main() {
	defer function_one()
	function_two()
}
