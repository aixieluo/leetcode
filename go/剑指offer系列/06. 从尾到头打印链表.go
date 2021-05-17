// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//  
//
// 示例 1：
//
// 输入：head = [1,3,2]
// 输出：[2,3,1]
//  
//
// 限制：
//
// 0 <= 链表长度 <= 10000
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	ans :=make([]int, 0, 100)
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	for l,r:=0, len(ans)-1; l< r; l,r= l+1, r-1  {
		ans[l], ans[r] = ans[r], ans[l]
	}
	return ans
}
