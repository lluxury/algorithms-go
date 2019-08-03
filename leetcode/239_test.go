package leetcode

import (
	"testing"
	// "github.com/stretchr/testify/assert"
	"github.com/lluxury/algorithms-go/test"
)

func Test_239(t *testing.T) {
	for _, f := range []func(nums []int, k int) []int{maxSlidingWindow} {
		test.Runs(t, f, []*test.Case{
			{Input: (`[1,3,-1,-3,5,3,6,7] \n 3`), Output: `[3,3,5,5,6,7]`},
		})
	}

}

// leetcode 测试通过,本地有些不兼容,后面处理
