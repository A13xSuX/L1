package main

import "fmt"

func uniq(str []string) []string {
	var result []string
	set := make(map[string]bool)
	for _, s := range str {
		set[s] = true
	}

	for s := range set {
		result = append(result, s)
	}
	return result
}

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(uniq(str))
}
