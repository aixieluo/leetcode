package main

import "fmt"

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/

func main()  {
	fmt.Println(verifyPostorder([]int{10,3,2,6,5}))
}

func verifyPostorder(postorder []int) bool {
	length := len(postorder)
	if  length <2 {
		return true
	}
	rootVal := postorder[length -1]
	i := 0
	for ;i< length; i++ {
		if rootVal <= postorder[i] {
			break
		}
	}
	left := i
	for ;i < length; i++ {
		if rootVal > postorder[i] {
			return false
		}
	}
	return verifyPostorder(postorder[:left]) && verifyPostorder(postorder[left:length-1])
}

