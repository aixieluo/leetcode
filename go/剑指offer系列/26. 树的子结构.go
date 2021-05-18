package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree(nums []int) *TreeNode  {
	nodes := make([]*TreeNode, 0, 10)
	for _,v:=range nums {
		nodes = append(nodes, &TreeNode{Val:v})
	}
	for i:=1; i< len(nodes); i++
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	if A.Val == B.Val {
		if B.Left == nil && B.Right == nil {
			return A.Left == nil && A.Right == nil
		}
		if B.Left == nil {
			return isSubStructure(A.Right, B.Right) && A.Left == nil
		} else if B.Right == nil {
			return isSubStructure(A.Left, B.Left) && A.Right == nil
		} else {
			return isSubStructure(A.Left, B.Left) && isSubStructure(A.Right, B.Right)
		}
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
