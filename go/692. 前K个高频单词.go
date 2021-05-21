// 给一非空的单词列表，返回前 k 个出现次数最多的单词。
//
// 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。
//
// 示例 1：
//
// 输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
// 输出: ["i", "love"]
// 解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
//    注意，按字母顺序 "i" 在 "love" 之前。
//  
//
// 示例 2：
//
// 输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
// 输出: ["the", "is", "sunny", "day"]
// 解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
//    出现次数依次为 4, 3, 2 和 1 次。
//  
//
// 注意：
//
// 假定 k 总为有效值， 1 ≤ k ≤ 集合元素数。
// 输入的单词均由小写字母组成。
//  
//
// 扩展练习：
//
// 尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/top-k-frequent-words

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent2([]string{"a", "aa", "aaa"}, 2))
	fmt.Println(topKFrequent([]string{"glarko", "zlfiwwb", "nsfspyox", "pwqvwmlgri", "qggx", "qrkgmliewc", "zskaqzwo", "zskaqzwo", "ijy", "htpvnmozay", "jqrlad", "ccjel", "qrkgmliewc", "qkjzgws", "fqizrrnmif", "jqrlad", "nbuorw", "qrkgmliewc", "htpvnmozay", "nftk", "glarko", "hdemkfr", "axyak", "hdemkfr", "nsfspyox", "nsfspyox", "qrkgmliewc", "nftk", "nftk", "ccjel", "qrkgmliewc", "ocgjsu", "ijy", "glarko", "nbuorw", "nsfspyox", "qkjzgws", "qkjzgws", "fqizrrnmif", "pwqvwmlgri", "nftk", "qrkgmliewc", "jqrlad", "nftk", "zskaqzwo", "glarko", "nsfspyox", "zlfiwwb", "hwlvqgkdbo", "htpvnmozay", "nsfspyox", "zskaqzwo", "htpvnmozay", "zskaqzwo", "nbuorw", "qkjzgws", "zlfiwwb", "pwqvwmlgri", "zskaqzwo", "qengse", "glarko", "qkjzgws", "pwqvwmlgri", "fqizrrnmif", "nbuorw", "nftk", "ijy", "hdemkfr", "nftk", "qkjzgws", "jqrlad", "nftk", "ccjel", "qggx", "ijy", "qengse", "nftk", "htpvnmozay", "qengse", "eonrg", "qengse", "fqizrrnmif", "hwlvqgkdbo", "qengse", "qengse", "qggx", "qkjzgws", "qggx", "pwqvwmlgri", "htpvnmozay", "qrkgmliewc", "qengse", "fqizrrnmif", "qkjzgws", "qengse", "nftk", "htpvnmozay", "qggx", "zlfiwwb", "bwp", "ocgjsu", "qrkgmliewc", "ccjel", "hdemkfr", "nsfspyox", "hdemkfr", "qggx", "zlfiwwb", "nsfspyox", "ijy", "qkjzgws", "fqizrrnmif", "qkjzgws", "qrkgmliewc", "glarko", "hdemkfr", "pwqvwmlgri"}, 14))
	fmt.Println(topKFrequent([]string{"plpaboutit", "jnoqzdute", "sfvkdqf", "mjc", "nkpllqzjzp", "foqqenbey", "ssnanizsav", "nkpllqzjzp", "sfvkdqf", "isnjmy", "pnqsz", "hhqpvvt", "fvvdtpnzx", "jkqonvenhx", "cyxwlef", "hhqpvvt", "fvvdtpnzx", "plpaboutit", "sfvkdqf", "mjc", "fvvdtpnzx", "bwumsj", "foqqenbey", "isnjmy", "nkpllqzjzp", "hhqpvvt", "foqqenbey", "fvvdtpnzx", "bwumsj", "hhqpvvt", "fvvdtpnzx", "jkqonvenhx", "jnoqzdute", "foqqenbey", "jnoqzdute", "foqqenbey", "hhqpvvt", "ssnanizsav", "mjc", "foqqenbey", "bwumsj", "ssnanizsav", "fvvdtpnzx", "nkpllqzjzp", "jkqonvenhx", "hhqpvvt", "mjc", "isnjmy", "bwumsj", "pnqsz", "hhqpvvt", "nkpllqzjzp", "jnoqzdute", "pnqsz", "nkpllqzjzp", "jnoqzdute", "foqqenbey", "nkpllqzjzp", "hhqpvvt", "fvvdtpnzx", "plpaboutit", "jnoqzdute", "sfvkdqf", "fvvdtpnzx", "jkqonvenhx", "jnoqzdute", "nkpllqzjzp", "jnoqzdute", "fvvdtpnzx", "jkqonvenhx", "hhqpvvt", "isnjmy", "jkqonvenhx", "ssnanizsav", "jnoqzdute", "jkqonvenhx", "fvvdtpnzx", "hhqpvvt", "bwumsj", "nkpllqzjzp", "bwumsj", "jkqonvenhx", "jnoqzdute", "pnqsz", "foqqenbey", "sfvkdqf", "sfvkdqf"}, 1))
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}

type pair struct {
	s string
	c int
}

type topK struct {
	pairs []pair
}

func (t topK) Len() int {
	return len(t.pairs)
}

func (t topK) Less(i, j int) bool {
	a, b := t.pairs[i], t.pairs[j]
	return a.c < b.c || a.c == b.c && a.s > b.s
}

func (t topK) Swap(i, j int) {
	t.pairs[i], t.pairs[j] = t.pairs[j], t.pairs[i]
}

func (t *topK) Push(x interface{}) {
	t.pairs = append(t.pairs, x.(pair))
}

func (t *topK) Pop() interface{} {
	x := t.pairs[t.Len()-1]
	t.pairs = t.pairs[:t.Len()-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	hash := make(map[string]int)
	for _, v := range words {
		hash[v]++
	}
	hp := &topK{}
	heap.Init(hp)
	for key, val := range hash {
		heap.Push(hp, pair{key, val})
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}
	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(hp).(pair).s
	}
	return ans
}

func topKFrequent2(words []string, k int) []string {
	hash := make(map[string]int)
	for _, v := range words {
		hash[v]++
	}
	ws := make([]pair, 0, len(hash))
	for key, val := range hash {
		ws = append(ws, pair{key, val})
	}
	qSort(ws, 0, len(ws)-1)
	ans := make([]string, k)
	for k := range ans {
		ans[k] = ws[k].s
	}
	return ans
}

func qSort(ws []pair, start, end int) {
	if start+1 > end {
		return
	}
	p := partition(ws, start, end)
	qSort(ws, start, p-1)
	qSort(ws, p+1, end)
}

func partition(ws []pair, start, end int) int {
	pivot := ws[end]
	p := start
	for i := start; i < end; i++ {
		if ws[i].c > pivot.c || ws[i].c == pivot.c && ws[i].s < pivot.s {
			ws[i], ws[p] = ws[p], ws[i]
			p++
		}
	}
	ws[end], ws[p] = ws[p], ws[end]
	return p
}
