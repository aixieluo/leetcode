// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
//
// 注意：本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/ 相同
//
//  
//
// 示例 1：
//
//
// 输入：root = [4,2,6,1,3]
// 输出：1
// 示例 2：
//
//
// 输入：root = [1,0,48,null,null,12,49]
// 输出：1
//  
//
// 提示：
//
// 树中节点数目在范围 [2, 100] 内
// 0 <= Node.val <= 105
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minDiffInBST(&TreeNode{100, &TreeNode{10, &TreeNode{99, nil, nil}, &TreeNode{13, nil, nil}}, nil}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func minDiffInBST2(root *TreeNode) int {
	var f func(root *TreeNode, point *TreeNode)
	res := root.Val
	f = func(root *TreeNode, point *TreeNode) {
		if point == nil {
			return
		}
		if root != point {
			res = min(res, root.Val-point.Val)
		}
		f(root, point.Left)
		f(root, point.Right)
	}
	var f2 func(root *TreeNode)
	f2 = func(r *TreeNode) {
		if r == nil {
			return
		}
		f(r, root.Left)
		f(r, root.Right)
		f2(r.Left)
		f2(r.Right)
	}
	f2(root)
	return res
}

func min(a, b int) int {
	a, b = abs(a), abs(b)
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
