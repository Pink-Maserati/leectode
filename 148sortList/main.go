package main

import (
	"fmt"
)

// 148. 排序链表
// 给你链表的头结点 head ，请将其按升序 排列并返回 排序后的链表 。

// 示例 1：
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]

// 示例 2：
// 输入：head = [-1,5,3,4,0]
// 输出：[-1,0,3,4,5]

// 示例 3：
// 输入：head = []
// 输出：[]

// 提示：
// 链表中节点的数目在范围 [0, 5 * 104] 内
// -105 <= Node.val <= 105

// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sort(head, tail *ListNode) *ListNode {
	//先把链表分成2部分
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	fast, slow := head, head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

//时间复杂度：O(nlogn)，其中 n 是链表的长度。

//空间复杂度：O(logn)，其中 n 是链表的长度。空间复杂度主要取决于递归调用的栈空间。

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tmp, tmp1, tmp2 := dummyHead, head1, head2
	for tmp1 != nil && tmp2 != nil {

		if tmp1.Val <= tmp2.Val {
			tmp.Next = tmp1

			tmp1 = tmp1.Next
		} else {

			tmp.Next = tmp2
			tmp2 = tmp2.Next
		}

		tmp = tmp.Next
	}

	if tmp1 != nil {
		tmp.Next = tmp1
	}
	if tmp2 != nil {
		tmp.Next = tmp2
	}
	return dummyHead.Next
}

func merge1(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func sortList1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge1(head1, head2)

			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummyHead.Next
}

func main() {
	//head = [4,2,1,3]

	node1 := ListNode{4, nil}
	node2 := ListNode{2, nil}
	node3 := ListNode{1, nil}
	node4 := ListNode{3, nil}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	fmt.Println(3 << 1)

	// cur := &node1
	// for cur != nil {
	// 	fmt.Println(cur.Val)
	// 	cur = cur.Next
	// }

	// cur := sortList(&node1)
	// //node := insertionSortList(&node1)
	// for cur != nil {

	// 	fmt.Println(cur.Val)
	// 	cur = cur.Next
	// }
	noder := sortList1(&node1)
	//node := insertionSortList(&node1)
	cur1 := noder
	for cur1 != nil {

		fmt.Println(cur1.Val)
		cur1 = cur1.Next
	}
}
