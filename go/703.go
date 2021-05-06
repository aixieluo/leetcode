package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	kth := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kth.Add(3))
	fmt.Println(kth.Add(5))
	fmt.Println(kth.Add(10))
	fmt.Println(kth.Add(9))
	fmt.Println(kth.Add(4))
}

type KthLargest struct {
	sort.IntSlice
	K    int
}

func (kth *KthLargest) Push(x interface{}) {
	kth.IntSlice = append(kth.IntSlice, x.(int))
}

func (kth *KthLargest) Pop() interface{} {
	a := kth.IntSlice
	kth.IntSlice = a[:a.Len()-1]
	return a[a.Len()-1]
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{K: k}
	for _, v := range nums {
		kth.Add(v)
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.K {
		heap.Pop(this)
	}
	return this.IntSlice[0]
}
