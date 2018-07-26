package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

//Input: root = [4,2,6,1,3,null,null]
//Output: 1
//Explanation:
//Note that root is a TreeNode object, not an array.
//
//The given tree [4,2,6,1,3,null,null] is represented by the following diagram:
//
//          4
//        /   \
//      2      6
//     / \
//    1   3
func Test_783(t *testing.T) {
	test.Runs(t, minDiffInBST, []*test.Case{
		{Input: `(4, (2,1,3), 6)`, Output: `1`},
	})
}
