// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//
//如果数组中不存在目标值 target，返回 [-1, -1]。
//
//你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
//
//
//
//示例 1：
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//示例 2：
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//示例 3：
//
//输入：nums = [], target = 0
//输出：[-1,-1]
//
//
//提示：
//
//0 <= nums.length <= 105
//-109 <= nums[i] <= 109
//nums 是一个非递减数组
//-109 <= target <= 109
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(searchRange([]int{1}, 1))
	fmt.Println(searchRange([]int{}, 6))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}

func searchRange(nums []int, target int) []int {
	ln := len(nums)
	if ln == 0 {
		return []int{-1, -1}
	}
	start := sort.SearchInts(nums, target)
	end := sort.SearchInts(nums, target+1) - 1
	if end < 0 {
		end = 0
	}
	if start >= ln || target != nums[start] {
		start = -1
	}
	if end >= ln || target != nums[end] {
		end = -1
	}
	return []int{start, end}
}
