package main

import "fmt"

// 19. 删除链表的倒数第 N 个结点
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

// 示例 1：
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]

// 示例 2：
// 输入：head = [1], n = 1
// 输出：[]

// 示例 3：
// 输入：head = [1,2], n = 1
// 输出：[1]

// 提示：
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//时间复杂度：O(L)，其中 LL 是链表的长度。空间复杂度：O(1)。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	len := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < len-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func getLength(head *ListNode) int {
	len := 0
	for node := head; node != nil; node = node.Next {
		len++

	}
	return len
}

////时间复杂度：O(L)，其中 LL 是链表的长度。空间复杂度：O(1)。
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast := head
	slow := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {

		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next

}

func main() {
	//head = [1,2,3,4,5], n = 2
	listNode1 := ListNode{1, nil}
	listNode2 := ListNode{2, nil}
	listNode3 := ListNode{3, nil}
	listNode4 := ListNode{4, nil}
	listNode5 := ListNode{5, nil}
	listNode1.Next = &listNode2
	listNode2.Next = &listNode3
	listNode3.Next = &listNode4
	listNode4.Next = &listNode5
	node := removeNthFromEnd1(&listNode1, 2)
	for ; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}

}
