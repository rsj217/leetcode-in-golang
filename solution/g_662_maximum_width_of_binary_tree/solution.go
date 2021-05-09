package g_662_maximum_width_of_binary_tree

import (
	. "github/rsj217/leetcode-in-golang/datastruct"
)

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0

	nodeQueue := make([]*TreeNode, 0, 20)
	seqQueue := make([]int, 0, 20)
	nodeQueue = append(nodeQueue, root)
	seqQueue = append(seqQueue, 0)

	for len(nodeQueue) > 0 {
		size := len(nodeQueue)
		start, end := 0, 0

		tmpNodeQueue := make([]*TreeNode, 0, 20)
		tmpSeqQueue := make([]int, 0, 20)

		for i := 0; i < size; i++ {
			node := nodeQueue[i]
			seq := seqQueue[i]
			if i == 0 {
				start = seq
			}
			if i == size-1 {
				end = seq
			}

			if node.Left != nil {
				tmpNodeQueue = append(tmpNodeQueue, node.Left)
				tmpSeqQueue = append(tmpSeqQueue, 2*seq+1)
			}
			if node.Right != nil {
				tmpNodeQueue = append(tmpNodeQueue, node.Right)
				tmpSeqQueue = append(tmpSeqQueue, 2*seq+2)
			}
		}
		nodeQueue = tmpNodeQueue
		seqQueue = tmpSeqQueue

		tmpAns := end - start + 1
		if ans < tmpAns {
			ans = tmpAns
		}
	}
	return ans
}
