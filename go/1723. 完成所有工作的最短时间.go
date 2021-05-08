// 给你一个整数数组 jobs ，其中 jobs[i] 是完成第 i 项工作要花费的时间。
//
// 请你将这些工作分配给 k 位工人。所有工作都应该分配给工人，且每项工作只能分配给一位工人。工人的 工作时间 是完成分配给他们的所有工作花费时间的总和。请你设计一套最佳的工作分配方案，使工人的 最大工作时间 得以 最小化 。
//
// 返回分配方案中尽可能 最小 的 最大工作时间 。
//
//  
//
// 示例 1：
//
// 输入：jobs = [3,2,3], k = 3
// 输出：3
// 解释：给每位工人分配一项工作，最大工作时间是 3 。
// 示例 2：
//
// 输入：jobs = [1,2,4,7,8], k = 2
// 输出：11
// 解释：按下述方式分配工作：
// 1 号工人：1、2、8（工作时间 = 1 + 2 + 8 = 11）
// 2 号工人：4、7（工作时间 = 4 + 7 = 11）
// 最大工作时间是 11 。
//  
//
// 提示：
//
// 1 <= k <= jobs.length <= 12
// 1 <= jobs[i] <= 107
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-minimum-time-to-finish-all-jobs

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// fmt.Println(minimumTimeRequired([]int{46, 13, 54, 51, 38, 49, 54, 67, 26, 78, 33}, 10))
	fmt.Println(minimumTimeRequired([]int{1, 2, 4, 7, 8}, 2))
	fmt.Println(minimumTimeRequired([]int{3, 2, 3}, 3))
}

func minimumTimeRequired(jobs []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(jobs)))
	l,r:= jobs[0], 0
	for _,v:=range jobs {
		r+=v
	}
	return l+ sort.Search(r-l, func(limit int) bool {
		limit+= l
		workers := make([]int, k)
		var f func(start int) bool
		f= func(start int) bool {
			for i:=0; i< k ;i++  {
				if start == len(jobs) {
					return true
				}
				if workers[i] + jobs[start] <= limit {
					workers[i] += jobs[start]
					if f(start+1) {
						return true
					}
					workers[i] -= jobs[start]
				}
				if workers[i] == 0 || workers[i] + jobs[i] == limit {
					break
				}
			}
			return false
		}
		return f(0)
	})
}

func minimumTimeRequired2(jobs []int, k int) int {
	dfs := make([]int, k)
	res := math.MaxInt64
	var f func(start int, used int, max int)
	f = func(start int, used int, max int) {
		if max >= res {
			return
		}
		if start == len(jobs) {
			res = max
			return
		}
		if used < k {
			dfs[used] += jobs[start]
			f(start+1, used+1, maxInt(dfs[used], max))
			dfs[used] -= jobs[start]
		}
		for k2 :=0; k2 < used; k2++ {
			dfs[k2] += jobs[start]
			f(start+1, used, maxInt(dfs[k2], max))
			dfs[k2] -= jobs[start]
		}
	}
	f(0, 0, 0)
	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
