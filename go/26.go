package main

import "fmt"

func main()  {
	fmt.Println(removeDuplicates([]int{1,1,2}))
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow:= 0
	for index:=1; index < len(nums); index++  {
		if nums[slow] != nums[index] {
			slow++
			nums[slow] = nums[index]
		}
	}
	return slow +1
}
