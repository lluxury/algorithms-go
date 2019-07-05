package leetcode

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func Test_460(t *testing.T) {
// 	as := assert.New(t)

// 	cache := Constructor_460(2)

// 	cache.Put(1, 1)
// 	cache.Put(2, 2)
// 	as.Equal(1, cache.Get(1))
// 	cache.Put(3, 3) // evicts key 2
// 	as.Equal(-1, cache.Get(2))
// 	as.Equal(3, cache.Get(3))
// 	cache.Put(4, 4) // evicts key 1.
// 	as.Equal(-1, cache.Get(1))
// 	as.Equal(3, cache.Get(3))
// 	as.Equal(4, cache.Get(4))
// }
