package main

/*
The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).
*/

import "fmt"

func fib(x uint) uint {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return fib(x-1) + fib(x-2)
	}
}

func main() {
	fmt.Println(fib(5))
	fmt.Println(fib(10))
}
