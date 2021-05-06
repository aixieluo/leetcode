// 给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。
//
// 如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。
//
//  
//
// 进阶：很容易想到时间复杂度为 O(n^2) 的解决方案，你可以设计一个时间复杂度为 O(n logn) 或 O(n) 的解决方案吗？
//
//  
//
// 示例 1：
//
// 输入：nums = [1,2,3,4]
// 输出：false
// 解释：序列中不存在 132 模式的子序列。
// 示例 2：
//
// 输入：nums = [3,1,4,2]
// 输出：true
// 解释：序列中有 1 个 132 模式的子序列： [1, 4, 2] 。
// 示例 3：
//
// 输入：nums = [-1,3,2,0]
// 输出：true
// 解释：序列中有 3 个 132 模式的的子序列：[-1, 3, 2]、[-1, 3, 0] 和 [-1, 2, 0] 。
//  
//
// 提示：
//
// n == nums.length
// 1 <= n <= 104
// -109 <= nums[i] <= 109
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/132-pattern

package main

import "fmt"

func main() {
	fmt.Println(find132pattern([]int{1,2,3,4}))
}

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return  false
	}
	for i:=0; i< len(nums)-2 ;i++  {
		if nums[i] == nums[i+1] {
			continue
		}
		res := make([]int, 1, 3)
		res[0] = nums[i]
		for j:=i+1;j < len(nums) ; j++ {
			if nums[j] <= nums[i] {
				continue
			}
			if len(res) == 1 {
				res = append(res, nums[j])
			} else if len(res) == 2 {
				if res[1] <= nums[j] {
					res[1] = nums[j]
				} else {
					if res[0] < nums[j] {
						return true
					}
				}
			}
		}
	}
	return false
}

// 1 6
