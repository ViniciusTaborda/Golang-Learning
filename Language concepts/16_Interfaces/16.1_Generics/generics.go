package main

import "fmt"

func genericFunction(interf interface{}) interface{} {
	fmt.Println(interf)
	return interf
}

func main() {

	genericFunction("String")
	genericFunction(100)
	genericFunction(true)
	genericFunction(nil)
	genericFunction(23.2)
	genericFunction("")
}
