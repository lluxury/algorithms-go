package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [232] Implement Queue using Stacks
 *
 * https://leetcode.com/problems/implement-queue-using-stacks/description/
 *
 * algorithms
 * Easy (39.24%)
 * Total Accepted:    115.8K
 * Total Submissions: 294.6K
 * Testcase Example:  '["MyQueue","push","push","peek","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * Implement the following operations of a queue using stacks.
 *
 *
 * push(x) -- Push element x to the back of queue.
 * pop() -- Removes the element from in front of queue.
 * peek() -- Get the front element.
 * empty() -- Return whether the queue is empty.
 *
 *
 * Example:
 *
 *
 * MyQueue queue = new MyQueue();
 *
 * queue.push(1);
 * queue.push(2);
 * queue.peek();  // returns 1
 * queue.pop();   // returns 1
 * queue.empty(); // returns false
 *
 * Notes:
 *
 *
 * You must use only standard operations of a stack -- which means only push to
 * top, peek/pop from top, size, and is empty operations are valid.
 * Depending on your language, stack may not be supported natively. You may
 * simulate a stack by using a list or deque (double-ended queue), as long as
 * you use only standard operations of a stack.
 * You may assume that all operations are valid (for example, no pop or peek
 * operations will be called on an empty queue).
 *
 *
 */

/*

用栈实现队列
  * 栈是后入先出
  * 队列是先入先出
  * 实现
    * 因为栈和队列是反的，所以只要将栈的数据反向，然后一个个出栈即可。
    * 所以需要两个栈，一个保存入栈的原始数据，一个保存反向之后的数据
    * 入队都放在第一个栈
    * 出队：如果第二个栈有数据，直接pop，否则将第一个栈的所有数据pop，然后push到第二个栈
*/
type MyQueue struct {
	s1 *lib.Stack
	s2 *lib.Stack
}

/** Initialize your data structure here. */
func Constructor_232() MyQueue {
	return MyQueue{
		s1: lib.NewStack(),
		s2: lib.NewStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() {
			this.s2.Push(this.s1.Pop())
		}
	}

	return this.s2.Pop().(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() {
			this.s2.Push(this.s1.Pop())
		}
	}

	return this.s2.Peek().(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.s1.IsEmpty() && this.s2.IsEmpty()
}
