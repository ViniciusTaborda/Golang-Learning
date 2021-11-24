package main

import (
	"errors"
	"fmt"
)

func main() {
	var number_int16 int16 = 100
	fmt.Println(number_int16)

	var number_uint16 uint16 = 100
	fmt.Println(number_uint16)

	//alias
	// Rune is the same that int32
	var number_rune rune = 123456
	fmt.Println(number_rune)

	// Byte is the same that uint8
	var number_byte byte = 64
	fmt.Println(number_byte)

	var number_float float32 = 64.00001
	fmt.Println(number_float)

	var simple_string string = "string"
	fmt.Println(simple_string)

	char := 'A'
	fmt.Println(char)

	var text int
	fmt.Println(text)

	var bool_variable bool
	fmt.Println(bool_variable)

	var some_error error = errors.New("Internal error.")
	fmt.Println(some_error)
}
