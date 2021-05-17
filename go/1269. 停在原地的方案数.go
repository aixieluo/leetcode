// 有一个长度为 arrLen 的数组，开始有一个指针在索引 0 处。
//
// 每一步操作中，你可以将指针向左或向右移动 1 步，或者停在原地（指针不能被移动到数组范围外）。
//
// 给你两个整数 steps 和 arrLen ，请你计算并返回：在恰好执行 steps 次操作以后，指针仍然指向索引 0 处的方案数。
//
// 由于答案可能会很大，请返回方案数 模 10^9 + 7 后的结果。
//
//  
//
// 示例 1：
//
// 输入：steps = 3, arrLen = 2
// 输出：4
// 解释：3 步后，总共有 4 种不同的方法可以停在索引 0 处。
// 向右，向左，不动
// 不动，向右，向左
// 向右，不动，向左
// 不动，不动，不动
// 示例  2：
//
// 输入：steps = 2, arrLen = 4
// 输出：2
// 解释：2 步后，总共有 2 种不同的方法可以停在索引 0 处。
// 向右，向左
// 不动，不动
// 示例 3：
//
// 输入：steps = 4, arrLen = 2
// 输出：8
//  
//
// 提示：
//
// 1 <= steps <= 500
// 1 <= arrLen <= 10^6
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps

package main

import "fmt"

func main() {
	fmt.Println(numWays(47, 38), "318671228 <- 47, 38")
	fmt.Println(numWays(4, 2), "8 <- 4, 2")
	fmt.Println(numWays(3, 2), "4 <- 3, 2")
	fmt.Println(numWays(2, 4), "2 <- 2, 4")
}

func numWays(steps int, arrLen int) int {
	const m = 1000000007
	hash := make(map[int]int)
	var dfs func(s int, curr int) int
	dfs = func(s int, curr int) int {
		key := getKey(s, curr)
		if v, has := hash[key]; has {
			return v
		}
		if curr == 0 && s == 0 {
			return 1
		}
		if curr > s {
			return 0
		}
		sum := 0
		if curr != 0 {
			sum += dfs(s-1, curr-1)
		}
		if curr != arrLen-1 {
			sum += dfs(s-1, curr+1)
		}
		sum += dfs(s-1, curr)
		if _,has:=hash[key]; !has {
			hash[key] = sum % m
		}
		return sum
	}
	dfs(steps, 0)
	return hash[getKey(steps, 0)]
}

func getKey(step int, curr int) int {
	return step*1000 + curr
}
