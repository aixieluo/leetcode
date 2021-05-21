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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := make([][]int, 0)
	nodes := []*TreeNode{root}
	deep := []int{0}
	for len(nodes) > 0 {
		node := nodes[0]
		dep := deep[0]
		if len(ans) == dep {
			ans = append(ans, []int{node.Val})
		} else {
			ans[dep] = append(ans[dep], node.Val)
		}
		nodes = nodes[1:]
		deep = deep[1:]
		if node.Left != nil {
			nodes = append(nodes, node.Left)
			deep = append(deep, dep+1)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
			deep = append(deep, dep+1)
		}
	}
	return ans
}
