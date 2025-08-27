package main

import "fmt"

func distribution(m map[int][]float64, temp []float64) {
	for _, t := range temp {
		key := int(t/10) * 10
		m[key] = append(m[key], t)
	}
}

func main() {
	solution := make(map[int][]float64)
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 9.0, -9.0}
	distribution(solution, temp)
	fmt.Println(solution)
}
