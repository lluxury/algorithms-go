package test

import (
	"testing"

	"github.com/Chyroc/algorithms-go/lib"
)

func TestRunCase(t *testing.T) {
	t.Run("init", func(t *testing.T) {
		Runs(t, func(a []int) int { return len(a) }, []*Case{
			{Input: `[]`, Output: `0`},
		})
	})

	t.Run("ListNode", func(t *testing.T) {
		Runs(t, func(node *lib.ListNode) *lib.ListNode { return node.Clone() }, []*Case{
			{Input: `1 -> 2 -> 3`, Output: "1 -> 2->3"},
			{Input: `1`, Output: "1"},
		})
	})

	t.Run("(int)->int", func(t *testing.T) {
		Runs(t, func(a int) int { return a + 1 }, []*Case{
			{Input: `1`, Output: "2"},
			{Input: `999`, Output: "1000"},
		})
	})

	t.Run("(int,int)->int", func(t *testing.T) {
		Runs(t, func(a, b int) int { return a + b }, []*Case{
			{Input: `1\n2`, Output: "3"},
		})
	})

	t.Run("(string)->string", func(t *testing.T) {
		Runs(t, func(a string) string { return a + "1" }, []*Case{
			{Input: "a", Output: "a1"},
		})
	})

	t.Run("(int)->bool", func(t *testing.T) {
		Runs(t, func(i int) bool { return i < 10 }, []*Case{
			{Input: "1", Output: "true"},
			{Input: "10", Output: "false"},
		})
	})

	t.Run("(string,string)->string", func(t *testing.T) {
		Runs(t, func(a, b string) string { return a + b }, []*Case{
			{Input: `a\nb`, Output: "ab"},
		})
	})

	t.Run("(list<int>)->bool", func(t *testing.T) {
		Runs(t, func(s []int) bool { return len(s) > 2 }, []*Case{
			{Input: `[1,2,3]`, Output: "true"},
			{Input: `[1]`, Output: "false"},
		})
	})
}
