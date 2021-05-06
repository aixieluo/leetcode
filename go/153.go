package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{2, 1}))
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
}

func findMin2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right, mid := 0, len(nums)-1, (len(nums)-1)/2
	for left < right {
		if nums[left] <= nums[mid] && nums[left] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
		mid = (left + right) / 2
	}
	return nums[left]
}

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] < nums[high] {
			high = mid
		} else {
			low++
		}
	}
	return nums[low]
}
