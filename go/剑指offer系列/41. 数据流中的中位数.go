package main

import (
	"container/heap"
	"sort"
)

func main()  {

}

type MaxHeap struct {
	sort.IntSlice
}

func (m *MaxHeap) Less(i,j int) bool {
	return m.IntSlice.Less(j, i)
}

func (m *MaxHeap) Push(x interface{}) {
	m.IntSlice = append(m.IntSlice, x.(int))
}

func (m *MaxHeap) Pop() interface{} {
	x := m.IntSlice[m.Len()-1]
	m.IntSlice = m.IntSlice[:m.Len()-1]
	return x
}

type MinHeap struct {
	sort.IntSlice
}

func (m *MinHeap) Push(x interface{}) {
	m.IntSlice = append(m.IntSlice, x.(int))
}

func (m *MinHeap) Pop() interface{} {
	x := m.IntSlice[m.Len()-1]
	m.IntSlice = m.IntSlice[:m.Len()-1]
	return x
}

type MedianFinder struct {
	*MaxHeap
	*MinHeap
	Mid int
	Count int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	m:= MedianFinder{&MaxHeap{}, &MinHeap{}, 0, 0}
	heap.Init(m.MaxHeap)
	heap.Init(m.MinHeap)
	return m
}


func (this *MedianFinder) AddNum(num int)  {
	if this.Count == 0 {
		this.Mid = num
	} else {
		if num > this.Mid {
			heap.Push(this.MinHeap, num)
		} else {
			heap.Push(this.MaxHeap, num)
		}
		if this.MaxHeap.Len() - this.MinHeap.Len() > 1 {
			heap.Push(this.MinHeap, this.Mid)
			this.Mid = heap.Pop(this.MaxHeap).(int)
		} else if this.MinHeap.Len() - this.MaxHeap.Len() > 1 {
			heap.Push(this.MaxHeap, this.Mid)
			this.Mid = heap.Pop(this.MinHeap).(int)
		}
	}
	this.Count++
}


func (this *MedianFinder) FindMedian() float64 {
	if this.MaxHeap.Len() == this.MinHeap.Len() {
		return float64(this.Mid)
	} else if this.MaxHeap.Len() > this.MinHeap.Len() {
		return (float64(this.MaxHeap.IntSlice[0]) + float64(this.Mid))/2
	} else {
		return (float64(this.MinHeap.IntSlice[0]) + float64(this.Mid))/2
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
