// 你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
//
// 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
//
// 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
//
// 字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
//
//  
//
// 示例 1:
//
// 输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
// 输出：6
// 解释：
// 可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
// 注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
// 因为当拨动到 "0102" 时这个锁就会被锁定。
// 示例 2:
//
// 输入: deadends = ["8888"], target = "0009"
// 输出：1
// 解释：
// 把最后一位反向旋转一次即可 "0000" -> "0009"。
// 示例 3:
//
// 输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
// 输出：-1
// 解释：
// 无法旋转到目标数字且不被锁定。
// 示例 4:
//
// 输入: deadends = ["0000"], target = "8888"
// 输出：-1
//  
//
// 提示：
//
// 1 <= deadends.length <= 500
// deadends[i].length == 4
// target.length == 4
// target 不在 deadends 之中
// target 和 deadends[i] 仅由若干位数字组成
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/open-the-lock

package main

import (
	"fmt"
)

func main() {
	fmt.Println(openLock([]string{"8887","8889","8878","8898","8788","8988","7888","9888"}, "8888"))
	fmt.Println(openLock([]string{"8888"}, "0008"))
}

type lock struct {
	T []byte
	C int
	Mp map[string]bool
}

func (l lock) IsLock() bool  {
	return l.Mp[string(l.T)]
}

func (l lock) Up(index int) lock {
	nums := make([]byte, 4)
	copy(nums, l.T)
	if nums[index] == '9' {
		nums[index] = '0'
	} else {
		nums[index]++
	}
	l.T = nums
	l.C++
	return l
}

func (l lock) Down(index int) lock {
	nums := make([]byte, 4)
	copy(nums, l.T)
	if nums[index] == '0' {
		nums[index] = '9'
	} else {
		nums[index]--
	}
	l.T = nums
	l.C++
	return l
}

type queues struct {
	Ls []lock
	Mp map[string]bool
}

func (q *queues) App(l lock) {
	if !q.Mp[string(l.T)] && !l.IsLock() {
		q.Ls = append(q.Ls, l)
		q.Mp[string(l.T)] = true
	}
}

func (q * queues) Pop() lock {
	l := q.Ls[0]
	q.Ls = q.Ls[1:]
	return l
}

func (q queues) Len() int {
	return len(q.Ls)
}

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	mp := map[string]bool{}
	for _, v := range deadends {
		mp[v] = true
	}
	if mp["0000"] {
		return -1
	}
	queue := queues{Mp: map[string]bool{}}
	queue.App(lock{[]byte("0000"), 0, mp})
	ans := 1<<7 - 1
	for queue.Len() > 0 {
		q := queue.Pop()
		if string(q.T) == target {
			if q.C < ans {
				ans = q.C
			}
		}
		for i := 0; i < 4; i++ {
			up := q.Up(i)
			queue.App(up)
			down:= q.Down(i)
			queue.App(down)
		}
	}
	if ans == 1<<7-1 {
		return -1
	} else {
		return ans
	}
}
