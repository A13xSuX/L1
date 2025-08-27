package main

import "fmt"

func main() {
	x, y := 2, 5 //(x,y)
	x ^= y       // (x^y,y)
	y ^= x       //(x^y,y^x^y)==>(x^y,x)
	x ^= y       //(x^y^x,x)==>(y,x)
	fmt.Println(x, y)
}
