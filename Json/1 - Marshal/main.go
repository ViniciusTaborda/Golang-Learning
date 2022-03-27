package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age  uint   `json:"age"`
}

func main() {
	d := dog{"Rex", "yorkshire", 4}
	fmt.Println(d)

	dogJSON, err := json.Marshal(d)

	if err != nil {
		log.Fatal(err)
	}

	// Encondes the dog struct to json format in bytes
	fmt.Println(dogJSON)

	// Decodes the json from bytes to human readable string
	fmt.Println(bytes.NewBuffer(dogJSON))

	d2 := map[string]string{
		"name": "Toddy",
		"race": "Poodle",
	}

	dog2JSON, err := json.Marshal(d2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dog2JSON)

	fmt.Println(bytes.NewBuffer(dog2JSON))

}
