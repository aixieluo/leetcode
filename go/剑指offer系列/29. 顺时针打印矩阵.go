package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{6,7}}))
	fmt.Println(spiralOrder([][]int{{1,2,3}, {4,5,6}, {7,8,9}}))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	ans := make([]int, 0)
	for i := 0; i < len(matrix[0]); i++ {
		ans = append(ans, matrix[0][i])
	}
	if len(matrix) > 1 {
		for i := 1; i < len(matrix); i++ {
			ans = append(ans, matrix[i][len(matrix[0])-1])
		}
		if len(matrix[0]) > 1 {
			for i := len(matrix[0]) - 2; i >= 0; i-- {
				ans = append(ans, matrix[len(matrix)-1][i])
			}
			for i := len(matrix) - 2; i > 0; i-- {
				ans = append(ans, matrix[i][0])
			}
		}
	}
	if len(matrix) < 3 || len(matrix[0]) < 3 {
		return ans
	}
	matrix = matrix[1 : len(matrix)-1]
	for k := range matrix {
		matrix[k] = matrix[k][1 : len(matrix[k])-1]
	}
	return append(ans, spiralOrder(matrix)...)
}
