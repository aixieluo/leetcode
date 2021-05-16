// 给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。
//
//进阶：你可以在 O(n) 的时间解决这个问题吗？
//
// 
//
//示例 1：
//
//输入：nums = [3,10,5,25,2,8]
//输出：28
//解释：最大运算结果是 5 XOR 25 = 28.
//示例 2：
//
//输入：nums = [0]
//输出：0
//示例 3：
//
//输入：nums = [2,4]
//输出：6
//示例 4：
//
//输入：nums = [8,10,2]
//输出：10
//示例 5：
//
//输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
//输出：127
// 
//
//提示：
//
//1 <= nums.length <= 2 * 104
//0 <= nums[i] <= 231 - 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array

package main

import "fmt"

func main()  {
	fmt.Println(findMaximumXOR([]int{1, 12}))
	fmt.Println(findMaximumXOR([]int{3,10,5,25,2,8}))
}

const h = 30

func findMaximumXOR(nums []int) int {
    root := &Trie{}
    ans := 0
    for i:=1; i< len(nums); i++ {
        root.Add(nums[i-1])
        ans = max(ans, root.check(nums[i]))
    }
    return ans
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

type Trie struct {
    Left, Right *Trie
}

func (t *Trie) Add(num int) {
    node := t
    for i:=h; i >= 0; i-- {
        f := num >> i & 1
        if f == 0 {
            if node.Left == nil {
                node.Left = &Trie{}
            }
            node = node.Left
        } else {
            if node.Right == nil {
                node.Right = &Trie{}
            }
            node = node.Right
        }
    }
}

func (t *Trie) check(num int) int {
    node := t
    sum :=0
    for i:= h; i>=0; i-- {
        f := num >> i & 1
        if f == 0 {
            if node.Right != nil {
                node = node.Right
                sum = 2 * sum +1
            } else {
                sum = 2* sum
                node = node.Left
            }
        } else {
            if node.Left != nil {
                node = node.Left
                sum = 2* sum +1
            } else {
                sum = 2* sum
                node = node.Right
            }
        }
    }
    return sum
}