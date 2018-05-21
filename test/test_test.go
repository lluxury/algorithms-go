package test

import (
	"testing"

	"github.com/Chyroc/algorithms-go/lib"
	"strings"
)

func TestRunCase(t *testing.T) {
	t.Run("int-float-string-bool", func(t *testing.T) {
		Runs(t, func(a int) int { return a + 1 }, []*Case{
			{Input: `1`, Output: "2"},
		})
		Runs(t, func(a int) float64 { return float64(a) / 2 }, []*Case{
			{Input: `1`, Output: "0.5"},
		})
		Runs(t, func(a string) string { return a + "1" }, []*Case{
			{Input: "a", Output: "a1"},
		})
		Runs(t, func(i int) bool { return i < 10 }, []*Case{
			{Input: "1", Output: "true"},
			{Input: "10", Output: "false"},
		})
	})

	t.Run("multi params", func(t *testing.T) {
		Runs(t, func(a, b int) int { return a + b }, []*Case{
			{Input: `1\n2`, Output: "3"},
		})
		Runs(t, func(a, b string) string { return a + b }, []*Case{
			{Input: `1\n2`, Output: "12"},
		})
		Runs(t, func(a string) []string { return strings.Split(a, ",") }, []*Case{
			{Input: `12,xx`, Output: "[12,xx]"},
		})
	})

	t.Run("array", func(t *testing.T) {
		Runs(t, func(a []int) int { return len(a) }, []*Case{
			{Input: `[]`, Output: `0`},
		})
		Runs(t, func(s []int) bool { return len(s) > 2 }, []*Case{
			{Input: `[1,2,3]`, Output: "true"},
			{Input: `[1]`, Output: "false"},
		})
	})

	t.Run("external type", func(t *testing.T) {
		Runs(t, func(node *lib.ListNode) *lib.ListNode { return node.Clone() }, []*Case{
			{Input: `1 -> 2 -> 3`, Output: "1 -> 2->3"},
			{Input: `1`, Output: "1"},
		})
	})
}
