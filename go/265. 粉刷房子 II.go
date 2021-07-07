// 假如有一排房子，共 n 个，每个房子可以被粉刷成 k 种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。
//
// 当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n x k 的矩阵来表示的。
//
// 例如，costs[0][0] 表示第 0 号房子粉刷成 0 号颜色的成本花费；costs[1][2] 表示第 1 号房子粉刷成 2 号颜色的成本花费，以此类推。请你计算出粉刷完所有房子最少的花费成本。
//
// 注意：
//
// 所有花费均为正整数。
//
// 示例：
//
// 输入: [[1,5,3],[2,9,4]]
// 输出: 5
// 解释: 将 0 号房子粉刷成 0 号颜色，1 号房子粉刷成 2 号颜色。最少花费: 1 + 4 = 5;
//      或者将 0 号房子粉刷成 2 号颜色，1 号房子粉刷成 0 号颜色。最少花费: 3 + 2 = 5.
// 进阶：
// 您能否在 O(nk) 的时间复杂度下解决此问题？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/paint-house-ii

package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCostII([][]int{{1,3}, {2,4}}))
	fmt.Println(minCostII([][]int{{17,2,17}, {16,16,5}, {14,3,19}}))
	fmt.Println(minCostII([][]int{{1, 5, 3}, {2, 9, 4}}))
}

func minCostII(costs [][]int) int {
	n := len(costs)
	k := len(costs[0])
	mp := make([][]int, n)
	mp[0] = costs[0]
	for key := range mp[1:] {
		mp[key+1] = make([]int, k)
	}
	for i := 1; i < n; i++ {
		for j:=0; j<k; j++ {
			minVal := 1<< 31 -1
			for j2 := 0; j2 <k; j2++ {
				if j2 == j {
					continue
				}
				minVal = min(mp[i-1][j2], minVal)
			}
			mp[i][j] = minVal + costs[i][j]
		}
	}
	minVal := 1<<31-1
	for _,v :=range mp[n-1] {
		minVal = min(minVal, v)
	}
	return minVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
