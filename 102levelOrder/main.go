package main

import (
	"fmt"
	"time"
)

// 102. 二叉树的层序遍历
// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

// 示例 1：
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]

// 示例 2：
// 输入：root = [1]
// 输出：[[1]]

// 示例 3：
// 输入：root = []
// 输出：[]

// 提示：
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	//将根节点放入队列中
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := []int{}
		///获取当前队列的长度，这个长度相当于 当前这一层的节点个数
		n := len(q)
		//将队列中的元素都拿出来(也就是获取这一层的节点)，放到临时list中
		//如果节点的左/右子树不为空，也放入队列中
		for n > 0 {
			node := q[0]
			//fmt.Println(len(q), node.Val, n)
			q = q[1:]
			tmp = append(tmp, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			n--
		}
		res = append(res, tmp)

	}
	return res
}

func levelOrder1(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func main() {
	//root = [3,9,20,null,null,15,7]

	//[[3],[9,20],[15,7]]
	treeNode1 := TreeNode{3, nil, nil}
	treeNode2 := TreeNode{9, nil, nil}
	treeNode3 := TreeNode{20, nil, nil}
	treeNode4 := TreeNode{15, nil, nil}
	treeNode5 := TreeNode{7, nil, nil}
	treeNode1.Left = &treeNode2
	treeNode1.Right = &treeNode3
	treeNode3.Left = &treeNode4
	treeNode3.Right = &treeNode5

	start := time.Now()
	res := levelOrder(&treeNode1)
	for _, v := range res {
		fmt.Println(v)
	}

	res1 := levelOrder1(&treeNode1)
	for _, v := range res1 {
		fmt.Println(v)
	}
	diff := time.Since(start)
	fmt.Println(diff)

}
