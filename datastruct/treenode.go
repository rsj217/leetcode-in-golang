package datastruct

import (
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func (n *TreeNode) getHeight(node *TreeNode) int {
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

func (n *TreeNode) Height() int {
	return n.getHeight(n)
}

func Create(nums []interface{}) *TreeNode {
	size := len(nums)
	if size <= 0 {
		return nil
	}

	val := nums[0].(int)
	root := NewTreeNode(val)
	queue := make([]*TreeNode, 0, len(nums))
	queue = append(queue, root)
	index := 0
	for 0 < len(queue) {
		tempQueue := make([]*TreeNode, 0, len(nums))
		for i, _ := range queue {
			index++
			if 2*index-1 < size && nums[2*index-1] != nil {
				queue[i].Left = NewTreeNode(nums[2*index-1].(int))
				tempQueue = append(tempQueue, queue[i].Left)
			}
			if 2*index < size && nums[2*index] != nil {
				queue[i].Right = NewTreeNode(nums[2*index].(int))
				tempQueue = append(tempQueue, queue[i].Right)
			}
		}
		queue = tempQueue
	}
	return root
}

func PrintTree(node *TreeNode) string {
	if node == nil {
		return ""
	}
	levels := make([]string, 0, 10)
	ans := printTreeBFS(node)
	for _, line := range ans {
		levels = append(levels, strings.Join(line, ""))
	}
	return strings.Join(levels, "\n")
}

type NodeInfo struct {
	node *TreeNode
	seq  int
}

func printTreeBFS(node *TreeNode) [][]string {
	if node == nil {
		return [][]string{}
	}
	height := node.Height()
	width := (1 << height) - 1
	deep := 0
	queue := make([]*NodeInfo, 0, width)
	queue = append(queue, &NodeInfo{node, 0})

	levels := make([][]string, 0, height)
	for i := 0; i < height; i++ {
		level := make([]string, width, width)
		for j := 0; j < width; j++ {
			level[j] = " "
		}
		levels = append(levels, level)
	}

	for len(queue) > 0 {
		deep += 1

		levelQueue := make([]*NodeInfo, 0, width)

		levelHeight := height - deep + 1
		step := int(math.Pow(float64(2), float64(levelHeight)))
		leftSeq := int(math.Pow(float64(2), float64(deep-1))) - 1
		leftIndex := int(math.Pow(float64(2), float64(levelHeight-1))) - 1
		for _, v := range queue {
			curIndex := leftIndex + (v.seq-leftSeq)*step
			levels[deep-1][curIndex] = strconv.Itoa(v.node.Val)
			if v.node.Left != nil {
				levelQueue = append(levelQueue, &NodeInfo{v.node.Left, 2*v.seq + 1})
			}
			if v.node.Right != nil {
				levelQueue = append(levelQueue, &NodeInfo{v.node.Right, 2*v.seq + 2})
			}
		}
		queue = levelQueue
	}
	return levels
}

func printTreeDFS(node *TreeNode) [][]string {
	height := node.Height()
	width := (1 << height) - 1

	levels := make([][]string, 0, height)
	for i := 0; i < height; i++ {
		level := make([]string, width, width)
		for j := 0; j < width; j++ {
			level[j] = " "
		}
		levels = append(levels, level)
	}
	printTreeDfs(node, levels, 1, 0, width)
	return levels
}

func printTreeDfs(node *TreeNode, levels [][]string, deep, lo, hi int) {
	if node == nil {
		return
	}
	mid := lo + (hi-lo)/2
	levels[deep-1][mid] = strconv.Itoa(node.Val)

	printTreeDfs(node.Left, levels, deep+1, lo, mid)
	printTreeDfs(node.Right, levels, deep+1, mid+1, hi)
}
