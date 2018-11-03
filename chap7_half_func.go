package main

import "fmt"

//func sum(tot_sum []int) int {
//}

func halves(x int) (int, bool) {
	half := x / 2 
	if x % 2 == 0 {
		return half, true
	} else {
		return half, false
	}
}


func main() {
	fmt.Println(halves(10))
	fmt.Println(halves(20))
	fmt.Println(halves(7))
}
