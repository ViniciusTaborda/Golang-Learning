package main

import "fmt"

func main() {
	channel := make(chan string, 2) // channel with buffer of size two.
	channel <- "Hello world!"
	channel <- "Hello world2!"
	message := <-channel
	message2 := <-channel
	fmt.Println(message, message2)
}
