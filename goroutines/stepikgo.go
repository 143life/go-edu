package goroutines

import (
	"fmt"
	"time"
)

func Start() {
	go myFunc()

	// Самовызывающаяся анонимная горутина
	go func() {
		fmt.Println("Privet, ya anonymous")
	}()

	c := make(chan int, 1) // Канал с буфером размером 1
	fmt.Println(len(c), cap(c))
	c <- 1
	fmt.Println(len(c), cap(c))
	<-c

	// Буферизированные каналы - это очередь
	channel := make(chan string, 3)
	channel <- "a"
	channel <- "b"
	fmt.Println(<-channel)
	channel <- "c"
	fmt.Println(<-channel, <-channel)

	time.Sleep(1 * time.Second)
}

func myFunc() {
	fmt.Println("Privet")
}
