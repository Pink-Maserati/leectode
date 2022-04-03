package main

import (
	"fmt"
	"sort"
)

// 56. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

// 示例 1：
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

// 示例 2：
// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]
// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

// 提示：
// 1 <= intervals.length <= 104
// intervals[i].length == 2
// 0 <= starti <= endi <= 104
func merge(intervals [][]int) [][]int {
	//从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res = [][]int{}
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		left := intervals[i][0]
		right := intervals[i][1]
		m := len(res)
		if res[m-1][1] < left {
			res = append(res, intervals[i])
		} else {
			res[m-1][1] = max(res[m-1][1], right)
		}

	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
