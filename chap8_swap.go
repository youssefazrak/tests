package main

import "fmt"

func swap(x,y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	x := 14
	y := 2
	fmt.Println("x and y before the swap: ", x, " ", y)
	swap(&x, &y)
	fmt.Println("x and y after the swap: ",x, " ", y)
}
