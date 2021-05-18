// 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
//
// 示例 1：
//
// 输入: 2
// 输出: 1
// 解释: 2 = 1 + 1, 1 × 1 = 1
// 示例 2:
//
// 输入: 10
// 输出: 36
// 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
// 提示：
//
// 2 <= n <= 58
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/jian-sheng-zi-lcof

package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(cuttingRope(10))
	fmt.Println(cuttingRope(120))
}

func cuttingRope(n int) int {
	var n2, n3 int
	if n== 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n % 2 == 0 {
		n3 = (n / 6) * 2
		n2 = (n - n3 * 3) / 2
	} else {
		n3 = n / 3
		if n- n3 *3 == 1 {
			n3--
		}
		n2 = (n- n3 *3) / 2
	}
	return int(math.Pow(3, float64(n3))* math.Pow(2, float64(n2)))
}
