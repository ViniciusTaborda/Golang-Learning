package main

import "fmt"

func recoverExecution() {
	// As this function is deferred in the main function, go will
	// try to recover if any panic occurs.
	fmt.Println("Trying to recover from panic!")
	if r := recover(); r != nil {
		fmt.Println("Successfullly recovered from panic...")
	}
}

func isApproved(n1, n2 float64) bool {
	mean := (n1 + n2) / 2
	if mean >= 0 && mean <= 6 {
		return false
	} else if mean > 6 {
		return true
	} else {
		panic("Unexpected error!!")
	}
}

func main() {
	fmt.Println("Panic and recover.")
	defer recoverExecution()
	fmt.Println(isApproved(6.2, 5.7))

}
