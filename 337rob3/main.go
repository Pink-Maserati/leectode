package main

import "fmt"

// 337. 打家劫舍 III
// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。

// 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。

// 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。

// 示例 1:
// 输入: root = [3,2,3,null,3,null,1]
// 输出: 7
// 解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7

// 示例 2:
// 输入: root = [3,4,5,1,3,null,1]
// 输出: 9
// 解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9

// 提示：
// 树的节点数在 [1, 104] 范围内
// 0 <= Node.val <= 104

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//时间复杂度是 O(n) 空间复杂度也是 O(n)
func rob(root *TreeNode) int {
	var f map[*TreeNode]int //表示选择o节点的情况下，o节点的子树上被选择的节点的最大权值和
	var g map[*TreeNode]int //不选择o节点的情况下，o 节点的子树上被选择的节点的最大权值和；l 和 r代表 o的左右孩子
	//当 o 被选中时，o 的左右孩子都不能被选中，故 o 被选中情况下子树上被选中点的最大权值和为 l 和 r 不被选中的最大权值和相加，即 f(o) = g(l) + g(r)。
	//当 o不被选中时，o 的左右孩子可以被选中，也可以不被选中。对于 o 的某个具体的孩子 x，它对 o的贡献是 x 被选中和不被选中情况下权值和的较大值。故g(o)=max{f(l),g(l)}+max{f(r),g(r)}。
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		f[node] = node.Val + g[node.Left] + g[node.Right]
		g[node] = max(f[node.Left], g[node.Left]) + max(f[node.Right], g[node.Right])

	}
	f = make(map[*TreeNode]int, 0)
	g = make(map[*TreeNode]int, 0)
	dfs(root)
	return max(f[root], g[root])

}

//时间复杂度：O(n)。对二叉树做了一次后序遍历。
//空间复杂度：O(n)。虽然优化过的版本省去了哈希表的空间，但是栈空间的使用代价依旧是 O(n)，故空间复杂度不变
func rob1(root *TreeNode) int {
	res := dfs(root)
	return max(res[0], res[1])

}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	selected := root.Val + l[1] + r[1]
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	treeNode1 := TreeNode{3, nil, nil}
	treeNode2 := TreeNode{4, nil, nil}
	treeNode3 := TreeNode{5, nil, nil}
	treeNode4 := TreeNode{1, nil, nil}
	treeNode5 := TreeNode{3, nil, nil}
	treeNode6 := TreeNode{1, nil, nil}
	treeNode1.Left = &treeNode2
	treeNode1.Right = &treeNode3
	treeNode2.Left = &treeNode4
	treeNode2.Right = &treeNode5
	treeNode3.Right = &treeNode6

	fmt.Println(rob(&treeNode1))
	fmt.Println(rob1(&treeNode1))

}
