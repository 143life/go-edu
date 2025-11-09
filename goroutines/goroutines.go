package goroutines

import (
	"fmt"
	"time"
)

func StartGoroutines() {
	// 001 example
	//go sleepyGopher()           // начало горутины
	//time.Sleep(4 * time.Second) // Ожидание храпа гофера

	// 002 example
	//for i := 0; i < 5; i++ {
	//	go sleepyGopher(i)
	//}
	//time.Sleep(4 * time.Second)

	// 003 example
	//channel := make(chan int) // Делает канал для связи
	//for i := 0; i < 5; i++ {
	//	go sleepyGopher(i, channel)
	//}
	//for i := 0; i < 5; i++ {
	//	gopherID := <-channel // получает значение от канала
	//	fmt.Println("gopher ", gopherID, "has finished sleeping")
	//}

	// WaitGroup
	// 004 example
	//hundredThereAndBackAgain()
	// 005 data race example
	// dataRace()

	// 006 Start stepik goroutines leraning
	// Start()

	// 007 Select-case construction
	firstFibonacciNumbers()
} // здесь все горутины останавливаются

// извращенно решаю проблему с отсутствием перегрузки методов
// используя kwargs для разных случаев ввода
func sleepyGopher(kwargs ...interface{}) {
	switch len(kwargs) {
	case 0:
		time.Sleep(3 * time.Second) // гофер спит
		fmt.Println("... snore ...")
	case 1:
		id := kwargs[0]
		time.Sleep(3 * time.Second)
		fmt.Println("... ", id, " snore ...")
	case 2:
		time.Sleep(3 * time.Second)
		fmt.Println("... ", kwargs[0], " snore ...")
		switch channel := kwargs[1].(type) {
		case chan int:
			if id, ok := kwargs[0].(int); ok {
				channel <- id
			} else {
				fmt.Print("error")
			}
		default:
			fmt.Print("error")
		}
	}
}
