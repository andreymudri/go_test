package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func createNodeList(list []int) *ListNode {
	var head *ListNode
	for i := len(list) - 1; i >= 0; i-- {
		head = &ListNode{Val: list[i], Next: head}
	}
	return head
}