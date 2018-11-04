package main

import "fmt"

func great(args ...int) int {
	greatest := 0
	for _, v := range args {
		if greatest < v {
			greatest = v
		}
	}
	return greatest
}

func main() {
	fmt.Println(great(1, 2, 30, 4, 5, 6, 7))
}
