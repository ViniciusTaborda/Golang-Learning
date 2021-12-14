package main

import "fmt"

func math_calcs(n1, n2 int) (sum int, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return
}

func main() {
	fmt.Println("Named returns")
	sum, subtraction := math_calcs(10, 20)
	fmt.Println(sum, subtraction)

}
