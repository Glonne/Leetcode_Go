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
	ListB := []int{5, 6, 1, 8, 4, 5}
	headA := createList(ListA)
	headB := createList(ListB)
	intersectVal := getIntersectionNode(headA, headB)
	if intersectVal != nil {
		fmt.Printf("%d", intersectVal.Val)
	} else {
		fmt.Println("两个链表无相交节点")
	}
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
