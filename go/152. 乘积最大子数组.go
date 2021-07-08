// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
//  
//
// 示例 1:
//
// 输入: [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
// 示例 2:
//
// 输入: [-2,0,-1]
// 输出: 0
// 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-product-subarray

package main

import "fmt"

func main()  {
	fmt.Println(maxProduct([]int{-2}))
	fmt.Println(maxProduct([]int{-2,-1,0,-3,-1}))
	fmt.Println(maxProduct([]int{-2,0,-1}))
	fmt.Println(maxProduct([]int{2,3,-2,4}))
}

func maxProduct(nums []int) int {
	zeros := make([]int, 0)
	for k, v :=range nums{
		if v == 0 {
			zeros = append(zeros, k)
		}
	}
	if len(zeros) == 0 {
		return fn(nums)
	} else {
		res := make([]int, 0)
		zeros = append(append([]int{-1}, zeros...), len(nums))
		for i:=0; i< len(zeros) -1 ; i++ {
			res = append(res , fn(nums[zeros[i]+1: zeros[i+1]]))
		}
		rlt := max(res...)
		if rlt < 0 {
			return 0
		} else {
			return rlt
		}
	}
}

func fn(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n := 0
	minIndex, maxIndex := len(nums), -1
	for k,v :=range nums {
		if v < 0 {
			if k < minIndex {
				minIndex = k
			}
			if k> maxIndex {
				maxIndex = k
			}
			n++
		}
	}
	if n % 2 ==0 {
		return j(nums)
	} else {
		return max(j(nums[minIndex+1:]), j(nums[:maxIndex]))
	}
}

func j(nums []int) int {
	if len(nums) ==0 {
		return 0
	}
	a :=1
	for _,v :=range nums {
		a*=v
	}
	return a
}

func max(a ...int) int {
	b := -1 <<31
	for _,v :=range a {
		if v > b {
			b=v
		}
	}
	return b
}
