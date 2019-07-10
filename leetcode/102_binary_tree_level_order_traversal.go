package leetcode

import (
	"github.com/lluxury/algorithms-go/lib"
 )

/*
 Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
*/

// 给定一个二叉树，逐层返回遍历的节点值
// root加入队列,得出队列长度
// 在目前长度内,获取指针
// 添加指针值Val,判断是否有左右子项
// 如果有,更新队列,填到最后
// 循环完lenQ后,跳过已循环部分,重定义q
// ret添加多个结果


func levelOrder(root *lib.TreeNode) [][]int  {
    return bfs(root)
}

func bfs(root *lib.TreeNode) [][]int  {
    if root == nil {
        return nil
    }

    var ret [][]int
    var q []*lib.TreeNode
    q = append(q, root)
    for len(q) > 0 {
        lenQ := len(q)
        var tmp []int
        for i := 0; i  < lenQ; i++ {
            n := q[i]
            tmp = append(tmp,n.Val)
            if n.Left != nil{
                q = append(q, n.Left)
            }

            if n.Right != nil {
                q = append(q, n.Right)
			}
			// 前序遍历
		}
		// print("lenQ=",len(q), " ")
		q = q[lenQ:]
        ret = append(ret, tmp)
    }
    return ret
}

