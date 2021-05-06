package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,1,2,3,3}))
}

func removeDuplicates(nums []int) int {
	if l := len(nums); l <= 2 {
		return l
	}
	for index := 2;; {
		if index >= len(nums) {
			break
		}
		if nums[index] == nums[index-1] && nums[index] == nums[index-2] {
			nums = append(nums[:index], nums[index+1:]...)
		} else {
			index++
		}
	}
	return len(nums)
}
