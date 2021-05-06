package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return append(children(nums), make([]int, 0))
}

func children(nums []int) [][]int {
	var result [][]int
	l := len(nums)
	if l >= 1 {
		first := 1
		res := children(nums[first:])
		result = append(result, append([]int(nil), nums[0]))
		for _, v := range res {
			if len(v) == 0 {
				continue
			}
			if v[0] != nums[0] {
				result = append(result, v)
			}
			result = append(result, append([]int{nums[0]}, v...))
		}
	}
	return result
}
