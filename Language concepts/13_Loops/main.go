package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops!")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	names := [3]string{"João", "Davi", "Lucas"}
	for index, name := range names {
		fmt.Println(index, name)
	}
	for index, char := range "LOTS OF WORDS" {
		fmt.Println(index, string(char))
	}

	user := map[string]string{
		"name":      "Vinicius",
		"last_name": "Costa",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

	for {
		fmt.Print(" Running forever ")
	}

}
