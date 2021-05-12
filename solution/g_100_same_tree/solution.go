package g_100_same_tree

import (
	. "github.com/rsj217/leetcode-in-golang/datastruct"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else { // (p != nil && q != nil)
		if p.Val != q.Val {
			return false
		}
		return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

	}
}
