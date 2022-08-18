//给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。
//
//
//
//示例 1：
//
//
//
//输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
//输出：15
//示例 2：
//
//输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
//输出：19
//
//
//提示：
//
//树中节点数目在范围 [1, 104] 之间。
//1 <= Node.val <= 100
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/deepest-leaves-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import "fmt"

func main() {
	fmt.Println(deepestLeavesSum(&TreeNode{Val: 10, Left: &TreeNode{Val: 11}, Right: &TreeNode{Val: 12}}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum1(root *TreeNode) int {
	mp := map[int]int{}
	sumNode(mp, root, 0)
	max := 0
	for k := range mp {
		if k > max {
			max = k
		}
	}
	return mp[max]
}

func sumNode(mp map[int]int, root *TreeNode, deep int) {
	if root == nil {
		return
	}
	mp[deep] += root.Val
	sumNode(mp, root.Left, deep+1)
	sumNode(mp, root.Right, deep+1)
}

func deepestLeavesSum(root *TreeNode) (sum int) {
	maxLevel := -1
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level > maxLevel {
			maxLevel = level
			sum = node.Val
		} else if level == maxLevel {
			sum += node.Val
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	return
}
