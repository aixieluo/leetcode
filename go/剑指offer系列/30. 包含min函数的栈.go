package main

func main()  {

}

type MinStack struct {
	stack1, stack2 []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	this.stack1 = append(this.stack1, x)
	if len(this.stack2) == 0 || this.stack2[len(this.stack2)-1] >= x {
		this.stack2 = append(this.stack2, x)
	}
}


func (this *MinStack) Pop()  {
	val := this.Top()
	if this.stack2[len(this.stack2)-1] == val {
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
	this.stack1 = this.stack1[:len(this.stack1)-1]
}


func (this *MinStack) Top() int {
	if len(this.stack1) >0 {
		return this.stack1[len(this.stack1)-1]
	}
	return 0
}


func (this *MinStack) Min() int {
	if len(this.stack2) >0 {
		return this.stack2[len(this.stack2)-1]
	}
	return 0
}


