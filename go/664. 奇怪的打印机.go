// 有台奇怪的打印机有以下两个特殊要求：
//
// 打印机每次只能打印由 同一个字符 组成的序列。
// 每次可以在任意起始和结束位置打印新字符，并且会覆盖掉原来已有的字符。
// 给你一个字符串 s ，你的任务是计算这个打印机打印它需要的最少打印次数。
//
//  
// 示例 1：
//
// 输入：s = "aaabbb"
// 输出：2
// 解释：首先打印 "aaa" 然后打印 "bbb"。
// 示例 2：
//
// 输入：s = "aba"
// 输出：2
// 解释：首先打印 "aaa" 然后在第二个位置打印 "b" 覆盖掉原来的字符 'a'。
//  
//
// 提示：
//
// 1 <= s.length <= 100
// s 由小写英文字母组成
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/strange-printer

package main

import "fmt"

func main() {
	fmt.Println(strangePrinter("aaabbb"))
}

func strangePrinter(s string) int {
	dp := make([][]int, len(s))
	for k := range dp {
		dp[k] = make([]int, len(s))
		dp[k][k] = 1
	}
	for i := len(s)-2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			dp[i][j] = dp[i][j-1] + 1
			for k:=i; k< j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k] + dp[k+1][j])
			}
			if s[i] == s[j] {
				dp[i][j] = min(dp[i][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
