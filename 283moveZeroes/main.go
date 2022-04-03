package main

import (
	"fmt"
)

// 283. 移动零
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

// 示例 1:

// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 示例 2:

// 输入: nums = [0]
// 输出: [0]

// 提示:

// 1 <= nums.length <= 104
// -231 <= nums[i] <= 231 - 1

// 进阶：你能尽量减少完成的操作次数吗？

//时间复杂度：O(n)，其中 n 为序列长度。每个位置至多被遍历两次。
//空间复杂度：O(1)。只需要常数的空间存放若干变量。
func moveZeroes(nums []int) {
	n := len(nums)
	left := 0
	right := 0
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

}

//使用for循环 时间复杂度：O(n) 声明临时数组 空间复杂度：O(n)
func moveZeroes1(nums []int) {
	i := 0

	for _, v := range nums {
		if v != 0 {
			nums[i] = v
			i++

		}
	}

	for i < len(nums) {

		nums[i] = 0
		i++

	}

}

func main() {
	// nums = [0,1,0,3,12]
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
	nums1 := []int{0, 1, 0, 3, 12}
	moveZeroes1(nums1)
	fmt.Println(nums1)
}
