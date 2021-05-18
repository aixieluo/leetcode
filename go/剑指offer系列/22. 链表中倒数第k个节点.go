package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	k2 := 0
	slow := head
	for head != nil {
		if k2 == k {
			slow = slow.Next
		} else {
			k2++
		}
		head = head.Next
	}
	return slow
}
