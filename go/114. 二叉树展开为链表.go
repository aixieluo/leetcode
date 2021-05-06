// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
//
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
//  
//
// 示例 1：
//
//
// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
// 示例 2：
//
// 输入：root = []
// 输出：[]
// 示例 3：
//
// 输入：root = [0]
// 输出：[0]
//  
//
// 提示：
//
// 树中结点数在范围 [0, 2000] 内
// -100 <= Node.val <= 100
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list

package main

import "fmt"

func main() {
	head := &TreeNode{1, &TreeNode{2, &TreeNode{Val: 3}, &TreeNode{Val: 4}}, &TreeNode{5, nil, &TreeNode{Val: 6}}}
	flatten(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Right
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var list []*TreeNode
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		list = append(list, node)
		f(node.Left)
		f(node.Right)
	}
	f(root)
	for index := 0; index <= len(list)-2; index++ {
		pre, cur := list[index], list[index+1]
		pre.Left, pre.Right = nil, cur
	}
}

func flatten2(root *TreeNode) {
	var head []int
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		head = append(head, node.Val)
		f(node.Left)
		f(node.Right)
	}
	f(root)
	for k, v := range head {
		if k == 0 {
			root.Left = nil
			continue
		}
		h := &TreeNode{Val: v}
		root.Right = h
		root = root.Right
	}
}
