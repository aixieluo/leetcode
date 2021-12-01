package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println(search([]int{1,1,3,4,5,5,5,5,5}, 3))
	fmt.Println(search([]int{2,2}, 2))
}

func search(nums []int, target int) int {
	n1 := sort.SearchInts(nums, target)
	if n1 == len(nums) || nums[n1] != target {
		return 0
	}
	n2 := sort.SearchInts(nums, target+1)
	return n2-n1
}
