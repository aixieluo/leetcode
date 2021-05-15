// 输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
// 
//
//例如，给出
//
//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//返回如下的二叉树：
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
// 
//
//限制：
//
//0 <= 节点个数 <= 5000
//
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := createNode(preorder[0])
	if len(preorder) == 1 {
		return root
	}
	pivot := 0
	for k,v :=range inorder {
		if v== preorder[0] {
			pivot = k
			break
		}
	}
	root.Left = buildTree(preorder[1:pivot+1], inorder[:pivot])
	root.Right = buildTree(preorder[pivot+1:], inorder[pivot+1:])
	return root
}