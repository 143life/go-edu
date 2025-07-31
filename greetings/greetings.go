package main

import (
	"fmt"

	"example.com/hello"
)

func main() {
	// Go to hello message and print it
	message := hello.Hello("Vika")
	fmt.Println(message)
	//fmt.Println(os.Getenv("GOPATH"))
}
