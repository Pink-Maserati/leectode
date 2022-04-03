package main

import "fmt"

// 144. 二叉树的前序遍历
// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

// 示例 1
// 输入：root = [1,null,2,3]
// 输出：[1,2,3]

// 示例 2：
// 输入：root = []
// 输出：[]

// 示例 3：
// 输入：root = [1]
// 输出：[1]

// 示例 4
// 输入：root = [1,2]
// 输出：[1,2]
// 示例 5：

// 输入：root = [1,null,2]
// 输出：[1,2]

// 提示：
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100

// Definition for a binary tree node.
type TreeNode struct {
	Val  int
	Left *TreeNode

	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return

}

func preorderTraversal1(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

func main() {
	//root = [1,null,2,3]
	treeNode1 := TreeNode{1, nil, nil}
	treeNode2 := TreeNode{2, nil, nil}
	treeNode3 := TreeNode{3, nil, nil}
	treeNode1.Right = &treeNode2
	treeNode2.Left = &treeNode3
	fmt.Println(preorderTraversal(&treeNode1))

	fmt.Println(preorderTraversal1(&treeNode1))

}
