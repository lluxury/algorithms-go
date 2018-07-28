package leetcode

import (
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

func Test_6(t *testing.T) {
	test.Runs(t, convert, []*test.Case{
		{Input: `PAYPALISHIRING \n 3`, Output: `PAHNAPLSIIGYIR`},
		{Input: `PAYPALISHIRING \n 4`, Output: `PINALSIGYAHRPI`},
		{Input: `AB \n 1`, Output: `AB`},
	})
}
