package main

import (
	"fmt"
)

// 160. 相交链表
// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

// 图示两个链表在节点 c1 开始相交：

// 题目数据 保证 整个链式结构中不存在环。

// 注意，函数返回结果后，链表必须 保持其原始结构 。

// 自定义评测：

// 评测系统 的输入如下（你设计的程序 不适用 此输入）：

// intersectVal - 相交的起始节点的值。如果不存在相交节点，这一值为 0
// listA - 第一个链表
// listB - 第二个链表
// skipA - 在 listA 中（从头节点开始）跳到交叉节点的节点数
// skipB - 在 listB 中（从头节点开始）跳到交叉节点的节点数
// 评测系统将根据这些输入创建链式数据结构，并将两个头节点 headA 和 headB 传递给你的程序。如果程序能够正确返回相交节点，那么你的解决方案将被 视作正确答案 。

// 示例 1：
// 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
// 输出：Intersected at '8'
// 解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
// 从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,6,1,8,4,5]。
// 在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

// 示例 2：
// 输入：intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
// 输出：Intersected at '2'
// 解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
// 从各自的表头开始算起，链表 A 为 [1,9,1,2,4]，链表 B 为 [3,2,4]。
// 在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。

// 示例 3：
// 输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
// 输出：null
// 解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
// 由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
// 这两个链表不相交，因此返回 null 。

// 提示：

// listA 中节点数目为 m
// listB 中节点数目为 n
// 1 <= m, n <= 3 * 104
// 1 <= Node.val <= 105
// 0 <= skipA <= m
// 0 <= skipB <= n
// 如果 listA 和 listB 没有交点，intersectVal 为 0
// 如果 listA 和 listB 有交点，intersectVal == listA[skipA] == listB[skipB]

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//双指针	使用双指针的方法，可以将空间复杂度降至 O(1)O。

	// 只有当链表 headA 和 headB 都不为空时，两个链表才可能相交。因此首先判断链表 headA 和 headB 是否为空，如果其中至少有一个链表为空，则两个链表一定不相交，返回 null。

	// 当链表 headA 和  headB 都不为空时，创建两个指针 pA pB，初始时分别指向两个链表的头节点 headA 和 headB，然后将两个指针依次遍历两个链表的每个节点。具体做法如下：

	// 每步操作需要同时更新指针 pA 和 pB。

	// 如果指针 pA 不为空，则将指针 A 移到下一个节点；如果指针 pB 不为空，则将指针 pB 移到下一个节点。

	// 如果指针 pA 为空，则将指针 pA 移到链表 headB 的头节点；如果指针 pB 为空，则将指针 pB 移到链表 headA 的头节点。
	//当指针 pA 和 pB 指向同一个节点或者都为空时，返回它们指向的节点或者 null。

	if headA == nil || headB == nil {
		return nil
	}
	nodeA := headA
	nodeB := headB
	for nodeA != nodeB {
		if nodeA == nil {
			nodeA = headB
		} else {
			nodeA = nodeA.Next
		}

		if nodeB == nil {
			nodeB = headA
		} else {
			nodeB = nodeB.Next
		}
	}

	return nodeA

}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	//哈希集合
	res := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		res[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if res[tmp] {
			return tmp
		}
	}
	return nil
}

func main() {
	//listA = [1,9,1,2,4], listB = [3,2,4]
	headA := ListNode{1, nil}
	A2 := ListNode{9, nil}
	A3 := ListNode{1, nil}
	A4 := ListNode{2, nil}
	A5 := ListNode{4, nil}

	node := ListNode{2, nil}
	headA.Next = &node

	node.Next = &A2

	A2.Next = &A3
	A3.Next = &A4
	A4.Next = &A5

	headB := ListNode{3, nil}
	headB.Next = &A4
	res := getIntersectionNode(&headA, &headB)
	fmt.Println("res:")
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
	res1 := getIntersectionNode1(&headA, &headB)
	fmt.Println("res1:")
	for res1 != nil {
		fmt.Println(res1.Val)
		res1 = res1.Next
	}
}
