package g_94_binary_tree_inorder_traversal

import (
	. "github.com/rsj217/leetcode-in-golang/datastruct"
)

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0, 1)
	inorderTraversalIter(root, &ans)
	return ans
}

func inorderTraversalIter(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}
	inorderTraversalIter(node.Left, ans)
	*ans = append(*ans, node.Val)
	inorderTraversalIter(node.Right, ans)
	return
}

func inorderTraversalRecursion(root *TreeNode) []int {
	ans := make([]int, 0, 10)
	stack := make([]*TreeNode, 0, 10)

	node := root
	for {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		if len(stack) == 0 {
			break
		}

		node = stack[len(stack)-1]
		ans = append(ans, node.Val)
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return ans
}
