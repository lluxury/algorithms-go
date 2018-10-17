package leetcode

import (
	"github.com/Chyroc/algorithms-go/test"
	"testing"
)

func Test_599(t *testing.T) {
	test.Runs(t, findRestaurant, []*test.Case{
		{Input: `[Shogun, Tapioca Express, Burger King, KFC]\n[Piatti, The Grill at Torrey Pines, Hungry Hunter Steakhouse, Shogun]`, Output: `[Shogun]`},
	})
}
