package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_4(t *testing.T) {
	test.Runs(t, findMedianSortedArrays, []*test.Case{
		{Input: `[1,3] \n [2]`, Output: "2"},
		{Input: `[3,4] \n [1,2,3]`, Output: "3"},
		{Input: `[1,2] \n [3,4]`, Output: "2.5"},
		{Input: `[2,3] \n []`, Output: "2.5"},
		{Input: `[2] \n []`, Output: "2"},
		{Input: `[] \n [2,3]`, Output: "2.5"},
		{Input: `[] \n [2]`, Output: "2"},
	})
}
