package main

func main() {

}

func exchange(nums []int) []int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast]%2 == 1 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
	}
	return nums
}
