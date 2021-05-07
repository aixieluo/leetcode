package main

import "fmt"

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
	// fmt.Println(twoSum([]int{2,7,11,15}, 9))
	// fmt.Println(twoSum([]int{3,2,4}, 6))
	// fmt.Println(twoSum([]int{3,3}, 6))
	// fmt.Println(climbStairs(2))
	// fmt.Println(climbStairs(3))
	// fmt.Println(climbStairs(4))
}

func minimumTotal(triangle [][]int) int {
	high := len(triangle)
	dp := make([][]int, high)
	for i:=0; i< high ;i++  {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = triangle[0][0]
	for x := 1; x < high; x++ {
		for y := 0; y <= x; y++ {
			if y == 0 {
				dp[x][y] = triangle[x][y] + dp[x-1][y]
			} else if  y == x {
				dp[x][y] = triangle[x][y] + dp[x-1][y-1]
			} else {
				dp[x][y] = min(dp[x-1][y], dp[x-1][y-1]) + triangle[x][y]
			}
		}
	}
	m := dp[high-1][0]
	for i:=1; i< high ; i++  {
		m = min(m, dp[high-1][i])
	}
	return m
}

func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}

func twoSum(nums []int, target int) []int {
	for first := 0; first < len(nums)-1; first++ {
		for second := first + 1; second < len(nums); second++ {
			if nums[first]+nums[second] == target {
				return []int{first, second}
			}
		}
	}
	return make([]int, 0)
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
