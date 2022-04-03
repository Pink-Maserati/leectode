package main

import (
	"fmt"
	"sort"
)

// 90. 子集 II
// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

// 示例 1：
// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]

// 示例 2：
// 输入：nums = [0]
// 输出：[[],[0]]

// 提示：
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	set := []int{}
	var dfs func(bool, int)
	dfs = func(chooice bool, cur int) {
		if cur == len(nums) {
			result = append(result, append([]int(nil), set...))
			return
		}

		dfs(false, cur+1)
		//在递归时，若发现没有选择上一个数，且当前数字与上一个数相同，则可以跳过当前生成的子集。
		if !chooice && cur > 0 && nums[cur] == nums[cur-1] {
			return
		}

		set = append(set, nums[cur])
		dfs(true, cur+1)
		set = set[:len(set)-1]

	}
	dfs(false, 0)
	return result

}

func main() {
	nums := []int{1, 1}
	fmt.Println(subsetsWithDup(nums))

}
