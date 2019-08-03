package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_102(t *testing.T) {
	test.Runs(t, levelOrder, []*test.Case{
		{Input: `(3,9,(20,15,7))`, Output: `[[3], [9, 20], [15, 7]]`},
	})
}

// actual  : [][]int{[]int{3}, []int{9, 15}, []int{20, 7}}
// 貌似默认的遍历次序不一样,要怎么解决, 看代码是前序的,但测试怎么过的啊
