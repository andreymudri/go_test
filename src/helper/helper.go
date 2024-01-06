package helper

import "github.com/andreymudri/go_test/types"


func CreateNodeList(list []int) *types.ListNode {
	var head *types.ListNode
	for i := len(list) - 1; i >= 0; i-- {
		head = &types.ListNode{Val: list[i], Next: head}
	}
	return head
}
