package main

func main()  {

}

func majorityElement(nums []int) int {
	hash := make(map[int]int)
	muchNum := nums[0]
	count := 1
	for _,v :=range nums {
		hash[v]++
		if hash[v] > count {
			count = hash[v]
			muchNum = v
		}
	}
	return muchNum
}
