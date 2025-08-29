package main

import (
	"fmt"
)

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left := []int{}
	equal := []int{}
	right := []int{}

	elem := nums[len(nums)/2]

	for _, num := range nums {
		if num < elem {
			left = append(left, num)
		} else if elem == num {
			equal = append(equal, num)
		} else {
			right = append(right, num)
		}
	}
	return append(append(quicksort(left), equal...), quicksort(right)...)
}

func main() {
	nums := []int{3}
	fmt.Println(quicksort(nums))
	//test on new pc
}
