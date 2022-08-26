// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
//
//
// 示例 1:
//
// 输入: nums = [-1,0,3,5,9,12], target = 9
// 输出: 4
// 解释: 9 出现在 nums 中并且下标为 4
// 示例 2:
//
// 输入: nums = [-1,0,3,5,9,12], target = 2
// 输出: -1
// 解释: 2 不存在 nums 中因此返回 -1
//  
//
// 提示：
//
// 你可以假设 nums 中的所有元素是不重复的。
// n 将在 [1, 10000]之间。
// nums 的每个元素都将在 [-9999, 9999]之间。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-search

package main

import "sort"

func main() {

}

func search(nums []int, target int) int {
	l := len(nums)
	ans := -1
	for i, j := 0, l-1; i <= j && j < l; {
		mid := (i + j) >> 1
		if nums[mid] > target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			ans = mid
			break
		}
	}
	return ans
}

func search2(nums []int, target int) int {
	x := sort.SearchInts(nums, target)
	if x < len(nums) && nums[x] == target {
		return x
	}
	return -1
}
