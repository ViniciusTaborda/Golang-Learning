package main

import "fmt"

func fibonaccci(position uint) uint {
	if position <= 1 {
		return position
	}
	return fibonaccci(position-2) + fibonaccci(position-1)
}

func main() {
	fmt.Println("Recursive function!")
	position := uint(10)
	for i := uint(1); i < position; i++ {
		fmt.Println(fibonaccci(i))
	}
}
