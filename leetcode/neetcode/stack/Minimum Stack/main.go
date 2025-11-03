package main

import "math"

type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	min := math.MaxInt

	for i := range this.stack {
		if min > this.stack[i] {
			min = this.stack[i]
		}
	}
	return min
}

func main() {

	minStack := Constructor()
	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(0)
	minStack.GetMin()
	minStack.Pop()
	minStack.Top()
	minStack.GetMin()
}
