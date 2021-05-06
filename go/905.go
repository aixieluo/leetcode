package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(A []int) []int {
	var a, b []int
	a, b = make([]int, 0), make([]int, 0)
	for _, v:= range A {
		if v%2 ==0 {
			a = append(a, v)
		} else {
			b = append(b, v)
		}
	}
	return append(a, b...)
}
