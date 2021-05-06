// 给你一个 m x n 的矩阵 matrix 和一个整数 k ，找出并返回矩阵内部矩形区域的不超过 k 的最大数值和。
//
// 题目数据保证总会存在一个数值和不超过 k 的矩形区域。
//
//  
//
// 示例 1：
//
//
// 输入：matrix = [[1,0,1],[0,-2,3]], k = 2
// 输出：2
// 解释：蓝色边框圈出来的矩形区域 [[0, 1], [-2, 3]] 的数值和是 2，且 2 是不超过 k 的最大数字（k = 2）。
// 示例 2：
//
// 输入：matrix = [[2,2,-1]], k = 3
// 输出：3
//  
//
// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -100 <= matrix[i][j] <= 100
// -105 <= k <= 105
//  
//
// 进阶：如果行数远大于列数，该如何设计解决方案？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSumSubmatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2))
	fmt.Println(maxSumSubmatrix([][]int{{2, 2, -1}}, 3))
	fmt.Println(maxSumSubmatrix([][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}, 10))
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	maxVal := math.MinInt32
	xl, yl := len(matrix[0]), len(matrix)
	bl := min(len(matrix), len(matrix[0]))
	points := make([][2][2]int, 0, 10)
	for b := 0; b < bl; b++ {
		x:= 0
		for ; x+b < xl ; x++ {
			y:=0
			for ; y+b < yl; y++ {
				point := [2][2]int{{x,y}, {x+b, y+b}}
				points = append(points, point)
			}
		}
	}
	for _,v :=range points {
		sum:= 0
		for y:=v[0][1]; y<= v[1][1]; y++ {
			for x:= v[0][0]; x<= v[0][0]; x++ {
				sum+= matrix[y][x]
			}
		}
		if sum <= k {
			maxVal = max(maxVal, sum)
		}
	}

	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
