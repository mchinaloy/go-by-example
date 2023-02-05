package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	go func() {
		messages <- "ping"
	}()
	msg := <-messages
	fmt.Println(msg)

	make_channel_with_buffer()
	sync()
}

func make_channel_with_buffer() {
	messages := make(chan string, 2)
	messages <- "Hello"
	messages <- "World"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func sync() {
	messages := make(chan bool)
	go worker(messages)
	<-messages
}

func worker(messages chan bool) {
	fmt.Println("Working...")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Done")
	messages <- true
}
