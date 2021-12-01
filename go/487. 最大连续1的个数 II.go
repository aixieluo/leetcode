// 给定一个二进制数组，你可以最多将 1 个 0 翻转为 1，找出其中最大连续 1 的个数。
//
// 示例 1：
//
// 输入：[1,0,1,1,0]
// 输出：4
// 解释：翻转第一个 0 可以得到最长的连续 1。
//      当翻转以后，最大连续 1 的个数为 4。
//  
//
// 注：
//
// 输入数组只包含 0 和 1.
// 输入数组的长度为正整数，且不超过 10,000
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/max-consecutive-ones-ii

package main

import "fmt"

func main()  {
	fmt.Println(findMaxConsecutiveOnes([]int{1,0,1,1,0, 1, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	left := 0
	index := -1
	ans := 1
	for i:=0; i< len(nums); i++ {
		num := nums[i]
		if num == 0 {
			if index == -1 {
				index = i
			} else {
				left = index+1
				index = i
			}
		}
		if i - left + 1 > ans {
			ans = i - left + 1
		}
	}
	return ans
}
