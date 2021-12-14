package main

import "fmt"

type address struct {
	street_name   string
	street_number int16
}

type user struct {
	name         string
	age          uint8
	user_address address
}

func main() {

	var some_address address
	some_address.street_name = "Rua dos bobos"
	some_address.street_number = 0

	var some_user user
	some_user.age = 20
	some_user.name = "Vinicius Taborda"
	some_user.user_address = some_address
	fmt.Println(some_user)

	another_user := user{"Vinicius Taborda", 20, some_address}
	fmt.Println(another_user)

	fmt.Println(some_user == another_user)
}
