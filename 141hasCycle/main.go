package main

import (
	"fmt"
)

// 141. 环形链表
// 给你一个链表的头节点 head ，判断链表中是否有环。

// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。

// 如果链表中存在环 ，则返回 true 。 否则，返回 false 。

// 示例 1：
// 输入：head = [3,2,0,-4], pos = 1
// 输出：true
// 解释：链表中有一个环，其尾部连接到第二个节点。

// 示例 2：
// 输入：head = [1,2], pos = 0
// 输出：true
// 解释：链表中有一个环，其尾部连接到第一个节点。

// 示例 3：
// 输入：head = [1], pos = -1
// 输出：false
// 解释：链表中没有环。

// 提示：
// 链表中节点的数目范围是 [0, 104]
// -105 <= Node.val <= 105
// pos 为 -1 或者链表中的一个 有效索引 。

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	var fast *ListNode
	var slow *ListNode
	slow = head
	fast = head

	fmt.Println(fast)
	fmt.Println(slow)
	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next
		fmt.Println(fast)
		fmt.Println(slow)
		if fast == slow {
			return true
		}

	}

	return false
}

//哈希表：遍历所有节点，每次遍历到一个节点时，判断该节点此前是否被访问过
//使用哈希表来存储所有已经访问过的节点。每次我们到达一个节点，如果该节点已经存在于哈希表中，则说明该链表是环形链表，
//否则就将该节点加入哈希表中
func hasCycle1(head *ListNode) bool {

	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next

	}

	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {
	//head = [3,2,0,-4], pos = 1
	var node2 ListNode
	// var node3 ListNode
	// var node4 ListNode
	node1 := ListNode{3, &node2}
	node2 = ListNode{2, nil}
	// node3 = ListNode{0, &node4}
	// node4 = ListNode{-4, &node2}
	fmt.Println(hasCycle(&node1))

	fmt.Println(hasCycle1(&node1))
	fmt.Println(hasCycle2(&node1))
	// fmt.Println(node1)
	// fmt.Println(node2)

}
