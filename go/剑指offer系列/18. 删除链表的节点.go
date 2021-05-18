package main

func main()  {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	node := head.Next
	prev := head
	if prev.Val == val {
		return node
	}
	for node != nil {
		if node.Val == val {
			prev.Next = node.Next
		}
		prev = node
		node = node.Next
	}
	return head
}
