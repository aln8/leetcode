/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
import "github.com/Marsyyh/go-datastructure/container/stack"

type MyQueue struct {
	inStack  stack.Stack
	outStack stack.Stack
	length   int
}

// Constructor to creat MyQueue
func Constructor() MyQueue {
	return MyQueue{stack.NewLinkedListStack(), stack.NewLinkedListStack(), 0}
}

// Push element x to the back of queue.
func (q *MyQueue) Push(x int) {
	q.inStack.Push(x)
	q.length++
}

// Pop the element from in front of queue and returns that element.
func (q *MyQueue) Pop() int {
	if !q.adjust() {
		return -1
	}
	q.length--
	if num, ok := q.outStack.Pop().(int); ok {
		return num
	}
	return -1
}

// Peek the front element. */
func (q *MyQueue) Peek() int {
	if !q.adjust() {
		return -1
	}
	q.adjust()
	if num, ok := q.outStack.Peek().(int); ok {
		return num
	}
	return -1
}

// Empty returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return q.length == 0
}

func (q *MyQueue) adjust() bool {
	if q.length == 0 {
		return false
	}
	if q.outStack.IsEmpty() {
		for q.inStack.Peek() != nil {
			q.outStack.Push(q.inStack.Pop())
		}
	}
	return true
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

