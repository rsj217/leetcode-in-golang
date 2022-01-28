package datastruct

import "math"

type BSTreeNode struct {
	Key   int
	Left  *BSTreeNode
	Right *BSTreeNode
}

func NewBSTreeNode(key int) *BSTreeNode {
	return &BSTreeNode{Key: key}
}

func (n *BSTreeNode) getHeight(node *BSTreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := node.getHeight(node.Left)
	rightHeight := node.getHeight(node.Right)
	if leftHeight < rightHeight {
		return 1 + rightHeight
	}
	return 1 + leftHeight
}

func (n *BSTreeNode) Height() int {
	return n.getHeight(n)
}

func (n *BSTreeNode) MinimumDFS(node *BSTreeNode) *BSTreeNode {
	if node.Left == nil {
		return node
	}
	return n.MinimumDFS(node.Left)
}

func (n *BSTreeNode) InsertDFS(node *BSTreeNode, key int) *BSTreeNode {
	if node == nil {
		return NewBSTreeNode(key)
	}
	if key < node.Key {
		node.Left = n.InsertDFS(node.Left, key)
	} else if node.Key < key {
		node.Right = n.InsertDFS(node.Right, key)
	} else {
		node.Key = key
	}
	return node
}

func (n *BSTreeNode) DeleteDFS(node *BSTreeNode, key int) *BSTreeNode {
	if node == nil {
		panic("node not exit")
	}
	if key < node.Key {
		node.Left = n.DeleteDFS(node.Left, key)
	} else if node.Key < key {
		node.Right = n.DeleteDFS(node.Right, key)
	} else {
		if node.Left != nil && node.Right != nil {
			successor := n.MinimumDFS(node.Right)
			successor.Right = n.DeleteDFS(node.Right, successor.Key)
			successor.Left = node.Left
			return successor
		} else if node.Left != nil {
			return node.Left
		} else if node.Right != nil {
			return node.Right
		} else {
			return nil
		}
	}
	return node
}

func IsBST(node *BSTreeNode) bool {
	var dfs func(node *BSTreeNode, l, r int) bool
	dfs = func(node *BSTreeNode, l, r int) bool {
		if node == nil {
			return true
		}
		if l < node.Key && node.Key < r {
			return dfs(node.Left, l, node.Key) && dfs(node.Right, node.Key, r)
		}
		return false
	}
	return dfs(node, math.MinInt64, math.MaxInt64)
}

type BSTree struct {
	root *BSTreeNode
}

func NewBSTree() *BSTree {
	return &BSTree{}
}

func (bst *BSTree) Insert(key int) {
	if bst.root == nil {
		bst.root = NewBSTreeNode(key)
	} else {
		bst.root = bst.root.InsertDFS(bst.root, key)
	}
}

func (bst *BSTree) Delete(key int) {
	if bst.root != nil {
		bst.root = bst.root.DeleteDFS(bst.root, key)
	}
}
