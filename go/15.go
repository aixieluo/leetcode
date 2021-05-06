package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for first := 0; first < len(nums)-2; first++ {
		if first > 0 && (nums[first-1] == nums[first]) {
			continue
		}
		for second, third := first+1, len(nums)-1; second < third; second++ {
			if second > first+1 && (nums[second-1] == nums[second]) {
				continue
			}
			for ; nums[first]+nums[second]+nums[third] > 0; {
				third--
				if second == third {
					break
				}
			}
			if nums[first]+nums[second]+nums[third] == 0 {
				fmt.Println(nums[first]+nums[second]+nums[third])
				res = append(res, append([]int(nil), nums[first], nums[second], nums[third]))
			}
		}
	}
	return res
}

func check(nums []int, f, s, t int) ([]int, bool) {
	if nums[f]+nums[s]+nums[t] == 0 {
		return []int{nums[f], nums[s], nums[t]}, true
	}
	return nil, false
}

func sum(nums []int, f, s, t int) int {
	return nums[f] + nums[s] + nums[t]
}
