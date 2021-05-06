// 给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//
//  
//
// 示例 1：
//
//
// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
// 示例 2：
//
// 输入：root = []
// 输出：[]
// 示例 3：
//
// 输入：root = [1]
// 输出：[1]
// 示例 4：
//
//
// 输入：root = [1,2]
// 输出：[2,1]
// 示例 5：
//
//
// 输入：root = [1,null,2]
// 输出：[1,2]
//  
//
// 提示：
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal

package main

func main()  {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		f(node.Left)
		res = append(res, node.Val)
		f(node.Right)
	}
	f(root)
	return res
}
