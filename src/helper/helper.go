package helper

import (
	"sort"

	"github.com/andreymudri/go_test/src/types"
)


func CreateNodeList(list []int) *types.ListNode {
	var head *types.ListNode
	for i := len(list) - 1; i >= 0; i-- {
		head = &types.ListNode{Val: list[i], Next: head}
	}
	return head
}

func ListNodetoSortedArray(merged *types.ListNode) []int {
	var mergedList []int
	for list := merged; list != nil; list = list.Next {
			mergedList = append(mergedList, list.Val)
	}
	// sort integer list
	sort.Ints(mergedList)
	return mergedList
}