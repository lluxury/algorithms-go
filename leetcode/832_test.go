package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_832(t *testing.T) {
	test.Runs(t, flipAndInvertImage, []*test.Case{
		//{Input: `[[1,1,0],[1,0,1],[0,0,0]]`, Output: `[[1,0,0],[0,1,0],[1,1,1]]`},
		{Input: `[[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]`, Output: `[[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]`},
	})
}

// 比较链表，双向链表，array，map，tree，heap，queue，stack
// 在插入，删除，获取长度，获取第第一个数据、最小数据、第n个数据，的对比
// 存储空间，
