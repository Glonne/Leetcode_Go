package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Your task: implement this function
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root} // 初始化队列，放入根节点
	depth := 0                 // 记录深度

	for len(queue) > 0 {
		// 当前层的节点数量（关键：用于区分不同层）
		currentLevelSize := len(queue)

		// 遍历当前层的所有节点
		for i := 0; i < currentLevelSize; i++ {
			// 出队（队列是先进先出，移除第一个元素）
			node := queue[0]
			queue = queue[1:]

			// 左孩子不为空则入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// 右孩子不为空则入队
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 每处理完一层，深度+1
		depth++
	}

	return depth
}

// Helper function to build a sample binary tree:
//
//	    1
//	   / \
//	  2   3
//	 /
//	4
func buildSampleTree() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
}

func main() {
	root := buildSampleTree()
	depth := maxDepth(root)
	fmt.Println("二叉树的最大深度是:", depth)
}
