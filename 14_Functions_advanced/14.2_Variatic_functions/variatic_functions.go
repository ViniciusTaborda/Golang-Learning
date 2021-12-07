package main

import "fmt"

func sum(numbers ...int) (total int) {
	for _, number := range numbers {
		total += number
	}
	return
}

func main() {
	fmt.Println("Variatic functions")
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1))
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 2, 2, 2, 2, 2, 2))

}
