package main

import (
	"fmt"
	"math/big"
)

func main() {
	num1 := new(big.Float)
	num2 := new(big.Float)
	num1.SetString("100000000000000000000.0")
	num2.SetString("3.0")
	num1.SetPrec(100) //бит точности
	num2.SetPrec(100)
	sum := new(big.Float).Add(num1, num2)
	product := new(big.Float).Mul(num1, num2)
	difference := new(big.Float).Sub(num1, num2)
	division := new(big.Float).Quo(num1, num2)
	fmt.Printf("Сложение: %v + %v = %v\n", num1, num2, sum)
	fmt.Printf("Умножение: %v * %v = %v\n", num1, num2, product)
	fmt.Printf("Вычитание: %v - %v = %v\n", num1, num2, difference)
	fmt.Printf("Деление: %v / %v = %v\n", num1, num2, division)
}
