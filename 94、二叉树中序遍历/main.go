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
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}
func inorderTraversal1(root *TreeNode) []int { // 找到最左下角的节点，出栈最左下角的节点，然后看他有没有右孩子
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.Val)
		curr = curr.Right
	}
	return result
}

// Helper function to build a sample binary tree:
//
//	1
//	 \
//	  2
//	 /
//	3
func buildSampleTree() *TreeNode {
	return &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
}

func main() {
	root := buildSampleTree()
	result := inorderTraversal1(root)
	fmt.Println("中序遍历结果:", result)
}
