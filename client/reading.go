package main

import (
	"os"
)

func main() {
	file, err := os.Create("./write.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("I am doing good")
	// fmt.Printf("Type: %T\n", data)
	// fmt.Printf("Length: %d\n", len(data))
	// fmt.Println(string(data))
}
