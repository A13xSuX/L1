package main

import "fmt"

func itype(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Это int")
	case string:
		fmt.Println("Это string")
	case bool:
		fmt.Println("Это bool")
	case chan int:
		fmt.Println("Это chan int")
	case chan string:
		fmt.Println("Это chan string")
	case chan bool:
		fmt.Println("Это chan bool")
	case chan interface{}:
		fmt.Println("Это chan interface{}")
	default:
		fmt.Print("Неизвестное значение")
	}

}

func main() {
	itype(10)
	itype("Hello")
	itype(true)
	var ch chan int
	itype(ch)
	itype(2.4)
}
