package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func main() {
	c:= Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r)

}

func circleArea(c Circle) float64 {
	return math.Pi * c.r*c.r
}
