package main

import "fmt"

func deleteEl(arr []int, el int) []int {
	result := make([]int, len(arr)-1)
	copy(result, arr[:el])
	copy(result[el:], arr[el+1:])
	return result
}

func main() {
	source := []int{1, 2, 3, 4, 5}
	el := 2
	fmt.Print(deleteEl(source, el))
}
