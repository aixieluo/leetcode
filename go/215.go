package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(findKthLargest([]int{3,2,1,5,6,4}, 2))
}

type heapArr []int

func (h heapArr) Len() int { return len(h) }
func (h heapArr) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *heapArr) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *heapArr) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *heapArr) Pop() interface{} {
	n := h.Len()
	top := (*h)[n-1]
	*h = (*h)[:n-1]
	return top
}

func findKthLargest(nums []int, k int) int {
	ha := heapArr(nums)
	heap.Init(&ha)
	for {
		top := heap.Pop(&ha)
		if k == 1 {
			return top.(int)
		}
		k--
	}
}
