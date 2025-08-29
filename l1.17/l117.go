package main

import (
	"fmt"
	"slices"
)

func binarySearch(arr []int, num int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == num {
			return mid
		} else if arr[mid] < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1

}

func main() {
	arr := []int{2, 3, 5, 6, 8}
	num := 4
	slices.Sort(arr)
	fmt.Printf("Searched element %d", binarySearch(arr, num))
}
