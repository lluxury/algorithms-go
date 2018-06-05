package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Rand(t *testing.T) {
	as := assert.New(t)
	s := RandSlice(10)
	as.Len(s, 10)
}
