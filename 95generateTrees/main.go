package main

import "fmt"

// 95. 不同的二叉搜索树 II
// 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。

// 示例 1：
// 输入：n = 3
// 输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

// 示例 2：
// 输入：n = 1
// 输出：[[1]]

// 提示：
// 1 <= n <= 8

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)

}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTreeNodes := []*TreeNode{}
	//枚举可行根节点
	for i := start; i <= end; i++ {
		//获得所有可行的左子树集合
		leftTreeNodes := generate(start, i-1)
		//获得所有可行的右子树集合
		rightTreeNodes := generate(i+1, end)
		//从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点
		for _, l := range leftTreeNodes {
			for _, r := range rightTreeNodes {
				tmp := &TreeNode{i, nil, nil}
				tmp.Left = l
				tmp.Right = r
				allTreeNodes = append(allTreeNodes, tmp)
			}
		}

	}
	return allTreeNodes
}

func main() {
	n := 3
	treeNodes := generateTrees(n)
	for _, v := range treeNodes {
		fmt.Println(v)
	}
}
