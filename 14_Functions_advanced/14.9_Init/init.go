package main

import "fmt"

func main() {
	fmt.Println("Main function running.")
}

// The init function will run before the main function
func init() {
	fmt.Println("Init function running.")
}
