// 在一个 XY 坐标系中有一些点，我们用数组 coordinates 来分别记录它们的坐标，其中 coordinates[i] = [x, y] 表示横坐标为 x、纵坐标为 y 的点。
//
// 请你来判断，这些点是否在该坐标系中属于同一条直线上，是则返回 true，否则请返回 false。
//
//  
//
// 示例 1：
//
//
//
// 输入：coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
// 输出：true
// 示例 2：
//
//
//
// 输入：coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
// 输出：false
//  
//
// 提示：
//
// 2 <= coordinates.length <= 1000
// coordinates[i].length == 2
// -10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
// coordinates 中不含重复的点
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/check-if-it-is-a-straight-line

package main

import "fmt"

func main()  {
	fmt.Println(checkStraightLine([][]int{{0,0}, {0,5}, {5,5}, {5,0}}))
}

func checkStraightLine(coordinates [][]int) bool {
	x1,y1,x2,y2 := coordinates[0][0], coordinates[0][1], coordinates[1][0], coordinates[1][1]
	if x1 -x2 == 0 {
		for _, v :=range coordinates {
			if v[0] != x1 {
				return false
			}
		}
		return true
	}
	k:= float64(y1-y2) / float64(x1-x2)
	b:= float64(y1)- k * float64(x1)
	for i:=2; i< len(coordinates); i++ {
		if float64(coordinates[i][1]) != k* float64(coordinates[i][0]) + b {
			return false
		}
	}
	return true
}
