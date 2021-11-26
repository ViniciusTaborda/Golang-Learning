package main

import "fmt"

type person struct {
	name      string
	last_name string
	age       uint8
	height    float32
}
type student struct {
	person
	course     string
	university string
}

func main() {
	person1 := person{"Vinicius", "Taborda", 20, 1.87}
	student1 := student{person1, "Sistemas de informação", "PUC-PR"}
	fmt.Println(student1)
	fmt.Println(student1.name + " " + student1.last_name)

}
