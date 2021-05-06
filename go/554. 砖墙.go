// 你的面前有一堵矩形的、由 n 行砖块组成的砖墙。这些砖块高度相同（也就是一个单位高）但是宽度不同。每一行砖块的宽度之和应该相等。
//
// 你现在要画一条 自顶向下 的、穿过 最少 砖块的垂线。如果你画的线只是从砖块的边缘经过，就不算穿过这块砖。你不能沿着墙的两个垂直边缘之一画线，这样显然是没有穿过一块砖的。
//
// 给你一个二维数组 wall ，该数组包含这堵墙的相关信息。其中，wall[i] 是一个代表从左至右每块砖的宽度的数组。你需要找出怎样画才能使这条线 穿过的砖块数量最少 ，并且返回 穿过的砖块数量 。
//
//  
//
// 示例 1：
//
//
// 输入：wall = [[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]
// 输出：2
// 示例 2：
//
// 输入：wall = [[1],[1],[1]]
// 输出：3
//  
// 提示：
//
// n == wall.length
// 1 <= n <= 104
// 1 <= wall[i].length <= 104
// 1 <= sum(wall[i].length) <= 2 * 104
// 对于每一行 i ，sum(wall[i]) 应当是相同的
// 1 <= wall[i][j] <= 231 - 1
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/brick-wall

package main

import "fmt"

func main() {
	fmt.Println(leastBricks([][]int{{5,2,1,1,1}, {9,1}, {9,1}}))
	fmt.Println(leastBricks([][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}))
}

func leastBricks(wall [][]int) int {
	max := 0
	hash := make(map[int]int)
	for k := range wall {
		for k2 := range wall[k] {
			if k2 > 0 {
				wall[k][k2] += wall[k][k2-1]
			}
		}
		if wall[k][len(wall[k])-1] > max {
			max = wall[k][len(wall[k])-1]
		}
	}
	for k := range wall {
		for k2 :=range wall[k] {
			val := wall[k][k2]
			if val == max {
				break
			}
			if _, has := hash[val]; has {
				hash[val]++
			} else {
				hash[val] = 1
			}
		}
	}
	l := 0
	if len(hash) > 0 {
		for _, v:=range hash {
			if v > l {
				l = v
			}
		}
	}
	return len(wall)-l
}
