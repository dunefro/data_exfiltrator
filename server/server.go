package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	connHost = "localhost"
	// connPort = "8080"
	connType = "tcp"
)

func main() {

	// to get the port
	var connPort string
	if len(os.Args) <= 1 {
		fmt.Println("No port was specified, using default 8080")
		connPort = "8080"
	} else {
		connPort = os.Args[1]
	}

	fmt.Printf("Starting %s server on %s:%s\n", connType, connHost, connPort)
	conn, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Connection error", connHost+":"+connPort)
		panic(err.Error())
	}
	defer conn.Close()

	// running the loop for listening all the connections
	fmt.Println("Listening ... ")
	for {
		// Start accepting the connections
		client, err := conn.Accept()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Client", client.RemoteAddr().String(), "connected")
		go handleClientConnection(client)
	}
}

func handleClientConnection(conn net.Conn) {
	// handling buffer writes
	// it take the connection and then creates the buffer
	file, err := os.Create("./sample_output.txt")
	if err != nil {
		panic(err)
	}
	defer close(file)
	for {
		buffer, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			fmt.Println("Client left")
			conn.Close()
			// writer.Flush()
			return
		}
		file.WriteString(string(buffer[:]))
		fmt.Printf("%T", []byte("hello"))

		// Sending a reply back to client for synchronous connection
		conn.Write([]byte("Y\n"))
	}

}
func close(file *os.File) {
	fmt.Println("Closing the file")
	file.Close()
}
