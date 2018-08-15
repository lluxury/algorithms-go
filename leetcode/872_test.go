package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_872(t *testing.T) {
	test.Runs(t, leafSimilar, []*test.Case{
		{Input: `(3, (5, 6, (2, 7, 4)), (1, 9, 8)) \n (3, (5, 6, 7), (1, 4, (2, 9, 8)))`, Output: `true`}, // [3,5,1,6,2,9,8,null,null,7,4]  \n [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
	})
}
