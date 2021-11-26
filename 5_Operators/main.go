package main

import "fmt"

func main() {
	//Arithmetics
	sum := 1 + 2
	substract := 1 - 2
	division := 1 / 2
	multiplication := 1 * 2
	mod_div := 1 % 2
	fmt.Println(sum, substract, division, multiplication, mod_div)

	// Assignment
	var variable string = "String"
	variable2 := "String2"
	fmt.Println(variable, variable2)

	// Relational
	fmt.Println(2 > 1)
	fmt.Println(1 < 2)
	fmt.Println(2 >= 1)
	fmt.Println(1 <= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 1)
	fmt.Println(1 != 2)

	// Logic
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)

	// Unitary
	number := 10
	number++
	number += 15

	fmt.Println(number)

	number--
	number -= 15
	fmt.Println(number)

	number *= 15
	fmt.Println(number)

	number /= 15
	fmt.Println(number)

	number %= 15
	fmt.Println(number)
}
