package main

import "fmt"

func main() {
	fmt.Println("Control structures")

	some_number := 10
	if some_number > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Less or equals 15")
	}

	if another_number := some_number; another_number > 0 {
		fmt.Println("Another number is greater than zero")
	} else if another_number := some_number; another_number = 1 {
		fmt.Println("Another number is one")
	}

}
