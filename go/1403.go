package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minSubsequence([]int{6}))
	fmt.Println(minSubsequence([]int{4, 3, 10, 9, 8}))
}

func minSubsequence(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	allSum := sumInts(nums)
	for i := 1; i < len(nums); i++ {
		if sumInts(nums[0:i]) > allSum-sumInts(nums[0:i]) {
			return nums[0:i]
		}
	}
	return nums
}

func sumInts(ns []int) int {
	var s int
	for _, v := range ns {
		s += v
	}
	return s
}
