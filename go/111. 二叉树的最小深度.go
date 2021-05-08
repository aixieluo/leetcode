// 给定一个二叉树，找出其最小深度。
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
// 说明：叶子节点是指没有子节点的节点。
//
//  
//
// 示例 1：
//
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：2
// 示例 2：
//
// 输入：root = [2,null,3,null,4,null,5,null,6]
// 输出：5
//  
//
// 提示：
//
// 树中节点数的范围在 [0, 105] 内
// -1000 <= Node.val <= 1000
//
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree

package main

import "fmt"

func main() {
	fmt.Println(minDepth(&TreeNode{
		3,
		&TreeNode{9, nil, nil},
		&TreeNode{
			20,
			&TreeNode{15, nil, nil},
			&TreeNode{7,nil, nil},
		},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	count := make([]int, 0)
	queue = append(queue, root)
	count = append(count, 1)
	for i:=0; i<len(queue); i++ {
		node := queue[i]
		deep := count[i]
		if node.Left == nil && node.Right == nil {
			return deep
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			count = append(count, deep+1)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			count = append(count, deep+1)
		}
	}
	return 0
}
