package leetcode

import "github.com/Chyroc/algorithms-go/lib"

/*
 * [144] Binary Tree Preorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (48.02%)
 * Total Accepted:    253.8K
 * Total Submissions: 527.8K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the preorder traversal of its nodes' values.
 * 
 * Example:
 * 
 * 
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 * 
 * Output: [1,2,3]
 * 
 * 
 * Follow up: Recursive solution is trivial, could you do it iteratively?
 * 
 */


/*

* 求一个二叉树的先序遍历的值的数组
  * 方法1：递归，很简单
  * 方法2：《二叉树的非递归遍历》

《二叉树的非递归遍历》
* 都是通过模拟栈实现的

*/

func preorderTraversal_Recursive_1(root *lib.TreeNode) []int {
	var values []int
	Recursive_preorderTraversal(root, &values)
	return values
}

func Recursive_preorderTraversal(root *lib.TreeNode, values *[]int) {
	if root == nil {
		return
	}

	*values = append(*values, root.Val)
	Recursive_preorderTraversal(root.Left, values)
	Recursive_preorderTraversal(root.Right, values)
}

//（这个前序遍历也就是答案了）
// 前序遍历二叉树
// 需要通过模拟栈实现
// for for root != nil || len(stack) != 0 { } 二叉树不为空||栈不为空
//   for root != nil { } 处理树当前节点数据，然后入栈，然后指向左孩，循环
//   if len(stack) != 0 { } 前面那个循环结束之后就需要出栈，处理右孩
func preOrder_Iterative(root *lib.TreeNode) []int {
	var stack []lib.TreeNode
	var values []int

	for root != nil || len(stack) != 0 {
		for root != nil {
			values = append(values, root.Val) // 处理
			stack = append(stack, *root)
			root = root.Left
		}

		if len(stack) != 0 {
			root = &stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}

	return values
}

// 中序遍历二叉树
// 和前序遍历一样，只不过处理函数的位置变化了
func inOrder_Iterative(root *lib.TreeNode) []int {
	var stack []lib.TreeNode
	var values []int

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, *root)
			root = root.Left
		}

		if len(stack) != 0 {
			root = &stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			values = append(values, root.Val) // 处理

			root = root.Right
		}
	}

	return values
}

// 后序遍历二叉树
// tip：
//   需要先处理左孩，再处理右孩，再处理根节点，所以根节点、右孩、左孩依次入队
//   当 左孩和右孩都是空（叶子节点） / 已经处理完左孩和右孩 的时候，可以处理根节点
// 步骤：
//   根节入队
//   然后循环：
//     取栈顶节点
//     如果左孩或者右孩为空，进行处理，然后栈顶数据出栈 + 将当前节点设置为pre节点，表示上一次处理的节点
//     如果pre节点不为空，并且pre节点等于当前节点的左孩或者右孩 =》  根据入栈顺序「根节点、右孩、左孩」 =》 当前节点是跟节点，并且这个节点已经处理了左孩和右孩 =》 处理本节点
//     上面两个条件可以合并
//     else：将不为空的右孩入栈，不为空的左孩入栈
func postOrder_Iterative(root *lib.TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack []lib.TreeNode
	var cur *lib.TreeNode
	var pre *lib.TreeNode
	var values []int

	stack = append(stack, *root)
	for len(stack) != 0 {
		cur = &stack[len(stack)-1]

		if (cur.Left == nil && cur.Right == nil) || (pre != nil && (lib.EqualsToTreeNode(pre, cur.Left) || lib.EqualsToTreeNode(pre, cur.Right))) {
			values = append(values, cur.Val) // 处理
			stack = stack[:len(stack)-1]
			pre = cur
		} else {
			if cur.Right != nil {
				stack = append(stack, *cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, *cur.Left)
			}
		}
	}

	return values
}
