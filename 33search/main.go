package main

import "fmt"

//33 搜索旋转排序数组
//
//整数数组 nums 按升序排列，数组中的值 互不相同 。
//
//在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
//
//给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//
//
//
//示例 1：
//
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
//示例 2：
//
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1
//示例 3：
//
//输入：nums = [1], target = 0
//输出：-1
//
//
//提示：
//
//1 <= nums.length <= 5000
//-10^4 <= nums[i] <= 10^4
//nums 中的每个值都 独一无二
//题目数据保证 nums 在预先未知的某个下标上进行了旋转
//-10^4 <= target <= 10^4
//
//
//进阶：你可以设计一个时间复杂度为 O(log n) 的解决方案吗？
func search(nums []int, target int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1

}

func search1(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {

			return 0
		} else {
			return -1
		}

	}
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {

				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1

			} else {
				right = mid - 1
			}
		}

	}
	return -1

}

func search2(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {

			return 0
		} else {
			return -1
		}

	}
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {

				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				left = mid + 1

			} else {
				right = mid - 1
			}
		}

	}
	return -1

}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}

	target := 0
	ret := search(nums, target)
	fmt.Println(ret)

	ret1 := search1(nums, target)
	fmt.Println(ret1)

	ret2 := search2(nums, target)
	fmt.Println(ret2)
	fmt.Println((1 + 2) / 2)

}
