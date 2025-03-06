package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(400 * time.Millisecond)
	}
	channel <- "Done"
}

func main() { //  main goroutine
	channel := make(chan string)
	go printMessage("Go is great", channel) // open a new goroutine
	// go printMessage("FM rocks")
	// printMessage("We miss classes") // we need to keep main goroutine busy to see results of others
	response := <-channel // waiting for the channel value
	fmt.Println(response)
}
