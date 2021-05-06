package main

import (
	"fmt"
	"sort"
)

func main() {
	a:= []int{10,3,2}
	fmt.Println(sort.SearchInts(a, 5))
	fmt.Println(lengthOfLIS([]int{4,10,4,3,8,9}))
}

func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		i := sort.SearchInts(dp, num) //min_index
		if i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
	}
	return len(dp)
}
