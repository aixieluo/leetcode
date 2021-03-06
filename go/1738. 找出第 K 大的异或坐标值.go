// 给你一个二维矩阵 matrix 和一个整数 k ，矩阵大小为 m x n 由非负整数组成。
//
// 矩阵中坐标 (a, b) 的 值 可由对所有满足 0 <= i <= a < m 且 0 <= j <= b < n 的元素 matrix[i][j]（下标从 0 开始计数）执行异或运算得到。
//
// 请你找出 matrix 的所有坐标中第 k 大的值（k 的值从 1 开始计数）。
//
//  
//
// 示例 1：
//
// 输入：matrix = [[5,2],[1,6]], k = 1
// 输出：7
// 解释：坐标 (0,1) 的值是 5 XOR 2 = 7 ，为最大的值。
// 示例 2：
//
// 输入：matrix = [[5,2],[1,6]], k = 2
// 输出：5
// 解释：坐标 (0,0) 的值是 5 = 5 ，为第 2 大的值。
// 示例 3：
//
// 输入：matrix = [[5,2],[1,6]], k = 3
// 输出：4
// 解释：坐标 (1,0) 的值是 5 XOR 1 = 4 ，为第 3 大的值。
// 示例 4：
//
// 输入：matrix = [[5,2],[1,6]], k = 4
// 输出：0
// 解释：坐标 (1,1) 的值是 5 XOR 2 XOR 1 XOR 6 = 0 ，为第 4 大的值。
//  
//
// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 1000
// 0 <= matrix[i][j] <= 106
// 1 <= k <= m * n
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-kth-largest-xor-coordinate-value

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kthLargestValue([][]int{{8,10,5,8,5,7,6,0,1,4,10,6,4,3,6,8,7,9,4,2}}, 2))
	fmt.Println(kthLargestValue([][]int{{5,2}, {1,6}}, 3))
}

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, n)
	}
	nums := make([]int, 0, m*n)
	for y := 0; y < m; y++{
		for x :=0; x < n; x++ {
			dp[y][x] = matrix[y][x]
			if y > 0 {
				dp[y][x] ^= dp[y-1][x]
			}
			if x > 0 {
				dp[y][x] ^= dp[y][x-1]
			}
			if x >0 && y > 0 {
				dp[y][x] ^= dp[y-1][x-1]
			}
			nums = append(nums, dp[y][x])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums[k-1]
}
