package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_264(t *testing.T) {
	test.Runs(t, nthUglyNumber, []*test.Case{
		{Input: "0", Output: "-1"},
		{Input: "1", Output: "1"},
		{Input: "10", Output: "12"},
		{Input: "100", Output: "1536"},
	})
}
