package main

import (
	"fmt"
)

// 206. 反转链表
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

// 示例 1：
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]

// 示例 2：
// 输入：head = [1,2]
// 输出：[2,1]

// 示例 3：
// 输入：head = []
// 输出：[]

// 提示：
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//在遍历链表时，将当前节点的next指针改为指向前一个节点。
//由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。在更改引用之前，还需要存储后一个节点。最后返回新的头引用。
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	fmt.Println(pre, cur)
	if head == nil || head.Next == nil {
		return head
	}

	for cur != nil {

		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		fmt.Println(pre, cur)

	}
	return pre
	// var prev *ListNode
	// curr := head
	// for curr != nil {
	// 	next := curr.Next
	// 	curr.Next = prev
	// 	prev = curr
	// 	curr = next
	// }
	// return prev

}

func main() {
	//head = [1,2]
	node1 := ListNode{1, nil}
	node2 := ListNode{2, nil}
	node3 := ListNode{3, nil}
	node1.Next = &node2

	node2.Next = &node3
	fmt.Println(node1, node1.Next.Val, node2, node2.Next.Val)
	fmt.Println(reverseList(&node1))

}
