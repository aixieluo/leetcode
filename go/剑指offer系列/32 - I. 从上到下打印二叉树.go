package main

func main()  {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0)
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		node := nodes[0]
		ans = append(ans, node.Val)
		nodes = nodes[1:]
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}
	return ans
}
