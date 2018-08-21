package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_232(t *testing.T) {
	obj := Constructor_232()
	obj.Push(1)
	obj.Push(2)
	assert.Equal(t, 1, obj.Pop())
	assert.Equal(t, 2, obj.Pop())
	assert.True(t, obj.Empty())

}
