package main

import (
	"fmt"
	"sync"
	"time"
)

func sleep(duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	<-time.After(duration) //канал заблокируется пока не получит сигнал и не прочтет его()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println(time.Now())
	go sleep(5*time.Second, &wg)
	wg.Wait()
	fmt.Println(time.Now())
}
