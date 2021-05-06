package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3 ,nil}}}}}
	fmt.Println(deleteDuplicates(head))
}

func deleteDuplicates(head *ListNode) *ListNode {
	var ddl func (head *ListNode)
	ddl = func (head *ListNode) {
		if head == nil || head.Next == nil {
			return
		}
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			ddl(head)
		} else {
			ddl(head.Next)
		}
	}

	ddl(head)
	return head
}
