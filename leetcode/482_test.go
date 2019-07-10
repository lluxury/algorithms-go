package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_482(t *testing.T) {
	test.Runs(t, licenseKeyFormatting, []*test.Case{
		{Input: `5F3Z-2e-9-w \n 4`, Output: `5F3Z-2E9W`},
		{Input: `2-5g-3-J \n 2`, Output: `2-5G-3J`},
	})
}
