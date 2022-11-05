package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func foo(ctx context.Context, wg *sync.WaitGroup) {
	count := 0
	for {
		count++
		fmt.Println("Inside foo:", count)

		select {
		case <-ctx.Done():
			fmt.Println("ABORTED......")
			wg.Done()
			return
		}
	}
}

func bar(ctx context.Context, wg *sync.WaitGroup) {
	time.Sleep(10)
	fmt.Println("Inside bar")
	// select {
	// case <-ctx.Done():
	// 	fmt.Println("bar routine cancelled..")
	// case <-wg.Done():
	// 	fmt.Println("bar executed till end")
	// 	return
	// }
	wg.Done()
	return
}

func main() {
	var wg sync.WaitGroup

	//Gracefully termination of both the goroutines.
	Done := make(chan os.Signal, 1)
	signal.Notify(Done, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(2)
	go foo(ctx, &wg)
	go bar(ctx, &wg)

	select {
	// Shutdown. Cancel application context will kill all attached tasks.
	case <-Done:
		cancel()
		fmt.Println("Client received exit cmd..Shutting down all goroutines.")
	}

	wg.Wait()
	fmt.Println("FInsihed")
}
