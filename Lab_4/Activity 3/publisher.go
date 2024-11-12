package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subject := "updates"

	for i := 1; i <= 10; i++ {
		message := fmt.Sprintf("Hello, NATS! Message #%d", i)
		err := nc.Publish(subject, []byte(message))
		if err != nil {
			log.Printf("Error publishing message: %v", err)
		} else {
			fmt.Println("Sent:", message)
		}
		time.Sleep(500 * time.Millisecond)
	}

	err = nc.Flush()
	if err != nil {
		log.Fatalf("Error flushing connection: %v", err)
	}

	if err := nc.LastError(); err != nil {
		log.Fatalf("NATS encountered an error: %v", err)
	}

	fmt.Println("All messages published successfully.")
}
