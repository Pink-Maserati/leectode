package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	count := len(lists) - 1
	fmt.Println(count)
	return merge(lists, 0, count)
}

func merge(Lists []*ListNode, left int, right int) *ListNode {
	if left >= right {
		return Lists[left]
	}
	mid := (left + right) / 2
	return mergeTwoLists(merge(Lists, left, mid), merge(Lists, mid+1, right))
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummy := &ListNode{0, nil}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}
