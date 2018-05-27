package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Interval(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		{
			x, err := (&Interval{}).Marshal()
			as.Nil(err)
			as.Equal(`[0,0]`, x)
		}
		{
			x, err := (&Interval{1, 2}).Marshal()
			as.Nil(err)
			as.Equal(`[1,2]`, x)
		}
	})

	t.Run("", func(t *testing.T) {
		x, err := new(Interval).Unmarshal(`[1,2]`)
		as.Nil(err)
		m, ok := x.(*Interval)
		as.True(ok)
		as.Equal(&Interval{Start: 1, End: 2}, m)
	})
}
