package g_814_bianry_tree_pruning

import (
	. "github/rsj217/leetcode-in-golang/datastruct"
)

func pruneTree(root *TreeNode) *TreeNode {
	if dfs(root) {
		return nil
	}
	return root
}

func dfs(node *TreeNode) bool {
	if node == nil {
		return true
	}
	left := dfs(node.Left)
	right := dfs(node.Right)
	if left {
		node.Left = nil
	}
	if right {
		node.Right = nil
	}
	return left && right && node.Val == 0
}
