package leetcode

import (
	"github.com/lluxury/algorithms-go/test"
	"testing"
)

func Test_720(t *testing.T) {
	test.Runs(t, longestWord, []*test.Case{
		{Input: `[k,lg,it,oidd,oid,oiddm,kfk,y,mw,kf,l,o,mwaqz,oi,ych,m,mwa]`, Output: `oiddm`},
		{Input: `[w, wo, wor, worl, world]`, Output: `world`},
		{Input: `[a, banana, app, appl, ap, apply, apple]`, Output: `apple`},
	})
}
