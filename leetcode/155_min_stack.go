package leetcode

/*
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (32.63%)
 * Total Accepted:    210.8K
 * Total Submissions: 644.7K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 * 
 * 
 * push(x) -- Push element x onto stack.
 * 
 * 
 * pop() -- Removes the element on top of the stack.
 * 
 * 
 * top() -- Get the top element.
 * 
 * 
 * getMin() -- Retrieve the minimum element in the stack.
 * 
 * 
 * 
 * 
 * Example:
 * 
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> Returns -3.
 * minStack.pop();
 * minStack.top();      --> Returns 0.
 * minStack.getMin();   --> Returns -2.
 * 
 * 
 */



/*
> https://leetcode.com/problems/min-stack/description/


Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.

Example:
```
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
```

实现上面描述的接口

思路
* 一个list存数据
* 一个list存最小数据
* *最小栈的数据会越来越小*


*/

type MinStack struct {
	list []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.list = append(this.list, x)

	if l := len(this.min); l == 0 || x <= this.min[l-1] {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	l1 := len(this.list)
	l2 := len(this.min)

	if l1 == 0 {
		return
	}

	if this.list[l1-1] == this.min[l2-1] {
		this.min = this.min[:l2-1]
	}

	this.list = this.list[:l1-1]
}

func (this *MinStack) Top() int {
	if l := len(this.list); l == 0 {
		return 0
	} else {
		return this.list[l-1]
	}
}

func (this *MinStack) GetMin() int {
	if l := len(this.min); l == 0 {
		return 0
	} else {
		return this.min[l-1]
	}
}
