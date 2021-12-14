package main

import (
	"first_module/helper"
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from the main file.")
	helper.Write()
	err := checkmail.ValidateFormat("vinicius@gmail.com")
	fmt.Println(err)
}
