// 给你一个二元数组 nums ，和一个整数 goal ，请你统计并返回有多少个和为 goal 的 非空 子数组。
//
// 子数组 是数组的一段连续部分。
//
//  
//
// 示例 1：
//
// 输入：nums = [1,0,1,0,1], goal = 2
// 输出：4
// 解释：
// 如下面黑体所示，有 4 个满足题目要求的子数组：
// [1,0,1,0,1]
// [1,0,1,0,1]
// [1,0,1,0,1]
// [1,0,1,0,1]
// 示例 2：
//
// 输入：nums = [0,0,0,0,0], goal = 0
// 输出：15
//  
//
// 提示：
//
// 1 <= nums.length <= 3 * 104
// nums[i] 不是 0 就是 1
// 0 <= goal <= nums.length
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-subarrays-with-sum

package main

import "fmt"

func main()  {
	fmt.Println(numSubarraysWithSum([]int{0,0,0,0,0}, 0))
	fmt.Println(numSubarraysWithSum([]int{1,0,1,0,1}, 2))
}

func numSubarraysWithSum(nums []int, goal int) int {
	ans := 0
	mp := map[int]int{}
	sum := 0
	for _,v :=range nums {
		mp[sum]++
		sum+=v
		ans += mp[sum-goal]
	}
	return ans
}
