package goroutines

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(2)
}

func hundredThereAndBackAgain() {
	// create waiter
	var wg sync.WaitGroup

	// add to wg counter one goroutine
	wg.Add(1)
	// create first goroutine
	go func() {
		for i := 10; i >= 0; i-- {
			fmt.Printf("1 Goroutine: %d\n", i)
		}
		// Say to waiter that goroutine is ended
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Printf("2 goroutine: %d\n", i)
		}
		wg.Done()
	}()

	// wg waiting for all goroutines is ended
	wg.Wait()
}

// Race detection can be used using:
// go run -race .
// go build -race .
// go test - race.
// If something went wrong, check envs, and verify if mingv-64 (64!) installed
// and added to PATH on your machine.
func dataRace() {
	var counter int

	var wg sync.WaitGroup

	for g := 1; g <= 2; g++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 2; i++ {
				counter++
				fmt.Printf("Goroutine %d: counter = %d\n", g, counter)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
