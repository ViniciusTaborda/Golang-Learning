package main

import (
	"fmt"
	"time"
)

func printsFiveTimes(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}
	close(channel)

}

func main() {
	channel := make(chan string)
	go printsFiveTimes("Hello world!", channel)

	// for {
	// 	message, isOpen := <-channel

	// 	if !isOpen {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }
	for message := range channel {
		fmt.Println(message)
	}

}
