package main

import "fmt"

func main()  {
	fmt.Println(minArray([]int{4,5,6,7,0,1,2}))
}

func minArray(numbers []int) int {
    low, high := 0, len(numbers) - 1
	for low < high {
        pivot := (low + high) / 2
        if numbers[high] < numbers[pivot] {
            low = pivot + 1
        } else if numbers[high] > numbers[pivot] {
            high = pivot
        } else {
            high--
        }
    }
    return numbers[low]
}

func minArray2(numbers []int) int {
	if numbers[len(numbers)-1] > numbers[0] || len(numbers) == 1 {
		return numbers[0]
	}
	i := len(numbers) / 2
	if numbers[0] > numbers[i] {
		return minArray(numbers[1:])
	}
	return min(minArray(numbers[:i]), minArray(numbers[i:]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}