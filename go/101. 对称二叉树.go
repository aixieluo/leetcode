// 给定一个二叉树，检查它是否是镜像对称的。
//
//  
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
// 3  4 4  3
//  
//
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//  
//
// 进阶：
//
// 你可以运用递归和迭代两种方法解决这个问题吗？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/symmetric-tree

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
