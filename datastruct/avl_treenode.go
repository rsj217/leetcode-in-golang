package datastruct

import "math"

type AVLTreeNode struct {
	Key    int
	Height int
	Left   *AVLTreeNode
	Right  *AVLTreeNode
}

func NewAVLTreeNode(key int) *AVLTreeNode {
	return &AVLTreeNode{Key: key, Height: 1}
}

func (n AVLTreeNode) MinimumDFS(node *AVLTreeNode) *AVLTreeNode {
	if node.Left == nil {
		return node
	}

	return n.MinimumDFS(node.Left)
}

func (n *AVLTreeNode) InsertDFS(node *AVLTreeNode, key int) *AVLTreeNode {
	if node == nil {
		return NewAVLTreeNode(key)
	}
	if key < node.Key {
		node.Left = n.InsertDFS(node.Left, key)
	} else if node.Key < key {
		node.Right = n.InsertDFS(node.Right, key)
	} else {
		node.Key = key
	}
	updateHeight(node)
	return balanceRotate(node)
}

func (n *AVLTreeNode) DeleteDFS(node *AVLTreeNode, key int) *AVLTreeNode {
	if node == nil {
		panic("node not exist")
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
			node = successor
		} else if node.Left != nil {
			node = node.Left
		} else if node.Right != nil {
			node = node.Right
		} else {
			node = nil
		}
	}
	updateHeight(node)
	return balanceRotate(node)
}

func updateHeight(node *AVLTreeNode) *AVLTreeNode {
	node.Height = 1 + int(math.Max(float64(getHeight(node.Left)), float64(getHeight(node.Right))))
	return node
}

func getHeight(node *AVLTreeNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func getBF(node *AVLTreeNode) int {
	if node == nil {
		return 0
	}
	return getHeight(node.Left) - getHeight(node.Right)
}

func rightRotate(node *AVLTreeNode) *AVLTreeNode {
	x := node
	y := x.Left
	T3 := y.Right
	y.Right = x
	x.Left = T3

	updateHeight(x)
	updateHeight(y)
	return y
}

func leftRotate(node *AVLTreeNode) *AVLTreeNode {
	x := node
	y := node.Right
	T2 := y.Left
	y.Left = x
	x.Right = T2

	updateHeight(x)
	updateHeight(y)
	return y
}

func balanceRotate(node *AVLTreeNode) *AVLTreeNode {
	bf := getBF(node)
	if 1 < bf {
		if 0 <= getBF(node.Left) {
			return rightRotate(node)
		} else {
			node.Left = leftRotate(node.Left)
			return rightRotate(node)
		}
	} else if bf < -1 {
		if getBF(node.Right) <= 0 {
			return leftRotate(node)
		} else {
			node.Right = rightRotate(node.Right)
			return leftRotate(node)
		}
	}
	return node
}

type AVLTree struct {
	root *AVLTreeNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (avl *AVLTree) Insert(key int) {
	if avl.root == nil {
		avl.root = NewAVLTreeNode(key)
	} else {
		avl.root = avl.root.InsertDFS(avl.root, key)
	}
}

func (avl *AVLTree) Delete(key int) {
	if avl.root != nil {
		avl.root = avl.root.DeleteDFS(avl.root, key)
	}
}
