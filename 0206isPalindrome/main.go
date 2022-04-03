package main

import (
	"fmt"
)

// 02.06. 回文链表
// 编写一个函数，检查输入的链表是否是回文的。

// 示例 1：
// 输入： 1->2
// 输出： false

// 示例 2：
// 输入： 1->2->2->1
// 输出： true

// 进阶：
// 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	arr := []int{}
	for ; head != nil; head = head.Next {
		arr = append(arr, head.Val)
	}

	n := len(arr)
	for i, v := range arr[:n/2] {
		if v != arr[n-1-i] {
			return false
		}
	}

	return true

}

func isPalindrome1(head *ListNode) bool {
	if head == nil {
		return true
	}
	//找出后半部分
	half := endOfFirstHalf(head)
	fmt.Println(half)
	//反转后半部分的链表
	tail := reverseList(half.Next)
	//判断是否是回文
	p1 := head
	p2 := tail
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	//还原链表并返回结果
	half.Next = reverseList(tail)

	return true

}

func endOfFirstHalf(head *ListNode) *ListNode {
	//快慢指针找到后半部分
	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

//反转链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {

	// 1->2->2->1
	node1 := ListNode{1, nil}
	node2 := ListNode{2, nil}
	node3 := ListNode{3, nil}
	node4 := ListNode{2, nil}
	node5 := ListNode{1, nil}
	node1.Next = &node2
	node2.Next = &node3

	node3.Next = &node4
	node4.Next = &node5
	fmt.Println(isPalindrome(&node1))
	fmt.Println(isPalindrome1(&node1))

}
