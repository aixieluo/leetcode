// 给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。
//
//  
//
// 示例 1：
//
// 输入：c = 5
// 输出：true
// 解释：1 * 1 + 2 * 2 = 5
// 示例 2：
//
// 输入：c = 3
// 输出：false
// 示例 3：
//
// 输入：c = 4
// 输出：true
// 示例 4：
//
// 输入：c = 2
// 输出：true
// 示例 5：
//
// 输入：c = 1
// 输出：true
//  
//
// 提示：
//
// 0 <= c <= 231 - 1
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sum-of-square-numbers

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum(0))
}

func judgeSquareSum(c int) bool {
	for a:=0; a *a <= c; a++ {
		b := math.Sqrt(float64(c-a*a))
		if b == math.Floor(b) {
			return true
		}
	}
	return false
}
