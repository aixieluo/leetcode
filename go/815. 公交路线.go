// 给你一个数组 routes ，表示一系列公交线路，其中每个 routes[i] 表示一条公交线路，第 i 辆公交车将会在上面循环行驶。
//
// 例如，路线 routes[0] = [1, 5, 7] 表示第 0 辆公交车会一直按序列 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ... 这样的车站路线行驶。
// 现在从 source 车站出发（初始时不在公交车上），要前往 target 车站。 期间仅可乘坐公交车。
//
// 求出 最少乘坐的公交车数量 。如果不可能到达终点车站，返回 -1 。
//
//  
//
// 示例 1：
//
// 输入：routes = [[1,2,7],[3,6,7]], source = 1, target = 6
// 输出：2
// 解释：最优策略是先乘坐第一辆公交车到达车站 7 , 然后换乘第二辆公交车到车站 6 。
// 示例 2：
//
// 输入：routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
// 输出：-1
//  
//
// 提示：
//
// 1 <= routes.length <= 500.
// 1 <= routes[i].length <= 105
// routes[i] 中的所有值 互不相同
// sum(routes[i].length) <= 105
// 0 <= routes[i][j] < 106
// 0 <= source, target < 106
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/bus-routes

package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(numBusesToDestination([][]int{{1,2,7}}, 1, 6))
	fmt.Println(numBusesToDestination([][]int{{7,12}, {4,5,15}, {6}, {15,19}, {9,12,13}}, 15, 12))
	fmt.Println(numBusesToDestination([][]int{{1,2,7}, {3,6,7}}, 1, 6))
}

func numBusesToDestination(routes [][]int, source, target int) int {
	if source == target {
		return 0
	}

	n := len(routes)
	edge := make([][]bool, n)
	for i := range edge {
		edge[i] = make([]bool, n)
	}
	rec := map[int][]int{}
	for i, route := range routes {
		for _, site := range route {
			for _, j := range rec[site] {
				edge[i][j] = true
				edge[j][i] = true
			}
			rec[site] = append(rec[site], i)
		}
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	q := []int{}
	for _, site := range rec[source] {
		dis[site] = 1
		q = append(q, site)
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for y, b := range edge[x] {
			if b && dis[y] == -1 {
				dis[y] = dis[x] + 1
				q = append(q, y)
			}
		}
	}

	ans := math.MaxInt32
	for _, site := range rec[target] {
		if dis[site] != -1 && dis[site] < ans {
			ans = dis[site]
		}
	}
	if ans < math.MaxInt32 {
		return ans
	}
	return -1
}

// 过不去最后一轮过大的数据，导致超时，逻辑正确
func numBusesToDestination2(routes [][]int, source int, target int) int {
	rs := make([]map[int]bool, len(routes))
	for k,v := range routes {
		rs[k] = map[int]bool{}
		for _,site := range v {
			rs[k][site] = true
		}
	}
	ans := 1<<31 -1
	hash := map[int]bool{}
	hash[source] = true
	queue := [][2]int{{source, 0}}
	for len(queue) > 0 {
		_source, c := queue[0][0], queue[0][1]
		queue = queue[1:]
		if _source == target && c < ans {
			ans = c
		}
		for _,r :=range rs {
			if r[_source] {
				for val :=range r {
					if !hash[val] {
						queue = append(queue, [2]int{val, c+1})
						hash[val] = true
					}
				}
			}
		}
	}

	if ans == 1<<31 -1 {
		return -1
	} else {
		return ans
	}
}
