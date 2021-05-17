// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用一次。
//
// 说明：
//
// 所有数字（包括目标数）都是正整数。
// 解集不能包含重复的组合。 
// 示例 1:
//
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 所求解集为:
// [
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
// ]
// 示例 2:
//
// 输入: candidates = [2,5,2,1,2], target = 5,
// 所求解集为:
// [
//   [1,2,2],
//   [5]
// ]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/combination-sum-ii

package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	carry := make([]int, 0)
	var dfs func(t int, start int)
	dfs = func(t int, start int) {
		if t== 0 && len(carry) > 0 {
			if len(res) > 0 {
				if equalSlice(res[len(res)-1], carry) {
					return
				}
			}
			res = append(res, append([]int{}, carry...))
		} else {
			for k:=start; k < len(candidates); k++ {
				v:= candidates[k]
				if t -v >= 0 {
					carry = append(carry, v)
					dfs(t-v, k+1)
					carry = carry[:len(carry)-1]
				}
			}
		}
	}
	dfs(target, 0)
	return res
}

func equalSlice(s1, s2 []int) bool  {
	if len(s1) != len(s2) {
		return false
	}
	if (s1 == nil) != (s2 == nil) {
		return false
	}
	for k,v:=range s1 {
		if v != s2[k] {
			return false
		}
	}
	return true
}
