package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_630(t *testing.T) {
	test.Runs(t, scheduleCourse, []*test.Case{
		{Input: `[[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]`, Output: `3`},
		{Input: `[[5,5],[4,6],[2,6]]`, Output: `2`},
		{Input: `[[5,15],[3,19],[6,7],[2,10],[5,16],[8,14],[10,11],[2,19]]`, Output: `5`},
	})
}
