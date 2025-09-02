// kv_store: A simple TCP server that responds to any client request with '+OK' (like a basic Redis PONG)
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Print startup message
	fmt.Println("Listening on Port 6379")

	// Start a TCP server listening on port 6379
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Accept a single client connection
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Ensure connection is closed when main exits
	defer conn.Close()

	// Main loop: handle client requests
	for {
		// Allocate buffer for incoming data
		buf := make([]byte, 1024)

		// Read message from client
		_, err := conn.Read(buf)
		if err != nil {
			// If client closed connection, exit loop
			if err == io.EOF {
				break
			}
			// Print error and exit if read fails
			fmt.Println("Error reading from client: ", err.Error())
			os.Exit(1)
		}

		// Ignore request and send back a simple OK response
		conn.Write([]byte("+OK\r\n"))
	}

}
