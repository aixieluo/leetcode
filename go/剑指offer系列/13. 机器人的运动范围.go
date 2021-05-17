// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
//
//  
//
// 示例 1：
//
// 输入：m = 2, n = 3, k = 1
// 输出：3
// 示例 2：
//
// 输入：m = 3, n = 1, k = 0
// 输出：1
// 提示：
//
// 1 <= n,m <= 100
// 0 <= k <= 20
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof

package main

import "fmt"

func main() {
	fmt.Println(movingCount(16, 8, 4))
	fmt.Println(movingCount(2, 3, 1))
}

func movingCount(m int, n int, k int) int {
	ans := 1
	dp := make([][]int, m)
	for k:=range dp {
		dp[k] = make([]int, n)
	}
	dp[0][0] = 1
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			kVal := 0
			row1 := row
			for row1 > 0 {
				kVal += row1 % 10
				row1 /= 10
			}
			col1 := col
			for col1 > 0 {
				kVal += col1 % 10
				col1 /= 10
			}
			if kVal <= k && ((col > 0 && dp[row][col-1] == 1) || (row > 0 && dp[row-1][col] == 1)) {
				dp[row][col] = 1
				ans++
			}
		}
	}
	return ans
}
