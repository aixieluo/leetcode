// 中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
//
// 例如，
//
// [2,3,4] 的中位数是 3
//
// [2,3] 的中位数是 (2 + 3) / 2 = 2.5
//
// 设计一个支持以下两种操作的数据结构：
//
// void addNum(int num) - 从数据流中添加一个整数到数据结构中。
// double findMedian() - 返回目前所有元素的中位数。
// 示例：
//
// addNum(1)
// addNum(2)
// findMedian() -> 1.5
// addNum(3)
// findMedian() -> 2
// 进阶:
//
// 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
// 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-median-from-data-stream

package main

import (
	"fmt"
)

func main() {
	m := Constructor()
	m.AddNum(1)
	fmt.Println(m.FindMedian())
	m.AddNum(2)
	fmt.Println(m.FindMedian())
	m.AddNum(3)
	fmt.Println(m.FindMedian())
	m.AddNum(4)
	fmt.Println(m.FindMedian())
	m.AddNum(5)
	fmt.Println(m.FindMedian())
	m.AddNum(6)
	fmt.Println(m.FindMedian())
	m.AddNum(7)
	fmt.Println(m.FindMedian())
	m.AddNum(8)
	fmt.Println(m.FindMedian())
	m.AddNum(9)
	fmt.Println(m.FindMedian())
	m.AddNum(10)
	fmt.Println(m.FindMedian())
	// fmt.Println(checkValidString("((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()"))
	// fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

type MedianFinder struct {
	List []int
	L    int
	M    int
	R    int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) Len() int {
	return len(this.List)
}

func (this *MedianFinder) AddNum(num int) {
	this.List = append(this.List, num)
	if this.Len() == 1 {
		this.L, this.M, this.R = num, num, num
		return
	}
	if this.Len() == 2 {
		this.L, this.R = min(num, this.M), max(num, this.M)
		return
	}
	if this.Len()%2 == 1 {
		this.M = min(max(num, this.L), this.R)
	} else {
		this.L = min(max(num, this.L), this.R)
		this.R = max(min(num, this.R), this.L)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	fmt.Println(this.L, this.M, this.R)
	list := this.List
	l := len(list)
	if l%2 == 0 {
		return (float64(this.L) + float64(this.R)) / 2
	} else {
		return float64(this.M)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
