package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {

	u := user{"Vinicius", "vinicius@gmail.com"}

	//Using dynamic data to render the html file.
	templates.ExecuteTemplate(w, "home.html", u)

}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	// Serves on port 5000
	println("Listening on port 5000.")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
