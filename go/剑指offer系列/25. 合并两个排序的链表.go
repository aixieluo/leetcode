package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val <= l2.Val {
		head = l1
		head.Next = mergeTwoLists(l1.Next, l2)
	} else {
		head = l2
		head.Next = mergeTwoLists(l1, l2.Next)
	}
	return head
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, ans *ListNode
	for l1 != nil || l2 != nil {
		node := &ListNode{}
		if l1 == nil {
			node = l2
			l2 = l2.Next
		} else if l2 == nil {
			node = l1
			l1 = l1.Next
		} else {
			if l1.Val <= l2.Val {
				node = l1
				l1 = l1.Next
			} else {
				node = l2
				l2 = l2.Next
			}
		}
		if head == nil {
			head = node
			ans = head
		} else {
			head.Next = node
			head = head.Next
		}
	}
	return ans
}
