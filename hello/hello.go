package hello

import (
	"fmt"

	"rsc.io/quote"
)

func All() {
	fmt.Println("Hello, world!")
	fmt.Println(quote.Go())
	for_statement()
	fmt.Println(sqrt(2), sqrt(-4))
}

// Hello returns a greeting for the named person
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
