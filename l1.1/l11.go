package main

import "fmt"

type human struct {
	name, lastname string
	height, weight float64 //height in metres,weight in kg
}
type action struct {
	human
	time uint //in second
}

func (h human) lengthOfStep() float64 {
	return h.height/4 + 0.37
}

func (h human) heightOfJump() float64 {
	return h.weight / 2
}

func (a action) walk() float64 {
	return a.lengthOfStep() * float64(a.time)
}

func main() {
	ac := action{
		human: human{
			name:     "Vasya",
			lastname: "Pupkin",
			height:   1.8,
			weight:   80,
		},
		time: 3600,
	}

	fmt.Printf("Length of your step: %v m.\n", ac.lengthOfStep())
	fmt.Printf("Height of your jump: %v sm.\n", ac.heightOfJump())
	fmt.Printf("You walked: %v m.\n", ac.walk())
	fmt.Printf("My name is %s \n", ac.name)
}
