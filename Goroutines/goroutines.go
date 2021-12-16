package main

import (
	"fmt"
	"time"
)

func printsForever(text string) {
	for {
		fmt.Println(text)
		// time.Sleep(time.Second)
	}
}

func main() {
	//concurrency != parallelism
	defer fmt.Println("Final")
	defer time.Sleep(time.Second)
	go printsForever("Hello...")
	go printsForever("world!")

}
