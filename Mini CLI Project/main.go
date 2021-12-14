package main

import (
	"cli-project/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Web Searcher!")
	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
