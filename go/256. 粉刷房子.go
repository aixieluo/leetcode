// 假如有一排房子，共 n 个，每个房子可以被粉刷成红色、蓝色或者绿色这三种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。
//
// 当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n x 3 的矩阵来表示的。
//
// 例如，costs[0][0] 表示第 0 号房子粉刷成红色的成本花费；costs[1][2] 表示第 1 号房子粉刷成绿色的花费，以此类推。请你计算出粉刷完所有房子最少的花费成本。
//
// 注意：
//
// 所有花费均为正整数。
//
// 示例：
//
// 输入: [[17,2,17],[16,16,5],[14,3,19]]
// 输出: 10
// 解释: 将 0 号房子粉刷成蓝色，1 号房子粉刷成绿色，2 号房子粉刷成蓝色。
//      最少花费: 2 + 5 + 3 = 10。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/paint-house

package main

import "fmt"

func main()  {
	fmt.Println(minCost([][]int{{17,2,17}, {16,16,5}, {14,3,19}}))
}

func minCost(costs [][]int) int {
	n := len(costs)
	mp := make([][]int, n)
	mp[0] = costs[0]
	for k :=range mp[1:] {
		mp[k+1] = make([]int,3)
	}
	for i:=1; i<n; i++ {
		mp[i][0] = min(mp[i-1][1], mp[i-1][2]) + costs[i][0]
		mp[i][1] = min(mp[i-1][0], mp[i-1][2]) + costs[i][1]
		mp[i][2] = min(mp[i-1][1], mp[i-1][0]) + costs[i][2]
	}
	return min(min(mp[n-1][0], mp[n-1][1]), mp[n-1][2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
