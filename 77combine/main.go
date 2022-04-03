package main

import "fmt"

// 77. 组合
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

// 你可以按 任何顺序 返回答案。

// 示例 1：
// 输入：n = 4, k = 2
// 输出：
// [
//   [2,4],
//   [3,4],
//   [2,3],
//   [1,2],
//   [1,3],
//   [1,4],
// ]

// 示例 2：
// 输入：n = 1, k = 1
// 输出：[[1]]

// 提示:
// 1 <= n <= 20
// 1 <= k <= n

func combine(n int, k int) [][]int {
	var res [][]int
	var set []int
	var dfs func(int)
	dfs = func(cur int) {
		if len(set)+(n-cur+1) < k {
			return
		}
		//记录合法答案
		if len(set) == k {
			tmp := make([]int, k)
			copy(tmp, set)
			res = append(res, tmp)
			return
		}
		// if cur == n+1 {
		// 	return

		// }
		//考虑当前位置
		set = append(set, cur)
		dfs(cur + 1)
		//不考虑
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(1)

	return res

}
func main() {

	//n = 4, k = 2
	// [
	//   [2,4],
	//   [3,4],
	//   [2,3],
	//   [1,2],
	//   [1,3],
	//   [1,4],
	// ]
	n := 4
	k := 2
	fmt.Println(combine(n, k))

}
