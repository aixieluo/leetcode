package main

func main()  {

}

func findRepeatNumber(nums []int) int {
	hash:= make(map[int]bool)
	for _,v :=range nums {
		if hash[v] {
			return v
		} else {
			hash[v] = true
		}
	}
	return 0
}
