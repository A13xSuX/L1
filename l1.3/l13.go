package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func worker(id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Воркер %d получил: %s\n", id, data)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup
	var workers int
	fmt.Print("Введите число желаемых воркеров: ")
	fmt.Fscan(os.Stdin, &workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}
	counter := 1
	for {
		message := fmt.Sprintf("Сообщение %d", counter)
		ch <- message
		fmt.Printf("Главная горутина отправила: %s\n", message)
		counter++
		time.Sleep(500 * time.Millisecond)

	}
	close(ch) //не дойдет сюда(это проблема или нет?)
	wg.Wait()

}
