package main

import "fmt"

func reverse(s string) string {
	var s_rev string
	runeSl := []rune{}
	for _, r := range s {
		runeSl = append(runeSl, r)
	}
	for i := len(runeSl) - 1; i >= 0; i-- {
		s_rev += string(runeSl[i])
	}
	return s_rev
}

func main() {
	s := "главрыба"
	fmt.Printf("Начальное слово: %s\n Перевернутое: %s", s, reverse(s))

}
