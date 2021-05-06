// 要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：
//
// "AAJF" ，将消息分组为 (1 1 10 6)
// "KJF" ，将消息分组为 (11 10 6)
// 注意，消息不能分组为  (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。
//
// 给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。
//
// 题目数据保证答案肯定是一个 32 位 的整数。
//
//  
//
// 示例 1：
//
// 输入：s = "12"
// 输出：2
// 解释：它可以解码为 "AB"（1 2）或者 "L"（12）。
// 示例 2：
//
// 输入：s = "226"
// 输出：3
// 解释：它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
// 示例 3：
//
// 输入：s = "0"
// 输出：0
// 解释：没有字符映射到以 0 开头的数字。
// 含有 0 的有效映射是 'J' -> "10" 和 'T'-> "20" 。
// 由于没有字符，因此没有有效的方法对此进行解码，因为所有数字都需要映射。
// 示例 4：
//
// 输入：s = "06"
// 输出：0
// 解释："06" 不能映射到 "F" ，因为字符串含有前导 0（"6" 和 "06" 在映射中并不等价）。
//  
//
// 提示：
//
// 1 <= s.length <= 100
// s 只包含数字，并且可能包含前导零。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/decode-ways

package main

import (
	"fmt"
)

func main() {
	fmt.Println("06", 0,numDecodings("06"))
	fmt.Println("111111111111111111111111111111111111111111111", 1836311903,numDecodings("111111111111111111111111111111111111111111111"))
	fmt.Println("2101", 1,numDecodings("2101"))
	fmt.Println("27", 1, numDecodings("27"))
	fmt.Println("207", 1,numDecodings("207"))
	fmt.Println("10011", 0,numDecodings("10011"))
	fmt.Println("226", 3,numDecodings("226"))
	fmt.Println("12", 2,numDecodings("12"))
	fmt.Println("1123", 5,numDecodings("1123"))
	fmt.Println("10", 1,numDecodings("10"))
}

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	index:=0
	for ; index< len(s) ;index++  {
		if s[index] == '0' {
			dp[index+1] = 0
		} else {
			dp[index+1] = dp[index]
		}
		if index>0 && (s[index-1] == '1' || (s[index-1] =='2' && s[index] <= '6')) {
			dp[index+1] += dp[index-1]
		}
	}
	return dp[index]
}
