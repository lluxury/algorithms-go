package leetcode

import "container/list"
import "errors"

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]*/


// 生成有效括号


func generateParenthesis(n int) []string {
    if n <= 0 {
        return nil
    }

    var combi []string
    genarateAllCombination(2*n, "", &combi)

    var ans []string
    for i := 0; i < len(combi); i++ {
        if isValid(combi[i]){
            ans = append(ans, combi[i])
        }
    }

    return ans
}

func isValid(s string) bool  {
    q := newQueue()
    valid := true

    for _, r := range s {
        switch r {
        case '(':
            q.enqueue(r)
        case ')':
            v, err := q.dequeue()
            if err != nil || v != '('{
                valid = false
                break
            }
        }
    }
    return valid && q.isEmpty()
}

func genarateAllCombination(n int, prefix string, res *[]string)  {
    if n == 0 {
        *res = append(*res, prefix)
        return
    }

    genarateAllCombination(n-1, prefix+"(",res)
    genarateAllCombination(n-1, prefix+")",res)
}

type queue struct {
    l *list.List
}

func newQueue() *queue  {
    return &queue{l: list.New()}
}

func (q *queue) enqueue(v interface{})  {
    q.l.PushBack(v)
}

func (q *queue) isEmpty() bool  {
    if q.l.Len() == 0{
        return true
    }
    return false
}

func (q *queue) dequeue() (interface{}, error)  {
    if q.isEmpty(){
        return nil, errors.New("dequeue failed cause of queue is empty")
    }
    return q.l.Remove(q.l.Back()), nil
}
