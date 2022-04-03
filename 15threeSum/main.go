package main

import (
	"fmt"
	"sort"
)

// 15. 三数之和
// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

// 注意：答案中不可以包含重复的三元组。

// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]

// 示例 2：
// 输入：nums = []
// 输出：[]

// 示例 3：
// 输入：nums = [0]
// 输出：[]

// 提示：

// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105
func threeSum(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	if n < 3 {
		return res
	}
	sort.Ints(nums)
	for first := 0; first < n-2; first++ {
		if nums[first] > 0 {
			break
		}
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second := first + 1
		third := n - 1
		for second < third {
			if nums[first]+nums[second]+nums[third] == 0 {
				res = append(res, []int{nums[first], nums[second], nums[third]})
				for second < third && nums[second] == nums[second+1] {
					second++
				}
				for second < third && nums[third] == nums[third-1] {
					third--
				}
				second++
				third--
			} else if nums[first]+nums[second]+nums[third] < 0 {
				second++
			} else {
				third--
			}
		}
	}
	return res

}
func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(nums)
	result := threeSum(nums)
	fmt.Println(result)
}
