package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSequenceSum(t *testing.T) {
	as := assert.New(t)

	as.Equal(20, MaxSequenceSum([]int{-2, 11, -4, 13, -5, -2}))
}
