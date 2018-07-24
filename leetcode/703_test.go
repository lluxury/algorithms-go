package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_703(t *testing.T) {
	as := assert.New(t)

	c := Constructor_703(3, []int{4, 5, 8, 2})
	as.Equal(4, c.Add(3))
	as.Equal(5, c.Add(5))
	as.Equal(5, c.Add(10))
	as.Equal(8, c.Add(9))
	as.Equal(8, c.Add(4))
}
