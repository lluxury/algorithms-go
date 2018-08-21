package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_692(t *testing.T) {
	test.Runs(t, topKFrequent, []*test.Case{
		{Input: `["i", "love", "leetcode", "i", "love", "coding"] \n 2`, Output: `["i", "love"]`},
		{Input: `["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"] \n 4`, Output: `["the", "is", "sunny", "day"]`},
	})
}
