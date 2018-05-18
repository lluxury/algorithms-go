package lib

func Example_stack() {
	t := NewStack()
	ExamplePrint(t)
	ExamplePrint(t.IsEmpty())
	t.Push(2)
	t.Push(3)
	ExamplePrint(t.IsEmpty())
	ExamplePrint(t.Pop())
	ExamplePrint(t.Peek())
	ExamplePrint(t.IsEmpty())
	ExamplePrint(t)
	t.Pop()
	ExamplePrint(t.IsEmpty())

	// Output:
	// &{[]}
	// true
	// false
	// 3
	// 2
	// false
	// &{[2]}
	// true
}

func Example_stack_ref() {
	var stack_ref = func(stack *Stack) {
		stack.Push(100)
	}
	t := NewStack()
	ExamplePrint(t)
	stack_ref(t)
	ExamplePrint(t)

	// Output:
	// &{[]}
	// &{[100]}
}
