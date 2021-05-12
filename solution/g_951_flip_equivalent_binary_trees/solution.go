package g_951_flip_equivalent_binary_trees

import (
	. "github.com/rsj217/leetcode-in-golang/datastruct"
)

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	return dfs(root1, root2)
}

func dfs(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	return (dfs(node1.Left, node2.Left) && dfs(node1.Right, node2.Right)) || (dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left))
}
