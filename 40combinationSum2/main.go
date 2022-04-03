package main

import (
	"fmt"
	"sort"
)

// 40. 组合总和 II
// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

// candidates 中的每个数字在每个组合中只能使用 一次 。

// 注意：解集不能包含重复的组合。

// 示例 1:
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 输出:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]

// 示例 2:
// 输入: candidates = [2,5,2,1,2], target = 5,
// 输出:
// [
// [1,2,2],
// [5]
// ]

// 提示:
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var set []int
	var dfs func(int, int, bool)
	sort.Ints(candidates)
	dfs = func(target int, cur int, chooice bool) {
		if target == 0 {
			res = append(res, append([]int(nil), set...))

			return
		}
		if cur == len(candidates) {
			return
		}

		dfs(target, cur+1, false)
		if !chooice && cur > 0 && candidates[cur-1] == candidates[cur] {
			return
		}

		if target-candidates[cur] >= 0 {
			set = append(set, candidates[cur])
			dfs(target-candidates[cur], cur+1, true)
			set = set[:len(set)-1]

		}

	}
	dfs(target, 0, false)

	return res

}

func combinationSum22(candidates []int, target int) [][]int {
	var res [][]int
	n := len(candidates)
	if n == 0 {
		return res
	}
	//排序是关键步骤
	sort.Ints(candidates)
	var path []int
	var dfs func(int, int)

	dfs = func(target int, cur int) {
		if target == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}

		for i := cur; i < n; i++ {
			// 大剪枝：减去 candidates[i] 小于 0，减去后面的 candidates[i + 1]、candidates[i + 2] 肯定也小于 0，因此用 break

			if target-candidates[i] < 0 {
				break
			}

			// 小剪枝：同一层相同数值的结点，从第 2 个开始，候选数更少，结果一定发生重复，因此跳过，用 continue
			if i > cur && candidates[i] == candidates[i-1] {
				continue
			}

			path = append(path, candidates[i])
			// 调试语句 ①
			fmt.Printf("递归之前 => %v ,剩余 = %v\n", path, target-candidates[i])

			// 因为元素不可以重复使用，这里递归传递下去的是 i + 1 而不是 i
			dfs(target-candidates[i], i+1)

			path = path[:len(path)-1]

			// 调试语句 ②
			fmt.Printf("递归之后 => %v ,剩余 = %v\n", path, target-candidates[i])
		}
	}
	dfs(target, 0)

	return res
}

func main() {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	fmt.Println([]int(nil))

	c := []int(nil)
	fmt.Println(c)
	fmt.Println(combinationSum2(candidates, target))
	fmt.Println(combinationSum22(candidates, target))

}
