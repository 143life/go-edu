package goroutines

import "fmt"

func firstFibonacciNumbers() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 15; i++ {
			fmt.Printf("%d, ", <-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for i := 0; ; i++ {
		select {
		case c <- x:
			x, y = y, y+x
		case <-quit:
			fmt.Println("Exit")
			return
		}
	}
}
