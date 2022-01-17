package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const connType = "tcp"

func main() {
	var connHost string
	var connPort string
	connHost, connPort = getHostandPort(os.Args)
	fmt.Printf("Connecting %s:%s over %s\n", connHost, connPort, connType)

	// dial will connect to the server
	conn, err := net.Dial(connType, connHost+":"+connPort)
	defer conn.Close()
	if err != nil {
		fmt.Println("Not able to connect to ", connHost, "at port", connPort)
		panic(err.Error())
	}

	// run loop forever, until exit.
	for {
		fileName := "./sample_input.txt"
		file, err := os.Open(fileName)
		defer file.Close()
		if err != nil {
			fmt.Println("Not able to read file", fileName)
			break
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var b = make([]byte, 3, 6)

			// We add \n because scanner.Text() removes the ending newline character
			conn.Write([]byte(scanner.Text() + "\n"))

			// Wait for the server message to indicate that the line is written
			conn.Read(b)
		}
		break
	}
}

func getHostandPort(args []string) (string, string) {
	var host string
	var port string
	if len(os.Args) <= 2 {
		fmt.Println("Host and port not specified")
		host = "localhost"
		port = "8080"
	} else {
		host = os.Args[1]
		port = os.Args[2]
	}
	return host, port
}
