package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

type KthMany struct {
	Arr []int
	List map[int]int
}

func (km KthMany) Len() int {
	return len(km.Arr)
}

func (km KthMany) Less(i, j int) bool {
	return km.List[km.Arr[i]] < km.List[km.Arr[j]]
}

func (km *KthMany)Swap(i, j int)  {
	km.Arr[i], km.Arr[j] = km.Arr[j], km.Arr[i]
}

func (km *KthMany) Push(v interface{}) {
	if _, ok := km.List[v.(int)]; ok {
		km.List[v.(int)]++
	} else {
		km.List[v.(int)] = 1
		km.Arr = append(km.Arr, v.(int))
	}
}

func (km *KthMany) Pop() interface{} {
	a:= km.Arr
	km.Arr = a[:km.Len()-1]
	return a[km.Len()-1]
}

func New(nums []int) KthMany {
	km := KthMany{}
	km.List = make(map[int]int)
	for _, v := range nums {
		km.Push(v)
	}
	heap.Init(&km)
	return km
}

func topKFrequent(nums []int, k int) []int {
	km := New(nums)
	for km.Len() > k {
		heap.Pop(&km)
	}
	return km.Arr
}
