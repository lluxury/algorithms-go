package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SplitWithToken(t *testing.T) {
	t.Run("", func(t *testing.T) {
		as := assert.New(t)
		var cases = []string{`[1]`, `[[1]]`, `[1,23]`, `[1,2,3,[1,2,3],2,3,[3,4]]`, `[1,2,3,[1,2,3],2,3,[3,4],2]`, `[ 1, [ 1, [ 1, [1] ] ] ]`}
		var results = [][]string{
			{"1"},
			{"[1]"},
			{"1", "23"},
			{"1", "2", "3", "[1,2,3]", "2", "3", "[3,4]"},
			{"1", "2", "3", "[1,2,3]", "2", "3", "[3,4]", "2"},
			{"1", "[1,[1,[1]]]"},
		}
		for k := range cases {
			x, err := SplitWithToken(cases[k], '[', ']')
			as.Nil(err)
			as.Equal(results[k], x)
		}
	})

	t.Run("", func(t *testing.T) {
		as := assert.New(t)
		var cases = []string{`(2,2,3)`, `(2,(2,1,),3)`, `(2,(2,,1),3)`, `(2,(2,2,1),3)`, `(2,3,(2,1,))`, `(5,(3,2,4),(6,,7))`, `(1,(2,,3),(2,(1,2,3),3))`}
		var results = [][]string{
			{"2", "2", "3"},
			{"2", "(2,1,)", "3"},
			{"2", "(2,,1)", "3"},
			{"2", "(2,2,1)", "3"},
			{"2", "3", "(2,1,)"},
			{"5", "(3,2,4)", "(6,,7)"},
			{"1", "(2,,3)", "(2,(1,2,3),3)"},
		}
		for k := range cases {
			x, err := SplitWithToken(cases[k], '(', ')')
			as.Nil(err)
			as.Equal(results[k], x)
			as.Len(x, 3)
		}
	})
}
