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
	ListA := []int{1, 1, 2, 1}
	headA := createList(ListA)

	fmt.Print(isPalindrome(headA))

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
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	head_p := reverseList(slow)
	p1, p2 := head, head_p
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
