package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplex(write("Hello "), write("world!"))

	for i := 0; i < 15; i++ {
		fmt.Println(<-channel)
	}
}

func multiplex(inputChannel1, inputChannel2 <-chan string) chan string {

	outputChannel := make(chan string)
	go func() {
		for {
			select {
			case message := <-inputChannel1:
				outputChannel <- message
			case message := <-inputChannel2:
				outputChannel <- message
			}
		}
	}()
	return outputChannel
}

func write(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {

			channel <- fmt.Sprintf("Received value: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)))
		}
	}()

	return channel

}
