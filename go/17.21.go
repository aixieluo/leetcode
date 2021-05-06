package main

import "fmt"

func main() {
	fmt.Println(trap2([]int{2, 0, 2}))
}

func trap1(height []int) int {
	n := len(height)
	if n < 1 {
		return 0
	}
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	rightMax[n-1] = height[n-1]
	for j := n - 2; j >= 0; j-- {
		rightMax[j] = max(height[j], rightMax[j+1])
	}
	sum := 0
	for k, v := range height {
		sum += min(leftMax[k], rightMax[k]) - v
	}
	return sum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func trap2(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}
	leftMax, rightMax, sum := 0, 0, 0
	for left, right := 0, n-1; left < right; {
		if height[left] > leftMax {
			leftMax = height[left]
		} else {
			sum += leftMax - height[left]
		}
		if height[right] > rightMax {
			rightMax = height[right]
		} else {
			sum += rightMax - height[right]
		}
		if leftMax < rightMax {
			left++
		} else {
			right--
		}
	}
	return sum
}
