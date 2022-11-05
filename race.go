package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// var mu sync.Mutex

//race Condition.
func main() {
	gs := 100
	counter := 0
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			v := counter
			v++
			time.Sleep(time.Microsecond)
			counter = v

			wg.Done()
		}()
		fmt.Println("Goroutine: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutine: ", runtime.NumGoroutine())
	fmt.Println(counter)

}
