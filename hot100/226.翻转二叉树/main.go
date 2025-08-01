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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// TODO: implement invertTree logic here
	stack := make([]*TreeNode, 0)
	curr := root
	stack = append(stack, curr)
	for len(stack) != 0 {
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr.Left, curr.Right = curr.Right, curr.Left
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}
	return root
}

// Helper function to build a sample binary tree:
//
//	     4
//	   /   \
//	  2     7
//	 / \   / \
//	1   3 6   9
func buildSampleTree() *TreeNode {
	return &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}
}

// Helper function to print the tree in level-order (BFS)
func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.Val, " ")
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}

func main() {
	root := buildSampleTree()
	fmt.Print("原始树（层序遍历）: ")
	printTree(root)

	inverted := invertTree(root)
	fmt.Print("翻转后树（层序遍历）: ")
	printTree(inverted)
}
