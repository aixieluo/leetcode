// 在一个二维平面空间中，给你 n 个点的坐标。问，是否能找出一条平行于 y 轴的直线，让这些点关于这条直线成镜像排布？
//
// 注意：题目数据中可能有重复的点。
//
//  
//
// 进阶：你能找到比 O(n2) 更优的解法吗?
//
//  
//
// 示例 1：
//
//
// 输入：points = [[1,1],[-1,1]]
// 输出：true
// 解释：可以找出 x = 0 这条线。
// 示例 2：
//
//
// 输入：points = [[1,1],[-1,-1]]
// 输出：false
// 解释：无法找出这样一条线。
//  
//
// 提示：
//
// n == points.length
// 1 <= n <= 10^4
// -10^8 <= points[i][j] <= 10^8
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/line-reflection

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isReflected([][]int{{1, 2}, {2,2}, {1,4}, {2,4}}))
}

func isReflected(points [][]int) bool {
	max, min := -1 << 63, 1 << 63 -1
	hash := make(map[int]map[int]bool)
	for _, p := range points {
		if max < p[0] {
			max = p[0]
		}
		if min > p[0] {
			min = p[0]
		}
		if _, ok := hash[p[0]]; !ok {
			hash[p[0]] = map[int]bool{p[1]:true}
		} else {
			hash[p[0]][p[1]] = true
		}
	}
	for _, p := range points {
		x := max + min - p[0]
		// if _, ok := hash[x]; !ok {
		// 	return false
		// }
		if _, ok := hash[x][p[1]]; !ok {
			return false
		}
	}
	return true
}
