// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的数字可以无限制重复被选取。
//
// 说明：
//
// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。 
// 示例 1：
//
// 输入：candidates = [2,3,6,7], target = 7,
// 所求解集为：
// [
//  [7],
//  [2,2,3]
// ]
// 示例 2：
//
// 输入：candidates = [2,3,5], target = 8,
// 所求解集为：
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]
//  
//
// 提示：
//
// 1 <= candidates.length <= 30
// 1 <= candidates[i] <= 200
// candidate 中的每个元素都是独一无二的。
// 1 <= target <= 500
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/combination-sum

package main

import "fmt"

func main()  {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	carry := make([]int, 0)
	var dfs func(t int, start int)
	dfs = func(t int, start int) {
		if t== 0 && len(carry) > 0 {
			res = append(res, append([]int{}, carry...))
		} else {
			for k,v :=range candidates {
				if k < start {
					continue
				}
				if t -v >= 0 {
					carry = append(carry, v)
					start = k
					dfs(t-v, start)
					carry = carry[:len(carry)-1]
					start -= k
				}
			}
		}
	}
	dfs(target, 0)
	return res
}
