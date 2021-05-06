package main

import "fmt"

func main() {
	m:= Constructor()
	m.AddNum(1)
	fmt.Println(m.FindMedian())
	m.AddNum(2)
	m.AddNum(3)
	fmt.Println(m.FindMedian())
	// fmt.Println(checkValidString("((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()"))
	// fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

type MedianFinder struct {
	List  []int
	Left  int
	Right int
	Mid   int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) Len() int {
	return len(this.List)
}

func (this *MedianFinder) AddNum(num int) {
	this.List = append(this.List, num)
	if this.Len() == 1 {
		this.Left, this.Right, this.Mid = num, -1, num
		return
	}
	if this.Len()%2 == 0 {
		if this.Left < num && this.Right > num {
			this.Left = num
		}
		if this.Right < 0 {
			this.Right = max(num,this.Left)
			this.Left = min(num, this.Left)
		}
	} else {
		if this.Mid < num {
			this.Mid = min(num, this.Right)
		} else {
			this.Mid = max(num, this.Left)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (this *MedianFinder) FindMedian() float64 {
	list := this.List
	l := len(list)
	if l%2 == 0 {
		return (float64(this.Left) + float64(this.Right)) / 2
	} else {
		return float64(this.Mid)
	}
}

func checkValidString(s string) bool {
	stack := make([]rune, 0, len(s))
	for _, v := range s {
		if v == ')' {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if last == '*' && len(stack) > 0 {
				for i := len(stack) - 1; i >= 0; i-- {
					if stack[i] == '(' {
						stack[i] = '*'
						break
					}
				}
			}
		} else {
			stack = append(stack, v)
		}
	}
	stack2 := make([]rune, 0)
	for _, v := range stack {
		if v == '(' {
			stack2 = append(stack2, v)
		} else {
			if len(stack2) > 0 {
				stack2 = stack2[:len(stack2)-1]
			}
		}
	}
	return len(stack2) == 0
}

func maxSubArray(nums []int) int {
	prev, res := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		prev = max(v, prev+v)
		res = max(res, prev)
	}
	return res
}
