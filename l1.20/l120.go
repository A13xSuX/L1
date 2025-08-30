package main

import (
	"fmt"
	"strings"
)

func reverseSentence(str string) string {
	final := ""
	arr := strings.Split(str, " ")
	for i := len(arr) - 1; i >= 0; i-- {
		if i == 0 {
			final += arr[i]
		} else {
			final += arr[i] + " "
		}
	}
	return final
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverseSentence(str))
}
