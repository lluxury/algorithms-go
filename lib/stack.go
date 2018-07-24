package lib

import (
	"fmt"
	"strings"
)

type Stack struct {
	is []interface{}
}

func (s *Stack) String() string {
	var ss []string
	for !s.IsEmpty() {
		ss = append(ss, fmt.Sprintf("%v", s.Pop()))
	}

	return strings.Join(ss, " | ")
}

func (s *Stack) Push(i interface{}) {
	s.is = append(s.is, i)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic("is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		panic("is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *Stack) Size() int {
	return len(s.is)
}

func (s *Stack) IsEmpty() bool {
	return len(s.is) == 0
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Clone() *Stack {
	s2 := NewStack()
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
