package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSortQuick(t *testing.T) {
	as := assert.New(t)

	//as.Equal([]int{1, 2, 3, 4, 5}, QuickSort([]int{5, 1, 2, 3, 4}))
	as.Equal([]int{1, 2, 3, 4, 5, 7}, QuickSort([]int{5, 1, 7, 2, 3, 4}))
}
