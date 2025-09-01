package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(other *Point) float64 {
	distance := math.Sqrt(math.Pow((other.x-p.x), 2) + math.Pow((other.y-p.y), 2))
	return distance
}
func main() {
	point := NewPoint(1, 3)
	other := NewPoint(2, 3)
	fmt.Println(point.Distance(other))
}
