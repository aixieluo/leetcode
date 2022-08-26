//给定一个整数 n ，返回 n! 结果中尾随零的数量。
//
//提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
//
//
//
//示例 1：
//
//输入：n = 3
//输出：0
//解释：3! = 6 ，不含尾随 0
//示例 2：
//
//输入：n = 5
//输出：1
//解释：5! = 120 ，有一个尾随 0
//示例 3：
//
//输入：n = 0
//输出：0
//
//
//提示：
//
//0 <= n <= 104
//
//
//进阶：你可以设计并实现对数时间复杂度的算法来解决此问题吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/factorial-trailing-zeroes
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
)

func main() {
	fmt.Println(trailingZeroes(30))
}

func trailingZeroes(n int) int {
	// 5 10 15 20 25 30
	sum := 0
	for i := 5; i <= n; i += 5 {
		sum += log5(i)
	}
	return sum
}

func log5(n int) int {
	num := 0
	for n > 0 {
		if m := n / 5 * 5; m == n {
			num++
			n /= 5
		} else {
			break
		}
	}
	return num
}
