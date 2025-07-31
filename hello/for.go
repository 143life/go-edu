package hello

import "fmt"

func for_statement() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// while
	/*
		C++ mentioned
	*/
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
