package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Sending: Message %d\n", i)
		ch <- fmt.Sprintf("Message %d", i)
	}
	close(ch)
}

func consumer(ch <-chan string) {
	for msg := range ch {
		fmt.Println("Received:", msg)
		time.Sleep(2 * time.Second) // Simulate processing time
	}
}

func main() {
	ch := make(chan string, 3) // Buffered channel with a capacity of 3
	go producer(ch)
	consumer(ch)
}
