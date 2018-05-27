package lib

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	s, _ := t.Marshal()
	return s
}

func EqualsTo(r, t *TreeNode) bool {
	if t == nil || r == nil {
		return r == t
	}

	if t.Val != r.Val {
		return false
	}

	return EqualsTo(t.Left, r.Left) && EqualsTo(t.Right, r.Right)
}

func treeNodeUnmarshal(data string) (*TreeNode, error) {
	node, err := new(TreeNode).Unmarshal(data)
	if err != nil {
		return nil, err
	} else if node == nil {
		return nil, nil
	}
	tree, ok := node.(*TreeNode)
	if !ok {
		return nil, fmt.Errorf("convert interface to *TreeNode err")
	}
	return tree, nil
}

func treeNodeUnmarshalAll(data, left, right string) (*TreeNode, error) {
	if data == "" || data == "()" {
		return nil, nil
	}
	var tree = new(TreeNode)
	var err error

	tree.Val, err = strconv.Atoi(data)
	if err != nil {
		return nil, err
	}

	tree.Left, err = treeNodeUnmarshal(left)
	if err != nil {
		return nil, err
	}

	tree.Right, err = treeNodeUnmarshal(right)
	if err != nil {
		return nil, err
	}

	return tree, nil
}

func (t *TreeNode) Unmarshal(data string) (interface{}, error) {
	if data == "" || data == "()" {
		return nil, nil
	}

	startWithLeftBrackets := strings.HasPrefix(data, `(`)
	startWithRightBrackets := strings.HasSuffix(data, `)`)
	if startWithLeftBrackets && startWithRightBrackets {
		results, err := SplitWithToken(data, '(', ')')
		if err != nil {
			return nil, err
		}

		var tree = new(TreeNode)

		for len(results) < 3 {
			results = append(results, "")
		}

		tree, err = treeNodeUnmarshalAll(results[0], results[1], results[2])

		if err != nil {
			return nil, err
		}
		return tree, err
	} else if !startWithLeftBrackets && !startWithRightBrackets {
		// int
		i, err := strconv.Atoi(data)
		if err != nil {
			return nil, err
		}
		return &TreeNode{Val: i}, nil
	} else {
		return nil, fmt.Errorf("must () or int")
	}
}

func (t *TreeNode) Marshal() (string, error) {
	if t == nil {
		return ``, nil
	}

	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val), nil
	} else {
		var buf bytes.Buffer

		buf.WriteString(fmt.Sprintf("(%d,", t.Val))

		l, err := t.Left.Marshal()
		if err != nil {
			return "", err
		}
		buf.WriteString(l)
		buf.WriteString(`,`)

		r, err := t.Right.Marshal()
		if err != nil {
			return "", err
		}
		buf.WriteString(r)
		buf.WriteString(`)`)
		return buf.String(), nil
	}
}

func (t *TreeNode) Dump() {
	t.dump("")
}

func (t *TreeNode) dump(indent string) {
	if t == nil {
		//fmt.Printf(indent+"<nil>\n")
		return
	}

	fmt.Printf(indent+"%d\n", t.Val)
	t.Left.dump(indent + "  ")
	t.Right.dump(indent + "  ")
}
