// 给出一个字符串 s（仅含有小写英文字母和括号）。
//
// 请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
//
// 注意，您的结果中 不应 包含任何括号。
//
//  
//
// 示例 1：
//
// 输入：s = "(abcd)"
// 输出："dcba"
// 示例 2：
//
// 输入：s = "(u(love)i)"
// 输出："iloveu"
// 示例 3：
//
// 输入：s = "(ed(et(oc))el)"
// 输出："leetcode"
// 示例 4：
//
// 输入：s = "a(bcdefghijkl(mno)p)q"
// 输出："apmnolkjihgfedcbq"
//  
//
// 提示：
//
// 0 <= s.length <= 2000
// s 中只有小写英文字母和括号
// 我们确保所有括号都是成对出现的
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses

package main

import "fmt"

func main() {
	fmt.Println(reverseParentheses("(ed(et(oc))el)"))
	fmt.Println(reverseParentheses("n(ev(t)((()lfevf))yd)cb()"))
	fmt.Println("ndyfvefltvecb")
	fmt.Println(reverseParentheses("ta()usw((((a))))"))
	fmt.Println(reverseParentheses("a(bcdefghijkl(mno)p)q"))
	fmt.Println(reverseParentheses("(u(love)i)"))
	fmt.Println(reverseParentheses("(abcd)"))
}

func reverseParentheses(s string) string {
	str, stack := make([]byte, 0), make([][]byte, 0)
	for k :=range s {
		if s[k] == '(' {
			stack = append(stack, str)
			str = []byte{}
		} else if s[k] == ')' {
			for i, n := 0, len(str); i< n /2; i++ {
				str[i], str[n-1-i] = str[n-1-i], str[i]
			}
			str = append(stack[len(stack)-1], str...)
			stack = stack[:len(stack)-1]
		} else {
			str = append(str, s[k])
		}
	}
	return string(str)
}

