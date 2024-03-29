// 给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
//
//示例 1 :
//
//输入: 2736
//输出: 7236
//解释: 交换数字2和数字7。
//示例 2 :
//
//输入: 9973
//输出: 9973
//解释: 不需要交换。
//注意:
//
//给定数字的范围是 [0, 108]
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/maximum-swap
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximumSwap(1993))
	fmt.Println(maximumSwap(98368))
	fmt.Println(maximumSwap(123219))
	fmt.Println(maximumSwap(0))
	fmt.Println(maximumSwap(19987))
}

func maximumSwap(num int) int {
	numStr := strconv.Itoa(num)
	nums := []byte(numStr)
	maxIdx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == '9' {
			continue
		}
		maxIdx = i
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] && nums[j] >= nums[maxIdx] || nums[i] == nums[j] && nums[j] > nums[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx != i {
			nums[maxIdx], nums[i] = nums[i], nums[maxIdx]
			break
		}
	}
	if maxIdx == 0 {
		return num
	}
	newNum, _ := strconv.Atoi(string(nums))
	return newNum
}
