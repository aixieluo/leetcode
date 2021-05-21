package main

import (
	"fmt"
)

func main() {
	a:= mirrorTree(CreateTree([]int{4,2,7,1,3,6,9}))
	fmt.Println(a)
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

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}

