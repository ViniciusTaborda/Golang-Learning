package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func main() {

	http.HandleFunc("/home", home)

	// Handle functions can be either nameless or pre-declared
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// The request return has to be a slice of bytes!
		w.Write([]byte("This is an users page!"))

	})

	// Serves on port 5000
	log.Fatal(http.ListenAndServe(":5000", nil))

}
