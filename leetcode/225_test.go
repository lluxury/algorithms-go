package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_225(t *testing.T) {
	as := assert.New(t)

	stack := Constructor_225()
	as.True(stack.Empty())
	stack.Push(1)
	as.False(stack.Empty())
	stack.Push(2)
	as.False(stack.Empty())
	as.Equal(2, stack.Pop())
	as.Equal(1, stack.Pop())
	as.True(stack.Empty())
}
