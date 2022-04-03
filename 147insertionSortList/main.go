package main

import (
	"fmt"
)

// 147. 对链表进行插入排序
// 给定单个链表的头 head ，使用 插入排序 对链表进行排序，并返回 排序后链表的头 。

// 插入排序 算法的步骤:

// 插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
// 每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
// 重复直到所有输入数据插入完为止。
// 下面是插入排序算法的一个图形示例。部分排序的列表(黑色)最初只包含列表中的第一个元素。每次迭代时，从输入数据中删除一个元素(红色)，并就地插入已排序的列表中。

// 对链表进行插入排序。

// 示例 1：
// 输入: head = [4,2,1,3]
// 输出: [1,2,3,4]

// 示例 2：
// 输入: head = [-1,5,3,4,0]
// 输出: [-1,0,3,4,5]

// 提示：
// 列表中的节点数在 [1, 5000]范围内
// -5000 <= Node.val <= 5000

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//插入排序 时复杂度o(N^2) N为链表的长度 空间复杂度 o(1)
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}

	//lastSorted为链表的已排序部分的最后一个节点
	lastSorted, cur := head, head.Next
	for cur != nil {
		if lastSorted.Val <= cur.Val {
			lastSorted = lastSorted.Next
		} else {
			pre := dummyHead
			for pre.Next.Val <= cur.Val {
				pre = pre.Next
			}
			lastSorted.Next = cur.Next
			cur.Next = pre.Next
			pre.Next = cur

		}
		cur = lastSorted.Next
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

	node := insertionSortList(&node1)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}

}
