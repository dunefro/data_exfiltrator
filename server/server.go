package main

import "fmt"

const (
	connHost = "localhost"
	connPort = 8080
	connType = "tcp"
)

func main() {
	fmt.Printf("Starting %s server on %s:%d", connType, connHost, connPort)
}
