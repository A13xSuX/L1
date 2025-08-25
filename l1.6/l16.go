package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// exit by condition
func workersByCondition(condition int) {
	for i := 0; i < condition; i++ {
		fmt.Println("goroutines work")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("goroutines finished by condition")
}

// exit by notification
func workersByNotification(stop <-chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("goroutines stopped")
			return
		default:
			fmt.Println("goroutines work")
			time.Sleep(1 * time.Second)
		}
	}
}

// exit by context
func workersByContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Print("goroutines finished by context")
			return
		default:
			fmt.Println("goroutines work")
			time.Sleep(1 * time.Second)
		}

	}
}

// exit by Goexit
func workersByGoexit() {
	defer fmt.Println("goroutines finished by Goexit")
	for i := 0; i < 3; i++ {
		fmt.Printf("goroutines work\n")
		time.Sleep(1 * time.Second)
		if i == 3 {
			runtime.Goexit()
		}
	}
}

// exit by close channel
func workerByChannelClose(stop <-chan struct{}) {
	for {
		select {
		case _, ok := <-stop:
			if !ok {
				fmt.Println("goroutines stop")
				return
			}
		default:
			fmt.Println("goroutines work")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// stop := make(chan bool)        //exit by notification
	// go workersByNotification(stop) // exit by notification
	// time.Sleep(3 * time.Second)    // exit by notification
	// stop <- true                   // exit by notification
	// go workersBy_Condition(5)//exit by condition

	// ctx, cancel := context.WithCancel(context.Background())//exit by context
	// go workersByContext(ctx)//exit by context
	// go workersByGoexit()//exit by Goexit
	stop := make(chan struct{})
	go workerByChannelClose(stop)
	time.Sleep(5 * time.Second)
	close(stop)
	time.Sleep(100 * time.Millisecond)
	// cancel()//exit by context
	fmt.Println("Program is finished")
}
