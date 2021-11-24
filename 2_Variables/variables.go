package main

import "fmt"

func main() {
	var string_one string = "String one."
	string_two := "String two"
	const string_three = "String three"
	var (
		string_four string = "String four"
		string_five string = "String five"
	)
	string_six, string_seven := "String six", "String seven"

	fmt.Println(string_one)
	fmt.Println(string_two)
	fmt.Println(string_three)
	fmt.Println(string_four)
	fmt.Println(string_five)
	fmt.Println(string_six)
	fmt.Println(string_seven)

	string_five, string_four = string_four, string_five

}
