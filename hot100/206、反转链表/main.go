package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(list []int) *ListNode {
	head := &ListNode{Val: list[0]}
	var Node_current *ListNode
	Node_current = head
	for i := 1; i < len(list); i++ {
		newhead := &ListNode{Val: list[i]}
		Node_current.Next = newhead
		Node_current = newhead
	}
	return head
}
func main() {
	ListA := []int{4, 1, 8, 4, 5}
	headA := createList(ListA)
	headA_r := reverseList(headA)
	for tmp := headA_r; tmp != nil; tmp = tmp.Next {
		fmt.Printf("The reverse list is %d\n", tmp.Val)
	}
}
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
