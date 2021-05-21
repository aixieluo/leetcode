package main

func main() {

}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed))
	slow := 0
	for _, v := range pushed {
		for i:= len(stack)-1; i >=0; i-- {
			if stack[i] == popped[slow] {
				slow++
				stack = stack[:len(stack)-1]
			}
		}
		if v == popped[slow] {
			slow++
		} else {
			stack = append(stack, v)
		}
	}
	for k,v:=range stack {
		if v != popped[len(popped)-k-1] {
			return false
		}
	}
	return true
}
