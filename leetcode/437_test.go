package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_437(t *testing.T) {
	test.Runs(t, pathSum_437, []*test.Case{
		{Input: `(1, (2,,(3,,(4,,5)))) \n 3`, Output: `2`},
		{Input: `(0,1,1) \n 0`, Output: `1`},
		{Input: `(1) \n 0`, Output: `0`},
		{Input: `(-2,,-3) \n -5`, Output: `1`},
		{Input: `(1,-2,-3) \n -2`, Output: `2`},
		{Input: `(10, (5, (3,3,-2), (2,,1)), (-3,,11)) \n 8`, Output: `3`},
		{Input: `(5, (4, (11,7,2)), (8, (13,5,1), 4)) \n 22`, Output: `3`},
	})
}

// 比较链表，双向链表，array，map，tree，heap，queue，stack
// 在插入，删除，获取长度，获取第第一个数据、最小数据、第n个数据，的对比
// 存储空间，
