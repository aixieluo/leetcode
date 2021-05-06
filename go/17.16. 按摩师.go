// 一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。
//
// 注意：本题相对原题稍作改动
//
//  
//
// 示例 1：
//
// 输入： [1,2,3,1]
// 输出： 4
// 解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。
// 示例 2：
//
// 输入： [2,7,9,3,1]
// 输出： 12
// 解释： 选择 1 号预约、 3 号预约和 5 号预约，总时长 = 2 + 9 + 1 = 12。
// 示例 3：
//
// 输入： [2,1,4,5,3,1,1,3]
// 输出： 12
// 解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/the-masseuse-lcci

package main

import "fmt"

func main() {
	fmt.Println(massage([]int{2,7,9,3,1}))
}

func massage(nums []int) int {
	dp := make([][]int, len(nums)+1)
	n := len(nums)
	for k:=range dp {
		dp[k] = make([]int, 2)
	}
	// dp[n][0] = max(dp[n-1][1], dp[n-1][0])
	// dp[n][1] = dp[n-1][0] + nums[n-1]
	for i:=0 ; i<n; i++{
		dp[i+1][1] = dp[i][0] + nums[i]
		dp[i+1][0] = max(dp[i][1], dp[i][0])
	}
	return max(dp[n][1], dp[n][0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
