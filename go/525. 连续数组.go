// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
//
//  
//
// 示例 1:
//
// 输入: nums = [0,1]
// 输出: 2
// 说明: [0, 1] 是具有相同数量0和1的最长连续子数组。
// 示例 2:
//
// 输入: nums = [0,1,0]
// 输出: 2
// 说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。
//  
//
// 提示：
//
// 1 <= nums.length <= 105
// nums[i] 不是 0 就是 1
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/contiguous-array

package main

import "fmt"

func main() {
	fmt.Println(findMaxLength([]int{0, 0, 1, 0, 0, 0, 1, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0}))
	fmt.Println(findMaxLength([]int{0, 1}))
}

func findMaxLength(nums []int) int {
	hash := map[int]int{0: -1}
	prev := 0
	c := 0
	for k, v := range nums {
		if v == 1 {
			prev++
		} else {
			prev--
		}
		if val, has:=hash[prev]; has {
			c = max(c, k - val)
		} else {
			hash[prev] = k
		}
	}
	return c
}

func findMaxLength2(nums []int) int {
	hash := make(map[int][]int)
	prev := 0
	for k, v := range nums {
		if v == 0 {
			prev--
		} else {
			prev++
		}
		if val, has := hash[prev]; has {
			hash[prev] = append(val, k)
		} else {
			hash[prev] = []int{k}
		}
	}
	count := 0
	for k, v := range hash {
		if k == 0 {
			count = max(count, v[len(v)-1]+1)
		} else {
			if len(v) > 1 {
				count = max(count, v[len(v)-1]-v[0])
			}
		}
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
