// 给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。
//
//  
//
// 示例 1：
//
//
// 输入：points = [[1,1],[2,2],[3,3]]
// 输出：3
// 示例 2：
//
//
// 输入：points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
// 输出：4
//  
//
// 提示：
//
// 1 <= points.length <= 300
// points[i].length == 2
// -104 <= xi, yi <= 104
// points 中的所有点 互不相同
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/max-points-on-a-line

package main

import (
	"fmt"
)

// kx+b = y => k = (y-b)/x
// 5 = 4 *k + b
// -1 = 4 * k + b

func main() {
	fmt.Println(maxPoints([][]int{{-6, -1}, {3, 1}, {12, 3}}), 3)
	fmt.Println(maxPoints([][]int{{4, 5}, {4, -1}, {4, 0}}), 3)
	fmt.Println(maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}), 4)
}

func maxPoints(points [][]int) (ans int) {
	n := len(points)
	if n <= 2 {
		return n
	}
	for i, p := range points {
		if ans >= n-i || ans > n/2 {
			break
		}
		cnt := map[int]int{}
		for _, q := range points[i+1:] {
			x, y := p[0]-q[0], p[1]-q[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					x, y = -x, -y
				}
				g := gcd(abs(x), abs(y))
				x /= g
				y /= g
			}
			cnt[y+x*10001]++
		}
		for _, c := range cnt {
			ans = max(ans, c+1)
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
