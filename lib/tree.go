package lib

import (
	"bytes"
	"fmt"
	"regexp"
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

var treeReg = regexp.MustCompile(`(\d*),(\(\d*,\d*,\d*\)|\d*),(\(\d*,\d*,\d*\)|\d*)`)

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
	if data == "" {
		return nil, nil
	}

	startWithLeftBrackets := strings.HasPrefix(data, `(`)
	startWithRightBrackets := strings.HasSuffix(data, `)`)
	if startWithLeftBrackets && startWithRightBrackets {
		// ()
		data = strings.TrimSuffix(strings.TrimPrefix(data, `(`), `)`)

		if data == "" {
			return nil, nil
		}

		var tree = new(TreeNode)
		var err error

		// ,数量
		commaCount := strings.Count(data, `,`)

		// 0 ''
		// 1 1,x
		// 2 1,x,y
		if commaCount < 3 {
			var datas = strings.Split(data, `,`)
			var left, right string
			if commaCount > 0 {
				left = datas[1]
			}
			if commaCount == 2 {
				right = datas[2]
			}
			tree, err = treeNodeUnmarshalAll(datas[0], left, right)
		} else {
			result := treeReg.FindStringSubmatch(data)
			if len(result) != 4 {
				return nil, fmt.Errorf("result(%v) len(%d) is invalid", result, len(result))
			}
			tree, err = treeNodeUnmarshalAll(result[1], result[2], result[3])
		}
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
