package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
	ans := make([][]int, 0)
	var dfs func(node *TreeNode, t int)
	nodes := make([]int, 0)
	dfs = func(node *TreeNode, t int) {
		if node == nil {
			return
		}
		nodes = append(nodes, node.Val)
		if t-node.Val == 0 && node.Left == nil && node.Right == nil {
			ans = append(ans, append([]int{}, nodes...))
		}
		dfs(node.Left, t-node.Val)
		dfs(node.Right, t-node.Val)
		nodes = nodes[:len(nodes)-1]
	}
	dfs(root, target)
	return ans
}
