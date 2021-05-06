// 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
//
//  
//
// 示例 1：
//
// 输入：nums = [10,2]
// 输出："210"
// 示例 2：
//
// 输入：nums = [3,30,34,5,9]
// 输出："9534330"
// 示例 3：
//
// 输入：nums = [1]
// 输出："1"
// 示例 4：
//
// 输入：nums = [10]
// 输出："10"
//  
//
// 提示：
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 109
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/largest-number

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main()  {
	fmt.Println(largestNumber([]int{3,30,34,5,9}))
}

type largeNums struct {
	sort.IntSlice
}

func (ln largeNums) Less(i, j int) bool {
	l, r:=strconv.Itoa(ln.IntSlice[i]), strconv.Itoa(ln.IntSlice[j])
	if l+r > r+l {
		return true
	}
	return false
}

func largestNumber(nums []int) string {
	n := &largeNums{nums}
	sort.Sort(n)
	var s string
	for _,v:=range n.IntSlice {
		s += strconv.Itoa(v)
	}
	if s[0] == '0' {
		return "0"
	}
	return s
}
