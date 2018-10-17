package leetcode

import "testing"

func Test_155(t *testing.T) {
	// TODO
	{
		m := Constructor_155()
		m.Push(-2)
		m.Push(0)
		m.Push(-3)

		if x := m.GetMin(); x != -3 {
			t.Fatal("get", x, "list", m.list, "min", m.min)
		}

		m.Pop()

		if x := m.Top(); x != 0 {
			t.Fatal("get", x)
		}

		if x := m.GetMin(); x != -2 {
			t.Fatal("get", x)
		}
	}
}
