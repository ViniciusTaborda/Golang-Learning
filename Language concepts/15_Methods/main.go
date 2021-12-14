package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving the user %s. \n", u.name)
}
func (u user) isMinor() bool {
	return u.age < 18
}
func (u *user) getOlder() {
	u.age++
}

func main() {
	fmt.Println("Methods!")
	user_1 := user{"Vinicius Taborda", 17}
	fmt.Println(user_1)
	user_1.save()
	fmt.Println(user_1.isMinor(), user_1.age)
	user_1.getOlder()
	fmt.Println(user_1.isMinor(), user_1.age)
}
