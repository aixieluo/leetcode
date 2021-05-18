package main

import (
	"fmt"
)

func main() {
	// fmt.Println(math.Pow(34.00515, -3))
	fmt.Println(myPow(34.00515, -3))
	fmt.Println(myPow(2, 10))
	fmt.Println(myPow(2, -2))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1/x
		n = -n
	}
	res := myPow(x, n/2)
	if n %2 == 1 {
		return  x *  res * res
	} else {
		return res * res
	}
}
