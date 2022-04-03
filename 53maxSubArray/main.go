package main

import (
	"fmt"
)

// 53. 最大子数组和
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

// 子数组 是数组中的一个连续部分。

// 示例 1：
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

// 示例 2
// 输入：nums = [1]
// 输出：1

// 示例 3：
// 输入：nums = [5,4,-1,7,8]
// 输出：23

// 提示：

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104

// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

//时间复杂度：O(n)，其中 n为 nums 数组的长度。我们只需要遍历一遍数组即可求得答案。
//空间复杂度：O(n)。我们只需要常数空间存放若干变量。
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))

	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])

	}
	return dp[len(nums)-1]

}

func maxSubArray1(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]

		}
		if nums[i] > max {
			max = nums[i]
		}

	}

	return max

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//nums = [5,4,-1,7,8]
	nums := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums))
	fmt.Println(maxSubArray1(nums))

}
