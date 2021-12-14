package main

import "fmt"

func invertedNumber(number int) int {
	return number * -1
}
func invertNumber(number *int) {
	*number = *number * -1
}

func main() {
	fmt.Println("Pointers.")
	number := 20

	inverted_number := invertedNumber(number)
	fmt.Println(inverted_number, number)

	another_number := 10
	invertNumber(&another_number)
	fmt.Println(another_number)
}
