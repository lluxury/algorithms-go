package lib

import (
	"strconv"
	"strings"
)

type Stack struct {
	is []int
}

func (s *Stack) String() string {
	var ss []string
	for !s.IsEmpty() {
		ss = append(ss, strconv.Itoa(s.Pop()))
	}

	return strings.Join(ss, " | ")
}

func (s *Stack) Push(i int) {
	s.is = append(s.is, i)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *Stack) Peek() int {
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
