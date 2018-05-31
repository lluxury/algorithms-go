package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeNode(t *testing.T) {
	as := assert.New(t)

	t.Run("TreeNode to string", func(t *testing.T) {
		{
			m, err := (&TreeNode{}).Marshal()
			as.Nil(err)
			as.Equal(`0`, m)
		}
		{
			m, err := (&TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}).Marshal()
			as.Nil(err)
			as.Equal(`(2,2,3)`, m)
		}
		{
			m, err := (&TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 3}}).Marshal()
			as.Nil(err)
			as.Equal(`(2,(2,1,),3)`, m)
		}
		{
			m, err := (&TreeNode{Val: 2, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 3}}).Marshal()
			as.Nil(err)
			as.Equal(`(2,(2,,1),3)`, m)
		}
		{
			m, err := (&TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 11}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 3}}).Marshal()
			as.Nil(err)
			as.Equal(`(2,(2,11,1),3)`, m)
		}
		{
			m, err := (&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3}},
			}).Marshal()
			as.Nil(err)
			as.Equal(`(1,(2,,3),(2,(1,2,3),3))`, m)
		}
	})

	t.Run("string to TreeNode", func(t *testing.T) {
		for _, v := range []string{``, `()`} {
			m, err := treeNodeUnmarshal(v)
			as.Nil(err)
			as.Nil(m)
		}

		for _, v := range []string{`0`, `(2,2,3)`, `(2,(2,1,),3)`, `(2,(2,,1),3)`, `(2,(2,2,1),3)`, `(2,3,(2,1,))`, `(5,(3,2,4),(6,,7))`, `(1,(2,,3),(2,(1,2,3),3))`} {
			m, err := treeNodeUnmarshal(v)
			as.Nil(err)
			as.NotNil(m)
			as.Equal(v, m.String())
		}
	})

	t.Run("equal", func(t *testing.T) {
		m1, err := treeNodeUnmarshal(`(2,3,(2,1,))`)
		as.Nil(err)
		as.NotNil(m1)

		m2, err := treeNodeUnmarshal(`(2,3,(2,1,2))`)
		as.Nil(err)
		as.NotNil(m2)

		m3, err := treeNodeUnmarshal(`(2,3,(2,1,))`)
		as.Nil(err)
		as.NotNil(m3)

		as.False(EqualsToTreeNode(m1, m2))
		as.True(EqualsToTreeNode(m1, m3))
	})
}
