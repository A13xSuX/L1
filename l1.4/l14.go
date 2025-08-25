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

func worker(id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Обрабочик %d получил сообщение: %s\n", id, data)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("Воркер %d завершил работу\n", id)
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	ch := make(chan string)
	var wg sync.WaitGroup
	var workers int
	fmt.Fscan(os.Stdin, &workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	go func() {
		<-signalChan
		fmt.Println("Канал закрыт по сигналу")
		cancel()
		close(ch)
	}()

	counter := 1
	for {
		message := fmt.Sprintf("Сообщение %d", counter)
		select {
		case <-ctx.Done():
			wg.Wait()
			fmt.Println("Все воркеры завершены")
			return
		default:
			select {
			case <-ctx.Done():
				continue
			case ch <- message:
				counter++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
}
