package main

import (
	"fmt"
	"time"
)

func produce(ch_p chan<- string, done <-chan struct{}) {
	defer close(ch_p) //чтобы избежать случая, получив сигнал done в recieve
	for i := 0; ; i++ {
		select {
		case <-done:
			fmt.Println("Producer: получен сигнал завершения")
			return
		default:
			ch_p <- fmt.Sprintf("Передано сообщение %d", i)
			time.Sleep(400 * time.Millisecond)
		}

	}
}

func recieve(ch <-chan string, done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Receiver: получен сигнал завершения")
			return
		case data, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("Получено сообщение %s\n", data)
		}

	}
}

func main() {
	timeout := 3 * time.Second
	done := make(chan struct{})
	ch_producer := make(chan string)

	time.AfterFunc(timeout, func() {
		close(done)
	})

	go produce(ch_producer, done)
	go recieve(ch_producer, done)

	<-done
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Конец")
}
