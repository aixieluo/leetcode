package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func search2(nums []int, target int) int {
	for k, v := range nums {
		if v == target {
			return k
		}
	}
	return -1
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 判断是否在前半部分查找
		if (nums[left] <= target && target <= nums[mid]) ||
			(nums[mid] <= nums[right] &&
				(target < nums[mid] || target > nums[right])) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
