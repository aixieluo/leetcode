// 给你一个二进制字符串 s（仅由 '0' 和 '1' 组成的字符串）。
//
// 返回所有字符都为 1 的子字符串的数目。
//
// 由于答案可能很大，请你将它对 10^9 + 7 取模后返回。
//
//  
//
// 示例 1：
//
// 输入：s = "0110111"
// 输出：9
// 解释：共有 9 个子字符串仅由 '1' 组成
// "1" -> 5 次
// "11" -> 3 次
// "111" -> 1 次
// 示例 2：
//
// 输入：s = "101"
// 输出：2
// 解释：子字符串 "1" 在 s 中共出现 2 次
// 示例 3：
//
// 输入：s = "111111"
// 输出：21
// 解释：每个子字符串都仅由 '1' 组成
// 示例 4：
//
// 输入：s = "000"
// 输出：0
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/number-of-substrings-with-only-1s

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSub("0110111"))
}

func numSub(s string) int {
	sum := 0
	count := 0
	for _, v := range s {
		if v == '1' {
			count++
		} else {
			sum += count * (count + 1) / 2
			count = 0
		}
	}
	if count > 0 {
		sum += count * (count + 1) / 2
	}
	return sum % 1000000007
}

func numSub2(s string) int {
	l := make([]int, 0, 20)
	max := 0
	slow := 0
	for fast := 0; fast < len(s); fast++ {
		n := 0
		if s[fast] == '0' {
			n = fast - slow
		} else if fast == len(s)-1 {
			n = fast - slow + 1
		} else {
			continue
		}
		slow = fast + 1
		if n > 0 {
			l = append(l, n)
			if n > max {
				max = n
			}
		}
	}
	hash := make(map[int]int)
	hash[1] = 1
	for i := 2; i <= max; i++ {
		hash[i] = i + hash[i-1]
	}
	sum := 0
	for _, v := range l {
		sum += hash[v]
	}
	return sum % int(math.Pow(10, 9)+7)
}
