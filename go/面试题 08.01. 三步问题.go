// 三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。
//
// 示例1:
//
// 输入：n = 3
// 输出：4
// 说明: 有四种走法
// 示例2:
//
// 输入：n = 5
// 输出：13
// 提示:
//
// n范围在[1, 1000000]之间
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/three-steps-problem-lcci

package main

import "fmt"

func main() {
	// fmt.Println(waysToStep(3))
	fmt.Println(waysToStep(61))
}

func waysToStep(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	}
	l1, l2, l3 := 1, 2, 4
	for index := 3; index < n; index++ {
		l3, l1, l2 = (l1+l2+l3)%1000000007, l2, l3
	}
	return l3
}
