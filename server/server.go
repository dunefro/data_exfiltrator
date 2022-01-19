package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	connType = "tcp"
)

func Serve(fileName, connHost, connPort string) error {

	fmt.Printf("Starting %s server on %s:%s\n", connType, connHost, connPort)
	conn, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		return fmt.Errorf("ConnectionError: Not able to connect %s", connHost+":"+connPort)
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
		go handleClientConnection(client, fileName)
		fmt.Println("You can press Ctrl+c to terminate the program")
	}
}

func handleClientConnection(conn net.Conn, fileName string) {
	// handling buffer writes
	// it take the connection and then creates the buffer
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer close(file)
	for {
		buffer, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			fmt.Println("Client left")
			conn.Close()
			return
		}
		file.WriteString(string(buffer[:]))

		// Sending a reply back to client for synchronous connection
		conn.Write([]byte("Y\n"))
	}

}
func close(file *os.File) {
	fmt.Println("Closing the file")
	fmt.Println()
	fmt.Println("Listening ... (press Ctrl+c to terminate)")
	file.Close()
}
