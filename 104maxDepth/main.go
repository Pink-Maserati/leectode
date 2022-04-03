package main

import "fmt"

// 104. 二叉树的最大深度
// 给定一个二叉树，找出其最大深度。

// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

// 说明: 叶子节点是指没有子节点的节点。

// 示例：
// 给定二叉树 [3,9,20,null,null,15,7]，

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回它的最大深度 3

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	treeNode1 := TreeNode{3, nil, nil}
	treeNode2 := TreeNode{9, nil, nil}
	treeNode3 := TreeNode{20, nil, nil}
	treeNode4 := TreeNode{15, nil, nil}
	treeNode5 := TreeNode{7, nil, nil}
	treeNode1.Left = &treeNode2
	treeNode1.Right = &treeNode3
	treeNode3.Left = &treeNode4
	treeNode3.Right = &treeNode5
	fmt.Println(maxDepth(&treeNode1))
}
