package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListNode(t *testing.T) {
	t.Run("clone", func(t *testing.T) {
		as := assert.New(t)
		l := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		}
		b := l.Clone()

		as.Equal(1, l.Val)
		as.Equal(1, b.Val)

		l.Val = 2
		as.Equal(2, l.Val)
		as.Equal(1, b.Val)

		l.Next = nil
		as.Nil(l.Next)
		as.NotNil(b.Next)
	})

	t.Run("add node", func(t *testing.T) {
		as := assert.New(t)

		node := &ListNode{Val: 1}
		node.Add(2)
		node.Add(3)

		as.Equal(1, node.Val)
		as.Equal(2, node.Next.Val)
		as.Equal(3, node.Next.Next.Val)
	})

	t.Run("marshal", func(t *testing.T) {
		as := assert.New(t)

		node2, err := new(ListNode).Unmarshal("1->2->3")
		as.Nil(err)
		node, ok := node2.(*ListNode)
		as.True(ok)

		as.Equal(1, node.Val)
		as.Equal(2, node.Next.Val)
		as.Equal(3, node.Next.Next.Val)

		s, err := node.Marshal()
		as.Nil(err)
		as.Equal("1->2->3", s)
	})
}
