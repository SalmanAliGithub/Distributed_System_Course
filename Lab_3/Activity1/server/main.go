package main

import (
	"bufio"
	"fmt"
	"net"
)

// Simple TCP server that listens for incoming connections and responds to messages
func main() {
	// Start the server and listen on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")

	for {
		// Accept a client1 connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the client1 connection in a new goroutine (allowing multiple clients)
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Read message from client1
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message received:", string(message))

	// Send response back to client1
	conn.Write([]byte("Message received: " + message))
}
