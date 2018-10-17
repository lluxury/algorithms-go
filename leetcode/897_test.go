package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_897(t *testing.T) {
	test.Runs(t, increasingBST, []*test.Case{
		{Input: `(5, (3, (2,1,), 4), (6, ,(8,7,9)))`, Output: `(1,,(2,,(3,,(4,,(5,,(6,,(7,,(8,,9))))))))`},
	})
}
