package main

import "fmt"

// 110. 平衡二叉树
// 给定一个二叉树，判断它是否是高度平衡的二叉树。

// 本题中，一棵高度平衡二叉树定义为：

// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。

// 示例 1：

// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
// 示例 2：

// 输入：root = [1,2,2,3,3,null,null,4,4]
// 输出：false
// 示例 3：

// 输入：root = []
// 输出：true

// 提示：

// 树中的节点数在范围 [0, 5000] 内
// -104 <= Node.val <= 104

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	// 退出条件
	// 1. 节点为空，则为平衡二叉树
	if root == nil {
		return true

	}

	//2. 左右节点不是平衡二叉树，则根节点不是平衡二叉树
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	//  // 单层逻辑
	// 1. 计算左节点深度
	// 2. 计算右节点深度
	// 3. 计算左右节点深度的差的绝对值是否>阈值1
	deepLeft := deep(root.Left)
	deepRight := deep(root.Right)
	if deepLeft-deepRight > 1 || deepRight-deepLeft > 1 {
		return false
	} else {
		return true
	}

}

func deep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//  单层逻辑 深度=Max(左节点深度,右节点深度)+1
	deepLeft := deep(root.Left)
	deepRight := deep(root.Right)
	if deepLeft > deepRight {
		return deepLeft + 1
	}
	return deepRight + 1

}

func isBalanced1(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	//输入：root = [3,9,20,null,null,15,7] 输出：true
	treeNode1 := TreeNode{3, nil, nil}
	treeNode2 := TreeNode{9, nil, nil}
	treeNode3 := TreeNode{20, nil, nil}
	treeNode4 := TreeNode{15, nil, nil}
	treeNode5 := TreeNode{7, nil, nil}

	treeNode1.Left = &treeNode2
	treeNode1.Right = &treeNode3
	treeNode3.Left = &treeNode4
	treeNode3.Right = &treeNode5
	fmt.Println(isBalanced(&treeNode1))
	fmt.Println(isBalanced1(&treeNode1))

}
