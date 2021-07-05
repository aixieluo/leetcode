// 给定一个字符串，请将字符串里的字符按照出现的频率降序排列。
//
// 示例 1:
//
// 输入:
// "tree"
//
// 输出:
// "eert"
//
// 解释:
// 'e'出现两次，'r'和't'都只出现一次。
// 因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
// 示例 2:
//
// 输入:
// "cccaaa"
//
// 输出:
// "cccaaa"
//
// 解释:
// 'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
// 注意"cacaca"是不正确的，因为相同的字母必须放在一起。
// 示例 3:
//
// 输入:
// "Aabb"
//
// 输出:
// "bbAa"
//
// 解释:
// 此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
// 注意'A'和'a'被认为是两种不同的字符。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sort-characters-by-frequency

package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(frequencySort("2a554442f544asfasssffffasss"))
	fmt.Println(frequencySort("raaeaedere"))
	fmt.Println(frequencySort("ada"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("ccacaa"))
}


func frequencySort(s string) string {
	var hash map[byte]int
	bs := make([]byte, 0)
	hash = map[byte]int{}
	for _, v := range []byte(s) {
		if _,has := hash[v]; !has {
			bs = append(bs, v)
		}
		hash[v]++
	}
	sort.Slice(bs, func(i, j int) bool {
		return hash[bs[i]] >= hash[bs[j]]
	})
	ans := make([]byte, 0)
	for _,v :=range bs {
		ans = append(ans, bytes.Repeat([]byte{v}, hash[v])...)
	}
	return string(ans)
}
