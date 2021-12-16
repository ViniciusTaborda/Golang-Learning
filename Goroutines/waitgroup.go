package main

import (
	"fmt"
	"sync"
	"time"
)

func printsFiveTimes(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(4)

	go func() {
		printsFiveTimes("Hello...")
		waitGroup.Done()
	}()
	go func() {
		printsFiveTimes("world!")
		waitGroup.Done()
	}()
	go func() {
		printsFiveTimes("world!1")
		waitGroup.Done()
	}()
	go func() {
		printsFiveTimes("world!2")
		waitGroup.Done()
	}()
	waitGroup.Wait()

}
