package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	as := assert.New(t)

	stack := NewStack()
	as.True(stack.IsEmpty())
	stack.Push(2)
	stack.Push(3)
	as.False(stack.IsEmpty())
	as.Equal(3, stack.Pop())
	as.Equal(2, stack.Peek())
	as.False(stack.IsEmpty())
	as.Equal(2, stack.Pop())
	as.True(stack.IsEmpty())
}

func Test_stack_ref_clone(t *testing.T) {
	as := assert.New(t)

	{
		var stackRef = func(stack *Stack) { stack.Push(100) }

		stack := NewStack()
		stackRef(stack)
		as.Equal(100, stack.Pop())
	}

	{
		s1 := NewStack()
		s1.Push(2)
		s1.Push(3)
		s2 := s1.Clone()

		s1.Push(4)
		as.Equal("4 | 3 | 2",s1.String())
		as.Equal("3 | 2",s2.String())
	}
}
