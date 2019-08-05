package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_1029(t *testing.T) {
	test.Runs(t, twoCitySchedCost, []*test.Case{
	    {Input: `[[10,20],[30,200],[400,50],[30,20]]`, Output: `110`},
	})
}

