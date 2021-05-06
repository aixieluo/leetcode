package main

import "fmt"

func main() {
	fmt.Println(search([]int{1,0,1,1,1}, 0))
}

func search(nums []int, target int) bool {
	if len(nums) ==0 {
		return false
	}
	if len(nums) == 1{
		return nums[0] == target
	}
	left, right := 0, len(nums) -1
	for left <= right {
		mid:= (right +left) /2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[left] && nums[mid] == nums[right] {
			left++
			right--
			continue
		}
		if (nums[left] <= target && target < nums[mid]) ||
			((target > nums[right] || target < nums[mid]) && (nums[mid] <= nums[right])) {
			right = mid -1
		} else {
			left = mid +1
		}
	}
	return false
}
