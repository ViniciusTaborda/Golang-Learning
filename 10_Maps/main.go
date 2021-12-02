package main

import "fmt"

func main() {
	fmt.Println("Maps")
	user := map[int]string{
		5:  "Vinicius",
		10: "Costa",
	}

	fmt.Println(user)

	user_2 := map[string]map[string]string{
		"name": {
			"first": "Vinicius",
			"last":  "Costa",
		},
		"course": {
			"name": "Systems of information",
			"city": "Curitiba",
		},
	}
	fmt.Println(user_2)

	delete(user_2, "name")
	fmt.Println(user_2)

	user_2["birthday"] = map[string]string{
		"day":   "23",
		"month": "11",
	}
	fmt.Println(user_2)

}
