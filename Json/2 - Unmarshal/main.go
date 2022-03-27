package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	// Name string `json:"-"` Using a `-` instead the name of the variable will hide it from the resulting JSON.
	Race string `json:"race"`
	Age  uint   `json:"age"`
}

func main() {

	dogJSON := `{"name":"Rex","race":"Yorkshire","age":4}`

	var d dog

	if err := json.Unmarshal([]byte(dogJSON), &d); err != nil {
		log.Fatal()
	}

	fmt.Println(d)

	dogJSON2 := `{"name":"Toddy","race":"Poodle"}`
	d2 := make(map[string]string)

	if err := json.Unmarshal([]byte(dogJSON2), &d2); err != nil {
		log.Fatal()
	}

	fmt.Println(d2)

}
