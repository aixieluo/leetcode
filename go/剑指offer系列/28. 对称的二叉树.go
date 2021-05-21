package main

import "fmt"

func main()  {
	fmt.Println(isSymmetric(CreateTree([]int{1,2,2,3,4,4,3})))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(nums []int) *TreeNode {
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

func isSymmetric(root *TreeNode) bool {
	var check func(node *TreeNode, node2 *TreeNode) bool
	check = func(node *TreeNode, node2 *TreeNode) bool {
		if node == nil && node2 == nil {
			return true
		}
		if node == nil || node2 == nil {
			return false
		}
		return node.Val == node2.Val && check(node.Left, node2.Right) && check(node.Right, node2.Left)
	}
	return check(root, root)
}
