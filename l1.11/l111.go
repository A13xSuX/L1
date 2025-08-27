package main

import "fmt"

func intersection(a, b []int) []int {
	set := make(map[int]bool)
	var result []int

	for _, elema := range a {
		set[elema] = true
	}

	for _, elemb := range b {
		if set[elemb] {
			result = append(result, elemb)
			set[elemb] = false //for delete duplicates
		}
	}
	return result

}
func main() {
	a := []int{1, 2, 3, 5, 5, 5}
	b := []int{2, 3, 4, 5, 5}
	fmt.Println(intersection(a, b))
}
