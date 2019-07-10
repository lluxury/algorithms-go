package leetcode

import (
	"testing"

	"github.com/lluxury/algorithms-go/test"
)

func Test_917(t *testing.T) {
	test.Runs(t, reverseOnlyLetters, []*test.Case{
		{Input: `ab-cd`, Output: `dc-ba`},
		{Input: `a-bC-dEf-ghIj`, Output: `j-Ih-gfE-dCba`},
		{Input: `Test1ng-Leet=code-Q!`, Output: `Qedo1ct-eeLg=ntse-T!`},
	})
}
