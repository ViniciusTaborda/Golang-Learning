package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and slices")

	var test_array [10]string
	some_array := [5]string{"1", "2", "3"}

	test_array[2] = "Strnig string"

	fmt.Println(test_array)
	fmt.Println(some_array)

	another_array := [...]int{1, 2, 3, 4, 5}
	fmt.Println(another_array)

	some_slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(some_slice)

	fmt.Println(reflect.TypeOf(some_array))
	fmt.Println(reflect.TypeOf(some_slice))

	some_slice = append(some_slice, 999)
	fmt.Println(some_slice)

	slice_of_array := another_array[2:4]
	fmt.Println(slice_of_array)

	// Internal arrays
	fmt.Println("Internal arrays")

	slice := make([]float32, 10, 15)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
