package main

import "fmt"

// 78. 子集
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

// 示例 1：

// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 示例 2：

// 输入：nums = [0]
// 输出：[[],[0]]

// 提示：

// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// nums 中的所有元素 互不相同

func subsets(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})

	for _, v := range nums {
		n := len(result)
		for j := 0; j < n; j++ {
			//这里一定要用copy
			tmp := make([]int, len(result[j]))
			copy(tmp, result[j])

			tmp = append(tmp, v)
			result = append(result, tmp)

		}

	}

	return result

}
func subsets1(nums []int) [][]int {
	var result [][]int

	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask&(1<<i) > 0 {
				set = append(set, v)
			}
		}
		result = append(result, append([]int(nil), set...))
	}

	return result
}

func subsets2(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}

func main() {
	//nums = [1,2,3]
	// nums := []int{1, 2, 3}
	nums := []int{9, 0, 3, 5, 7}

	fmt.Println(subsets(nums))

	fmt.Println(subsets1(nums))
	fmt.Println(subsets2(nums))

}
