// 设计实现双端队列。
//
//实现 MyCircularDeque 类:
//
//MyCircularDeque(int k) ：构造函数,双端队列最大为 k 。
//boolean insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true ，否则返回 false 。
//boolean insertLast() ：将一个元素添加到双端队列尾部。如果操作成功返回 true ，否则返回 false 。
//boolean deleteFront() ：从双端队列头部删除一个元素。 如果操作成功返回 true ，否则返回 false 。
//boolean deleteLast() ：从双端队列尾部删除一个元素。如果操作成功返回 true ，否则返回 false 。
//int getFront() )：从双端队列头部获得一个元素。如果双端队列为空，返回 -1 。
//int getRear() ：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1 。
//boolean isEmpty() ：若双端队列为空，则返回 true ，否则返回 false  。
//boolean isFull() ：若双端队列满了，则返回 true ，否则返回 false 。
//
//
//示例 1：
//
//输入
//["MyCircularDeque", "insertLast", "insertLast", "insertFront", "insertFront", "getRear", "isFull", "deleteLast", "insertFront", "getFront"]
//[[3], [1], [2], [3], [4], [], [], [], [4], []]
//输出
//[null, true, true, true, false, 2, true, true, true, 4]
//
//解释
//MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
//circularDeque.insertLast(1);			        // 返回 true
//circularDeque.insertLast(2);			        // 返回 true
//circularDeque.insertFront(3);			        // 返回 true
//circularDeque.insertFront(4);			        // 已经满了，返回 false
//circularDeque.getRear();  				// 返回 2
//circularDeque.isFull();				        // 返回 true
//circularDeque.deleteLast();			        // 返回 true
//circularDeque.insertFront(4);			        // 返回 true
//circularDeque.getFront();				// 返回 4
//
//
//
//提示：
//
//1 <= k <= 1000
//0 <= value <= 1000
//insertFront, insertLast, deleteFront, deleteLast, getFront, getRear, isEmpty, isFull  调用次数不大于 2000 次
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/design-circular-deque
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
)

func main() {
	circularDeque := Constructor(3)           // 设置容量大小为3
	fmt.Println(circularDeque.InsertLast(1))  // 返回 true
	fmt.Println(circularDeque.InsertLast(2))  // 返回 true
	fmt.Println(circularDeque.InsertFront(3)) // 返回 true
	fmt.Println(circularDeque.InsertFront(4)) // 已经满了，返回 false
	fmt.Println(circularDeque.GetRear())      // 返回 2
	fmt.Println(circularDeque.IsFull())       // 返回 true
	fmt.Println(circularDeque.DeleteLast())   // 返回 true
	fmt.Println(circularDeque.InsertFront(4)) // 返回 true
	fmt.Println(circularDeque.GetFront())
}

type MyCircularDeque struct {
	front, rear int
	list        []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{list: make([]int, k+1)}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.front = (this.front - 1 + len(this.list)) % len(this.list)
	this.list[this.front] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.list[this.rear] = value
	this.rear = (this.rear + 1) % len(this.list)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1 + len(this.list)) % len(this.list)
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.rear = (this.rear - 1 + len(this.list)) % len(this.list)
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.list[this.front%len(this.list)]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.list[(this.rear-1+len(this.list))%len(this.list)]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.front == (this.rear)%len(this.list)
}

func (this *MyCircularDeque) IsFull() bool {
	if this.front > this.rear {
		return this.front-this.rear == 1
	}
	return this.rear-this.front == len(this.list)-1
}
