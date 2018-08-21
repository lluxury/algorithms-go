package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_830(t *testing.T) {
	test.Runs(t, largeGroupPositions, []*test.Case{
		{Input: `abbxxxzxzzy`, Output: `[[3,5]]`},
		{Input: `abbxxzxxzzy`, Output: `[]`},
		{Input: `abbxxxxzzy`, Output: `[[3,6]]`},
		{Input: `abc`, Output: `[]`},
		{Input: `bababbabaa`, Output: `[]`},
		{Input: `aaa`, Output: `[[0,2]]`},
		{Input: `abcdddeeeeaabbbcd`, Output: `[[3,5],[6,9],[12,14]]`},
	})
}
