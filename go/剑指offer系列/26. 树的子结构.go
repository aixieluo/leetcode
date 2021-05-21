package main

import "fmt"
import "utils"

func main() {
	fmt.Println(isSubStructure(untils.CreateTree([]int{1, 0, 1, -4, -3}), untils.CreateTree([]int{1, -4})), false)
	fmt.Println(isSubStructure(untils.CreateTree([]int{10, 12, 6, 8, 3, 11}), untils.CreateTree([]int{10, 12, 6, 8})), true)
	fmt.Println(isSubStructure(untils.CreateTree([]int{4, 2, 3, 4, 5, 6, 7, 8, 9}), untils.CreateTree([]int{4, 8, 9})), true)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	return isSub(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSub(p, c *TreeNode) bool {
	if c == nil {
		return true
	}
	if p == nil {
		return false
	}
	return p.Val == c.Val && isSub(p.Left, c.Left) && isSub(p.Right, c.Right)
}
