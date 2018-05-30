package lib

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) Add(v int) {
	if node == nil {
		*node = ListNode{}
	}
	tmp := node
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &ListNode{Val: v}
}

func (node *ListNode) Clone() *ListNode {
	if node == nil {
		return nil
	}
	var l = &ListNode{Val: node.Val}
	var tmpL = l
	var tmpNode = node

	for tmpNode.Next != nil {
		tmpNode = tmpNode.Next
		tmpL.Next = &ListNode{Val: tmpNode.Val}
		tmpL = tmpL.Next
	}

	return l
}

func EqualsToListNode(a, b *ListNode) bool {
	if a == nil || b == nil {
		return b == b
	}

	return a.Val == b.Val && EqualsToListNode(a.Next, b.Next)
}

// 定义字符串能够通过转成定的类型: string -> typ
// 特定的类型也能转成字符串:      typ -> string
func (node *ListNode) Unmarshal(param string) (interface{}, error) {
	var list = new(ListNode)
	for k, v := range strings.Split(param, "->") {

		v = strings.TrimSpace(v)

		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		if k == 0 {
			// 差点出bug? 真的？
			//*list = ListNode{Val: i}
			list = &ListNode{Val: i}
		} else {
			list.Add(i)
		}
	}
	return list, nil
}

func (node *ListNode) Marshal() (string, error) {
	var s []string
	var tmp = node

	if node != nil {
		s = append(s, strconv.Itoa(node.Val))
	}

	for tmp.Next != nil {
		tmp = tmp.Next
		s = append(s, strconv.Itoa(tmp.Val))
	}
	return strings.Join(s, "->"), nil
}

func (node *ListNode) String() string {
	s, _ := node.Marshal()
	return s
}
