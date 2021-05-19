package main

import "fmt"

func main() {
	fmt.Println(isSubStructure(createTree([]int{1, 0, 1, -4, -3}), createTree([]int{1, -4})), false)
	fmt.Println(isSubStructure(createTree([]int{10, 12, 6, 8, 3, 11}), createTree([]int{10, 12, 6, 8})), true)
	fmt.Println(isSubStructure(createTree([]int{4, 2, 3, 4, 5, 6, 7, 8, 9}), createTree([]int{4, 8, 9})), true)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree(nums []int) *TreeNode {
	nodes := make([]*TreeNode, 0, 10)
	for _, v := range nums {
		nodes = append(nodes, &TreeNode{Val: v})
	}
	for i := 1; i < len(nodes); i++ {
		if i%2 == 1 {
			nodes[i/2].Left = nodes[i]
		} else {
			nodes[(i-1)/2].Right = nodes[i]
		}
	}
	return nodes[0]
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
