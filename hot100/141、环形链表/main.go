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
	Node_current.Next = head
	return head
}
func main() {
	ListA := []int{4, 1, 8, 4, 5}
	headA := createList(ListA)
	fmt.Print(hasCycle(headA))
}
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
