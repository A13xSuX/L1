package main

import "fmt"

func main() {
	var num int64 = 5
	var bitPosition int
	var bitValue int
	fmt.Print("Индекс бита: ")
	fmt.Scan(&bitPosition)
	fmt.Print("Установить 0 или 1?")
	fmt.Scan(&bitValue)

	fmt.Printf("Исходное число: %d - %b\n", num, num)

	if bitValue == 1 {
		num |= 1 << bitPosition
	} else if bitValue == 0 {
		num &= ^(1 << bitPosition)
	} else {
		fmt.Print("Введите корректное значение бита(0/1)")
	}
	fmt.Printf("Результат: %d - %b\n", num, num)
}
