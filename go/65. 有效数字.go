// 有效数字（按顺序）可以分成以下几个部分：
//
// 一个 小数 或者 整数
// （可选）一个 'e' 或 'E' ，后面跟着一个 整数
// 小数（按顺序）可以分成以下几个部分：
//
// （可选）一个符号字符（'+' 或 '-'）
// 下述格式之一：
// 至少一位数字，后面跟着一个点 '.'
// 至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
// 一个点 '.' ，后面跟着至少一位数字
// 整数（按顺序）可以分成以下几个部分：
//
// （可选）一个符号字符（'+' 或 '-'）
// 至少一位数字
// 部分有效数字列举如下：
//
// ["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"]
// 部分无效数字列举如下：
//
// ["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]
// 给你一个字符串 s ，如果 s 是一个 有效数字 ，请返回 true 。
//
//  
//
// 示例 1：
//
// 输入：s = "0"
// 输出：true
// 示例 2：
//
// 输入：s = "e"
// 输出：false
// 示例 3：
//
// 输入：s = "."
// 输出：false
// 示例 4：
//
// 输入：s = ".1"
// 输出：true
//  
//
// 提示：
//
// 1 <= s.length <= 20
// s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，或者点 '.' 。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-number

package main

import (
	"strings"
)

func main() {
}

func isNumber(s string) bool {
	ec := strings.Count(s, "e")
	Ec := strings.Count(s, "E")
	var sa []string
	if ec+Ec > 1 {
		return false
	} else if ec == 1 {
		sa = strings.Split(s, "e")
	} else if Ec == 1 {
		sa = strings.Split(s, "E")
	} else {
		sa = []string{s}
	}
	if len(sa) == 1 {
		return isZs(s) || isXs(s)
	} else {
		return (isZs(sa[0]) || isXs(sa[0])) && isZs(sa[1])
	}
}

func isZs(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	if strings.Count(s, "+")+strings.Count(s, "-") > 0 {
		return false
	}
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return len(s) > 0
}

func isXs(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	if strings.Count(s, "+")+strings.Count(s, "-") > 0 {
		return false
	}
	if strings.Count(s, ".") != 1 {
		return false
	}
	for _, v := range s {
		if (v < '0' || v > '9') && v != '.' {
			return false
		}
	}
	return len(s) > 1
}
