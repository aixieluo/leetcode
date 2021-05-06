// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 注意：给定 n 是一个正整数。
//
// 示例 1：
//
// 输入： 2
// 输出： 2
// 解释： 有两种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶
// 2.  2 阶
// 示例 2：
//
// 输入： 3
// 输出： 3
// 解释： 有三种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶 + 1 阶
// 2.  1 阶 + 2 阶
// 3.  2 阶 + 1 阶
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/climbing-stairs

package main

import "fmt"

func main() {
	fmt.Println(climbStairs(44))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp1, dp2 := 1,2
	for i:=2;i <n; i++ {
		dp1, dp2 = dp2, dp1+dp2
	}
	return dp2
}

func climbStairs3(n int) int {
	l, r, res := 0, 0, 1
	for ; n > 0; n-- {
		l = r
		r = res
		res = l + r
	}
	return res
}

// 正确答案但是超时
func climbStairs2(n int) int {
	var f func(num int) int
	f = func(num int) int {
		if num == 0 {
			return 1
		}
		if num > 0 {
			return f(num-1) + f(num-2)
		}
		return 0
	}
	return f(n)
}
