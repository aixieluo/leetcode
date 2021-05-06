package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5,6,7}
	rotate(nums, 10)
	fmt.Println( nums)
}

func rotate1(nums []int, k int)  {
	k %= len(nums)
	for k,v := range append(nums[len(nums)-k:], nums[:len(nums)-k]...) {
		nums[k] = v
	}
}

func reverse2(nums []int) {
	for left, right := 0, len(nums) -1; left < len(nums) /2 ; left++ {
		nums[left], nums[right] = nums[right], nums[left]
		right--
	}
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse2(nums)
	reverse2(nums[:k])
	reverse2(nums[k:])
}
