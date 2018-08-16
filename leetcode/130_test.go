package leetcode

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_130(t *testing.T) {
	{
		x := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
		solve(x)
		assert.Equal(t, [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}, x)
	}

	{
		x := [][]byte{}
		solve(x)
		assert.Equal(t, [][]byte{}, x)
	}
}
