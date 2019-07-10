package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_131(t *testing.T) {
	test.Runs(t, partition, []*test.Case{
		{Input: `aab`, Output: `[[a,a,b],[aa,b]]`},
	})
}
