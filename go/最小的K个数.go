package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(getLeastNumbers([]int{4, 5, 1, 6, 2, 7, 3}, 2))
}

type heapArr []int

func (h heapArr) Len() int { return len(h) }
func (h heapArr) Less(i, j int) bool {
	return h[i] < h[j]
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

func getLeastNumbers(arr []int, k int) []int {
	ha := heapArr(arr)
	heap.Init(&ha)
	result := make([]int, 0)
	for ; k > 0; k-- {
		result = append(result, heap.Pop(&ha).(int))
	}
	return result
}
