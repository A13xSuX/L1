package main

import (
	"fmt"
	"strings"
)

func duplicate(str string) bool {
	str = strings.ToLower(str)
	asci := make(map[rune]bool)
	for _, v := range str {
		if asci[v] {
			return false
		} else {
			asci[v] = true
		}
	}
	return true
}

func main() {
	str := "abCdef"
	fmt.Println(duplicate(str))
}
