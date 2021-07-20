package main

import (
	"fmt"
	"time"
)

func testChannels() {
	messages := make(chan string)
	
	go func() {
		messages <- "A piece of message."
		messages <- "Another piece of message."
	}()

	msg := <- messages
	fmt.Println("Message:", msg)
	msg = <- messages
	fmt.Println("Message:", msg)
}

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(5 * time.Second)
	fmt.Println("Done")

	done <- true
}

func testChannelSync() {
	done := make(chan bool)
	go worker(done)
	fmt.Println("Wating...")
	<- done
	fmt.Println("main finished.")
}

func main() {
	testChannels()
	testChannelSync()
}
