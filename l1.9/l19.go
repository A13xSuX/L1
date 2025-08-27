package main

import (
	"fmt"
	"sync"
)

func in(nums []int, firstCh chan<- int) {
	defer close(firstCh)
	for i := range nums {
		// fmt.Printf("Data sent %d\n", nums[i])
		firstCh <- nums[i]

	}
}

func multiplication(secondCh chan<- int, firstCh <-chan int) {
	defer close(secondCh)
	for num := range firstCh {
		secondCh <- num * 2
	}
}

func get(ch <-chan int) {
	for data := range ch {
		fmt.Printf("data get %d\n", data)
	}
}

func main() {
	var wg sync.WaitGroup
	nums := []int{}
	firstCh := make(chan int)
	secondCh := make(chan int)
	wg.Add(3)
	for i := 1; i <= 10; i++ {
		nums = append(nums, i)
	}
	go func() {
		defer wg.Done()
		in(nums, firstCh)
	}()
	go func() {
		defer wg.Done()
		multiplication(secondCh, firstCh)
	}()
	go func() {
		defer wg.Done()
		get(secondCh)
	}()

	wg.Wait()
}
