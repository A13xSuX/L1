package main

import (
	"fmt"
	"sync"
)

func sqaure(i int, nums []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(nums[i] * nums[i])
}

func main() {
	var wg sync.WaitGroup
	nums := []int{2, 4, 6, 8, 10}
	for i := range nums {
		wg.Add(1)
		go sqaure(i, nums, &wg)
	}
	wg.Wait()

}
