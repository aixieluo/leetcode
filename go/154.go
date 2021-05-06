package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{3, 3, 1, 3}))
}

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for ; low < high; {
		mid := (high+low)/2
		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return nums[low]
}

func findMin2(nums []int) int {
	low, high := 0, len(nums)-1
	for ; low < high; {
		mid := (low + high) / 2
		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low++
		} else {
			high--
			low++
		}
	}
	return nums[low]
}
